# go-gitlab-mock
[![Go Reference](https://pkg.go.dev/badge/github.com/rumenvasilev/go-gitlab-mock.svg)](https://pkg.go.dev/github.com/rumenvasilev/go-gitlab-mock) [![Go Report Card](https://goreportcard.com/badge/github.com/rumenvasilev/go-gitlab-mock)](https://goreportcard.com/report/github.com/rumenvasilev/go-gitlab-mock) [![Build](https://github.com/rumenvasilev/go-gitlab-mock/actions/workflows/build.yaml/badge.svg)](https://github.com/rumenvasilev/go-gitlab-mock/actions/workflows/build.yaml)

***THIS LIBRARY IS NO LONGER NEEDED, GITLAB HAVE PUBLISHED THEIR OWN https://gitlab.com/gitlab-org/api/client-go/-/tree/main/testing***

A library to aid unit testing Go code that uses Gitlab SDK (currently based on `github.com/xanzy/go-gitlab`)

***I've copied (and adapted) the code generation and mock structure from github.com/migueleliasweb/go-github-mock. That is reflected in the LICENSE as well. To get to the current state though, I had to build GitLab's OpenAPI spec. It is not available on the internet (at least I couldn't find a near-complete spec), which is why I have generated it off their source code and then used it for the mocks.***

# Installation

```bash
go get github.com/rumenvasilev/go-gitlab-mock
```

# Features

- Create mocks for successive calls for the same endpoint
- Pagination support
- Mock error returns
- High level abstraction helps writing readabe unittests (see `mock.WithRequestMatch`)
- Lower level abstraction for advanced uses (see `mock.WithRequestMatchHandler`)

# Example

```
import "github.com/rumenvasilev/go-gitlab-mock/mock"
```

## Multiple requests

```golang
mockedURL := NewMockedHTTPServer(
    WithRequestMatch(
        GetApiV4UsersById,
        &gitlab.User{
            ID:   5,
            Name: "foobar",
        },
    ),
    WithRequestMatch(
        GetApiV4Groups,
        []*gitlab.Group{
            {
                Name: "foobar123thisorgwasmocked",
            },
        },
    ),
    WithRequestMatchHandler(
        GetApiV4Projects,
        http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
            w.Write(MustMarshal([]gitlab.Project{
                {
                    Name: "mocked-proj-1",
                },
                {
                    Name: "mocked-proj-2",
                },
            }))
        }),
    ),
)

c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL))

user, _, userErr := c.Users.GetUser(1, gitlab.GetUsersOptions{})

// user.Name == "foobar"

groups, _, groupsErr := c.Groups.ListGroups(&gitlab.ListGroupsOptions{})

// groups[0].Name == "foobar123thisorgwasmocked"

projs, _, projsErr := c.Projects.ListProjects(&gitlab.ListProjectsOptions{})

// projs[0].Name == "mocked-proj-1"
// projs[1].Name == "mocked-proj-2"
```

## Returning empty results

```golang
mockedURL := NewMockedHTTPServer(
    WithRequestMatch(
        GetApiV4ProjectsIssuesById,
        []*gitlab.Issue{
            {
                ID:    123,
                Title: "Issue 1",
            },
            {
                ID:    456,
                Title: "Issue 2",
            },
        },
        []*gitlab.Issue{},
    ),
)
c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL))

issues1, _, issues1Err := c.Issues.ListProjectIssues(1, &gitlab.ListProjectIssuesOptions{})

// len(issues1) == 2
// issues1Err == nil

issues2, _, issues2Err := c.Issues.ListProjectIssues(2, &gitlab.ListProjectIssuesOptions{})

// len(issues2) == 0
// issues2Err == nil
```

## Mocking errors from the API

```golang
mockedURL := NewMockedHTTPServer(
    WithRequestMatchHandler(
        GetApiV4UsersById,
        http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            WriteError(
                w,
                http.StatusInternalServerError,
                "test error we've defined",
            )
        }),
    ),
)

c, err := gitlab.NewClient(
    "",
    gitlab.WithBaseURL(mockedURL),
    gitlab.WithCustomBackoff(
        func(_, _ time.Duration, _ int, _ *http.Response) time.Duration { return 0 }
    )
) // use timeout 0, so we can get the response immediately

user, _, userErr := c.Users.GetUser(1, gitlab.GetUsersOptions{})

// user == nil

if userErr == nil {	
    if glErr, ok := userErr.(*gitlab.ErrorResponse); ok {
        fmt.Println(glErr.Message) // == "test error we've defined"
    }
}

```

## Mocking with pagination

```golang
mockedURL := NewMockedHTTPServer(
    WithRequestMatchPages(
        GetApiV4GroupsProjectsById,
        []*gitlab.Project{
            {
                Name: "repo-A-on-first-page",
            },
            {
                Name: "repo-B-on-first-page",
            },
        },
        []*gitlab.Project{
            {
                Name: "repo-C-on-second-page",
            },
            {
                Name: "repo-D-on-second-page",
            },
        },
    ),
)

is := is.New(t)
c, err := gitlab.NewClient(
    "",
    gitlab.WithBaseURL(mockedURL),
    gitlab.WithCustomBackoff(
        func(_, _ time.Duration, _ int, _ *http.Response) time.Duration { return 0 }
    )
)
is.NoErr(err)

opt := &gitlab.ListGroupProjectsOptions{
    ListOptions: gitlab.ListOptions{
        // in fact, the perPage option is ignored
        // but this would be present in production code
        PerPage: 2,
    },
}

var allProjs []*gitlab.Project

for {
    projs, resp, listErr := c.Groups.ListGroupProjects(1, opt)
    is.NoErr(listErr)
    is.Equal(len(projs), 2)

    allProjs = append(allProjs, projs...)

    if resp.NextPage == 0 {
        break
    }

    opt.Page = resp.NextPage
}

// matches the mock definitions len(page[0]) + len(page[1])
// len(allRepos) == 4
```

# Why

I was working on a project that was using GitLab's REST API. And I wanted to write tests for that. And I couldn't. Until now.

# Thanks

Thanks to @migueleliasweb for his go-github-mock project that was the baseline for this one.

# License

This library is distributed under the MIT License found in LICENSE.
