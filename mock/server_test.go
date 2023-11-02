package mock

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/xanzy/go-gitlab"
)

func TestNewMockedHTTPClient(t *testing.T) {
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
	is := is.New(t)
	c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL))
	is.NoErr(err)

	t.Run("GetUser", func(t *testing.T) {
		user, _, userErr := c.Users.GetUser(1, gitlab.GetUsersOptions{})
		is.NoErr(userErr)
		is.True(user != nil)
		is.True(user.Name != "")
		is.True(user.Name == "foobar")
	})

	t.Run("ListGroups", func(t *testing.T) {
		groups, _, groupsErr := c.Groups.ListGroups(&gitlab.ListGroupsOptions{})
		is.NoErr(groupsErr)
		is.True(len(groups) == 1)
		is.Equal(groups[0].Name, "foobar123thisorgwasmocked")
	})

	t.Run("ListProjects", func(t *testing.T) {
		projs, _, projsErr := c.Projects.ListProjects(&gitlab.ListProjectsOptions{})
		is.NoErr(projsErr)
		is.True(len(projs) == 2)
		is.Equal(projs[0].Name, "mocked-proj-1")
	})
}

func TestMockErrorSimple(t *testing.T) {
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
	is := is.New(t)
	c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL), gitlab.WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration { return 0 }))
	is.NoErr(err)

	user, _, userErr := c.Users.GetUser(1, gitlab.GetUsersOptions{})
	is.True(userErr != nil)
	is.Equal(user, nil)

	glErr, ok := userErr.(*gitlab.ErrorResponse)
	is.True(ok)
	is.Equal(glErr.Message, "test error we've defined")
}

func TestMocksNotConfiguredError(t *testing.T) {
	mockedURL := NewMockedHTTPServer(
		WithRequestMatch(
			GetApiV4UsersById,
			gitlab.User{
				Name: "foobar",
			},
		),
	)
	is := is.New(t)
	c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL), gitlab.WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration { return 0 }))
	is.NoErr(err)

	t.Run("GetUser", func(t *testing.T) {
		user, _, userErr := c.Users.GetUser(1, gitlab.GetUsersOptions{})
		is.NoErr(userErr)
		is.True(user != nil)
		is.Equal(user.Name, "foobar")
	})

	t.Run("ListGroups", func(t *testing.T) {
		groups, _, groupsErr := c.Groups.ListGroups(&gitlab.ListGroupsOptions{})
		is.True(groupsErr != nil)
		is.True(len(groups) == 0)

		r, ok := groupsErr.(*gitlab.ErrorResponse)
		is.True(ok)
		is.True(strings.Contains(r.Message, "mock response not found for"))
	})
}

func TestMocksPaginationAllPages(t *testing.T) {
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
	c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL), gitlab.WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration { return 0 }))
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
		// matches mock definition
		is.Equal(len(projs), 2)

		allProjs = append(allProjs, projs...)

		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}
	is.Equal(len(allProjs), 4)
}

func TestEmptyArrayResult(t *testing.T) {
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
	is := is.New(t)
	c, err := gitlab.NewClient("", gitlab.WithBaseURL(mockedURL))
	is.NoErr(err)

	t.Run("has result", func(t *testing.T) {
		issues1, _, issues1Err := c.Issues.ListProjectIssues(1, &gitlab.ListProjectIssuesOptions{})
		is.NoErr(issues1Err)
		is.True(len(issues1) == 2)
		is.Equal(issues1[0].Title, "Issue 1")
		is.Equal(issues1[1].Title, "Issue 2")
	})
	t.Run("empty result", func(t *testing.T) {
		issues2, _, issues2Err := c.Issues.ListProjectIssues(2, &gitlab.ListProjectIssuesOptions{})
		is.NoErr(issues2Err)
		is.True(len(issues2) == 0)
	})
}
