package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/buger/jsonparser"
	"sigs.k8s.io/yaml"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/rumenvasilev/go-gitlab-mock/src/gen"
)

var debug bool
var specPath string

func init() {
	flag.BoolVar(&debug, "debug", false, "output debug information")
	flag.StringVar(&specPath, "spec-path", "", "path to openapi spec on disk")
}

func main() {
	flag.Parse()

	var l log.Logger

	l = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	l = log.With(l, "caller", log.DefaultCaller)

	if debug {
		l = level.NewFilter(l, level.AllowDebug())
		level.Debug(l).Log("msg", "running in debug mode")
	} else {
		l = level.NewFilter(l, level.AllowInfo())
	}

	var apiDefinition []byte
	var err error
	if specPath != "" {
		apiDefinition, err = os.ReadFile(specPath)
		if err != nil {
			l.Log("Couldn't read OpenAPI spec file from disk")
			os.Exit(1)
		}
		if strings.HasSuffix(specPath, "yaml") {
			apiDefinition, err = yaml.YAMLToJSON(apiDefinition)
			if err != nil {
				l.Log("Couldn't parse yaml")
				os.Exit(1)
			}
		}
	} else {
		y := gen.FetchAPIDefinition(l)
		apiDefinition, err = yaml.YAMLToJSON(y)
		if err != nil {
			l.Log("Couldn't parse yaml")
			os.Exit(1)
		}
	}

	buf := bytes.NewBuffer([]byte(gen.OUTPUT_FILE_HEADER))

	jsonparser.ObjectEach(
		apiDefinition,
		func(key, endpointDefinition []byte, _ jsonparser.ValueType, _ int) error {
			endpointPattern := string(key)
			httpMethods := []string{}

			jsonparser.ObjectEach(
				endpointDefinition,
				func(key, _ []byte, _ jsonparser.ValueType, _ int) error {
					httpMethods = append(httpMethods, string(key))

					return nil
				},
			)

			for _, httpMethod := range httpMethods {
				code := gen.FormatToGolangVarNameAndValue(
					l,
					gen.ScrapeResult{
						HTTPMethod:      httpMethod,
						EndpointPattern: endpointPattern,
					},
				)
				buf.WriteString(code)
			}

			return nil
		},
		"paths",
	)

	os.WriteFile(
		gen.OUTPUT_FILEPATH,
		buf.Bytes(),
		0755,
	)

	errorsFound := false

	// to catch possible format errors
	if err := exec.Command("gofmt", "-w", "mock/endpointpattern.go").Run(); err != nil {
		level.Error(l).Log("msg", fmt.Sprintf("error executing gofmt: %s", err.Error()))
		errorsFound = true
	}

	// to catch everything else (hopefully)
	if err := exec.Command("go", "vet", "./...").Run(); err != nil {
		level.Error(l).Log("msg", fmt.Sprintf("error executing go vet: %s", err.Error()))
		errorsFound = true
	}

	if errorsFound {
		os.Exit(1)
	}

	fmt.Println("All done.")
}
