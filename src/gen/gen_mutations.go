package gen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// https://docs.gitlab.com/ee/api/packages/nuget.html#odata-package-entry-endpoints
var enabledMutators = map[string]func(ScrapeResult) ScrapeResult{
	"/api/v4/projects/{project_id}/packages/nuget/v2/FindPackagesById\\(\\)":                                                                    specialLastParamMutatorHelper(),
	"/api/v4/projects/{project_id}/packages/nuget/v2/Packages\\(\\)":                                                                            specialLastParamMutatorHelper(),
	"/api/v4/projects/{project_id}/packages/nuget/v2/Packages\\(Id='*package_name',Version='*package_version'\\)":                               specialLastParamMutatorHelper(), // curl "https://gitlab.example.com/api/v4/projects/1/packages/nuget/v2/Packages(Id='mynugetpkg',Version='1.0.0')"
	"/api/v4/group/{id}/-/packages/composer/*package_name":                                                                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/group/{id}/-/packages/composer/p2/*package_name":                                                                                   allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/InRelease":                                                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/Release":                                                                         allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/Release.gpg":                                                                     allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/binary-{architecture}/Packages":                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256}":                  allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/Packages":                     allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256}": allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/source/Sources":                                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/debian/dists/*distribution/{component}/source/by-hash/SHA256/{file_sha256}":                                 allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/maven/*path/{file_name}":                                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/npm/*package_name":                                                                                          allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/npm/-/package/*package_name/dist-tags":                                                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/npm/-/package/*package_name/dist-tags/{tag}":                                                                allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/nuget/metadata/*package_name/*package_version":                                                              allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/nuget/metadata/*package_name/index":                                                                         allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/pypi/files/{sha256}/*file_identifier":                                                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/groups/{id}/-/packages/pypi/simple/*package_name":                                                                                  allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/maven/*path/{file_name}":                                                                                                  allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/npm/*package_name":                                                                                                        allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/npm/-/package/*package_name/dist-tags":                                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/npm/-/package/*package_name/dist-tags/{tag}":                                                                              allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version":                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version/download":                           allowExtendedLastParamMutatorHelper(),
	"/api/v4/packages/terraform/modules/v1/{module_namespace}/{module_name}/{module_system}/*module_version/file":                               allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/jobs/artifacts/{ref_name}/raw/*artifact_path":                                                                        allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/jobs/{job_id}/artifacts/*artifact_path":                                                                              allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/composer/archives/*package_name":                                                                            allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/InRelease":                                                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/Release":                                                                         allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/Release.gpg":                                                                     allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/binary-{architecture}/Packages":                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/binary-{architecture}/by-hash/SHA256/{file_sha256}":                  allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/Packages":                     allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/debian-installer/binary-{architecture}/by-hash/SHA256/{file_sha256}": allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/source/Sources":                                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/debian/dists/*distribution/{component}/source/by-hash/SHA256/{file_sha256}":                                 allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/generic/{package_name}/*package_version/{file_name}":                                                        allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/generic/{package_name}/*package_version/{file_name}/authorize":                                              allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/go/*module_name/@v/list":                                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.info":                                                                   allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.mod":                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/go/*module_name/@v/{module_version}.zip":                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/maven/*path/{file_name}":                                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/maven/*path/{file_name}/authorize":                                                                          allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/ml_models/{model_name}/*model_version/{file_name}":                                                          allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/ml_models/{model_name}/*model_version/{file_name}/authorize":                                                allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/npm/*package_name":                                                                                          allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/npm/*package_name/-/*file_name":                                                                             allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/npm/-/package/*package_name/dist-tags":                                                                      allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/npm/-/package/*package_name/dist-tags/{tag}":                                                                allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/nuget/*package_name/*package_version":                                                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/nuget/download/*package_name/*package_version/*package_filename":                                            allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/nuget/download/*package_name/index":                                                                         allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/nuget/metadata/*package_name/*package_version":                                                              allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/nuget/metadata/*package_name/index":                                                                         allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/pypi/files/{sha256}/*file_identifier":                                                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/pypi/simple/*package_name":                                                                                  allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/rpm/*package_file_id/*file_name":                                                                            allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/rpm/repodata/*file_name":                                                                                    allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version/file":                                       allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/packages/terraform/modules/{module_name}/{module_system}/*module_version/file/authorize":                             allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/releases/permalink/latest(/)(*suffix_path)":                                                                          allowExtendedLastParamMutatorHelper(),
	"/api/v4/projects/{id}/releases/{tag_name}/downloads/*direct_asset_path":                                                                    allowExtendedLastParamMutatorHelper(),
}

// allowExtendedLastParamMutatorHelper mutates the last param of the endpoint pattern
// allowing it to have any characters (including slashes)
func allowExtendedLastParamMutatorHelper() func(ScrapeResult) ScrapeResult {
	return func(sr ScrapeResult) ScrapeResult {
		endpointSplits := strings.Split(sr.EndpointPattern, "/")
		// if parameter begins with *, replace it with value:.+
		r := regexp.MustCompile(`^\*[a-zA-Z-_.0-9]+$`)
		for index, v := range endpointSplits {
			if r.MatchString(v) {
				endpointSplits[index] = fmt.Sprintf("{%s:.+}", v[1:])
			}
		}

		sr.EndpointPattern = strings.Join(endpointSplits, "/")

		return sr
	}
}

// specialLastParamMutatorHelper mutates the last param of the endpoint pattern
// containing backslashes, braces and glob patterns
func specialLastParamMutatorHelper() func(ScrapeResult) ScrapeResult {
	return func(sr ScrapeResult) ScrapeResult {
		endpointSplits := strings.Split(sr.EndpointPattern, "/")
		lastParam := endpointSplits[len(endpointSplits)-1]

		// strip backslashes
		lastParam = strings.ReplaceAll(lastParam, `\`, "")

		// Packages(Id='*package_name',Version='*package_version')
		r := regexp.MustCompile(`^([a-zA-Z]+\([a-zA-Z]+=)'(\*([a-z_]+))',([a-zA-Z]+=)'(\*([a-z_]+))'\)$`)

		res := r.FindStringSubmatch(lastParam)
		if len(res) > 0 {
			t := "${1}'{${3}}',${4}'{${6}}')"
			lastParam = r.ReplaceAllString(lastParam, t)
		}

		endpointSplits[len(endpointSplits)-1] = lastParam

		sr.EndpointPattern = strings.Join(endpointSplits, "/")

		return sr
	}
}

// applyMutation applies mutation to the scrape result if necessary.
//
// There are some edge cases due to inconsistencies between GitHub's OpenAPI definition
// compared to the real world.
func applyMutation(sr ScrapeResult, l log.Logger) ScrapeResult {
	if mutator, found := enabledMutators[sr.EndpointPattern]; found {
		level.Debug(l).Log("msg", fmt.Sprintf("Mutating result %s", sr.EndpointPattern))
		return mutator(sr)
	}

	return sr
}
