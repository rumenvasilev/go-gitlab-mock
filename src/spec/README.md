# Generating the OpenAPI spec for Gitlab

At the time of writing, Gitlab are not publishing a complete OpenAPI specification of their API. Or at least I couldn't find one. So I started writing just the minimal one, based on the available documentation here: https://docs.gitlab.com/ee/api/rest/. That's a herculean effort though. There has to be an easier way!

Turns out even GitLab don't have the full OpenAPI spec yet (https://gitlab.com/gitlab-org/gitlab-foss/-/issues/55028) and that work is still very much WIP (https://gitlab.com/gitlab-org/gitlab/-/blame/master/lib/api/api.rb#L205). The current OpenAPI specification you can find here: https://gitlab.com/gitlab-org/gitlab/-/tree/master/doc/api/openapi (v2 is the latest). You'll quickly realise there's barely anything in there though.

While I was researching how to extract that API spec, I came across an internal instruction video, how to onboard more APIs to the OpenAPI spec generator (ruby gem) and how to generate the spec afterwards - https://www.youtube.com/watch?v=V1hUlFOBbYY.

And this is what I used in the end.

## Generating OpenAPI spec for Gitlab REST API

* Clone https://gitlab.com/gitlab-org/gitlab
* Add a line `mount ::API::Issues` after 276 (`mount ::API:IssueLinks`) to this file https://gitlab.com/gitlab-org/gitlab/-/blob/master/lib/api/api.rb
* Build the Dockerfile and run it from the root of the gitlab cloned repository: `docker run -it --rm -v .:/tmp/code $myimage bash`
* Once in docker:

```bash
cd /tmp/code
bundle install
cp config/database.yml.postgresql config/database.yml
cp config/gitlab.yml.example config/gitlab.yml
rake gitlab:openapi:generate
```

You will end up with an updated OpenAPI specification under `doc/api/openapi/openapi_v2.yaml` (also available on your disk, not just within docker), which includes loads more.

## Regenerating mock endpoints for Gitlab's REST API

Copy the openapi spec to `src/spec` and run `go run main.go -spec-path src/spec/openapi.yaml`. You will end up with an updated `src/mock/endpointpattern.go`.
