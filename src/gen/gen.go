package gen

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/go-kit/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-kit/log/level"
)

const GITLAB_OPENAPI_DEFINITION_LOCATION = "https://gitlab.com/gitlab-org/gitlab/-/raw/master/doc/api/openapi/openapi.yaml"
const OUTPUT_FILE_HEADER = `package mock

// Code generated; DO NOT EDIT.

`
const OUTPUT_FILEPATH = "mock/endpointpattern.go"

type ScrapeResult struct {
	HTTPMethod      string
	EndpointPattern string
}

func FetchAPIDefinition(l log.Logger) []byte {
	resp, err := http.Get(GITLAB_OPENAPI_DEFINITION_LOCATION)

	if err != nil {
		level.Error(l).Log(
			"msg", "error fetching github's api definition",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		level.Error(l).Log(
			"msg", "error fetching github's api definition",
			"err", err.Error(),
		)

		os.Exit(1)
	}

	return bodyBytes
}

// FormatToGolangVarName generated the proper golang variable name
// given a endpoint format from the API
func FormatToGolangVarName(l log.Logger, sr ScrapeResult) string {
	c := cases.Title(language.English, cases.NoLower)
	result := c.String(strings.ToLower(sr.HTTPMethod))

	if sr.EndpointPattern == "/" {
		return result + "Slash"
	}

	// handles urls with dashes in them
	pattern := strings.ReplaceAll(sr.EndpointPattern, "-", "/")

	// cleans up varname when pattern was mutated
	// e.g see `GetApiV4GroupPackagesComposerByIdByPackageName`
	re := regexp.MustCompile(`[a-zA-Z0-9\/\{\}\_]+`)
	matches := re.FindAllString(pattern, -1)
	pattern = strings.Join(matches, "/")

	epSplit := strings.Split(
		pattern,
		"/",
	)

	// handle the first part of the variable name
	for _, part := range epSplit {
		if len(part) < 1 || string(part[0]) == "{" || part == "}" {
			continue
		}

		splitPart := strings.Split(part, "_")

		for _, p := range splitPart {
			result = result + c.String(p)
		}
	}

	//handle the "By`X`" part of the variable name
	for _, part := range epSplit {
		if len(part) < 1 {
			continue
		}

		if string(part[0]) == "{" {
			part = strings.ReplaceAll(part, "{", "")
			part = strings.ReplaceAll(part, "}", "")

			result += "By"

			for _, splitPart := range strings.Split(part, "_") {
				result += c.String(splitPart)
			}
		}
	}

	return result
}

func FormatToGolangVarNameAndValue(l log.Logger, sr ScrapeResult) string {
	sr = applyMutation(sr, l)

	return fmt.Sprintf(
		`var %s EndpointPattern = EndpointPattern{
	Pattern: "%s",
	Method:  "%s",
}
`,
		FormatToGolangVarName(l, sr),
		sr.EndpointPattern,
		strings.ToUpper(sr.HTTPMethod),
	) + "\n"
}
