package mock

import (
	"net/http"

	"github.com/gorilla/mux"
)

// WithRequestMatchHandler implements a request callback
// for the given `pattern`.
//
// For custom implementations, this handler usage is encouraged.
//
// Example:
//
//	WithRequestMatchHandler(
//		GetApiV4Projects,
//		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
//			w.Write(MustMarshal([]gitlab.Project{
//				{
//					Name: "mocked-proj-1",
//				},
//				{
//					Name: "mocked-proj-2",
//				},
//			}))
//		}),
//	)
func WithRequestMatchHandler(
	ep EndpointPattern,
	handler http.Handler,
) MockBackendOption {
	return func(router *mux.Router) {
		router.Handle(ep.Pattern, handler).Methods(ep.Method)
	}
}

// WithRequestMatch implements a simple FIFO for requests
// of the given `pattern`.
//
// Once all responses have been used, it shall panic()!
//
// Example:
//
//	WithRequestMatch(
//		GetApiV4UsersById,
//		gitlab.User{
//			Name: "foobar",
//		},
//	)
func WithRequestMatch(
	ep EndpointPattern,
	responsesFIFO ...interface{},
) MockBackendOption {
	responses := [][]byte{}

	for _, r := range responsesFIFO {
		switch v := r.(type) {
		case []byte:
			responses = append(responses, v)
		default:
			responses = append(responses, MustMarshal(r))
		}
	}

	return WithRequestMatchHandler(ep, &FIFOReponseHandler{
		Responses: responses,
	})
}

// WithRequestMatchPages honors pagination directives.
//
// Pages can be requested in any order and each page can be called multiple times.
//
// E.g.
//
//	mockedURL := NewMockedHTTPServer(
//		WithRequestMatchPages(
//			GetApiV4GroupsProjectsById,
//			[]*gitlab.Project{
//				{
//					Name: "repo-A-on-first-page",
//				},
//				{
//					Name: "repo-B-on-first-page",
//				},
//			},
//			[]*gitlab.Project{
//				{
//					Name: "repo-C-on-second-page",
//				},
//				{
//					Name: "repo-D-on-second-page",
//				},
//			},
//		),
//	)
func WithRequestMatchPages(
	ep EndpointPattern,
	pages ...interface{},
) MockBackendOption {
	p := [][]byte{}

	for _, r := range pages {
		p = append(p, MustMarshal(r))
	}

	return WithRequestMatchHandler(ep, &PaginatedReponseHandler{
		ResponsePages: p,
	})
}
