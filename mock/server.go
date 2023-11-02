package mock

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

// EndpointPattern models the Gitlab's API endpoints
type EndpointPattern struct {
	Pattern string // eg. "/repos/{owner}/{repo}/actions/artifacts"
	Method  string // "GET", "POST", "PATCH", etc
}

// MockBackendOption is used to configure the *mux.router
// for the mocked backend
type MockBackendOption func(*mux.Router)

// FIFOReponseHandler handler implementation that
// responds to the HTTP requests following a FIFO approach.
//
// Once all available `Responses` have been used, this handler will panic()!
type FIFOReponseHandler struct {
	lock         sync.Mutex
	Responses    [][]byte
	CurrentIndex int
}

// ServeHTTP implementation of `http.Handler`
func (srh *FIFOReponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srh.lock.Lock()
	defer srh.lock.Unlock()

	if srh.CurrentIndex > len(srh.Responses) {
		panic(fmt.Sprintf(
			"go-gitlab-mock: no more mocks available for %s",
			r.URL.Path,
		))
	}

	defer func() {
		srh.CurrentIndex++
	}()

	w.Write(srh.Responses[srh.CurrentIndex])
}

// PaginatedReponseHandler handler implementation that
// responds to the HTTP requests and honors the pagination headers
//
//	Header e.g: `Link: <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=15>; rel="next",
//	 <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=34>; rel="last",
//	 <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=1>; rel="first",
//	 <https://api.github.com/search/code?q=addClass+user%3Amozilla&page=13>; rel="prev"`
//
// See: https://docs.github.com/en/rest/guides/traversing-with-pagination
type PaginatedReponseHandler struct {
	ResponsePages [][]byte
}

func (prh *PaginatedReponseHandler) getCurrentPage(r *http.Request) int {
	strPage := r.URL.Query().Get("page")

	if strPage == "" {
		return 1
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err == nil {
		return page
	}

	// this should never happen as the request is being made by the SDK
	panic(fmt.Sprintf("invalid page: %s", strPage))
}

func (prh *PaginatedReponseHandler) getPerPage(r *http.Request) int {
	strPage := r.URL.Query().Get("per_page")

	if strPage == "" {
		return 1
	}

	page, err := strconv.Atoi(r.URL.Query().Get("per_page"))

	if err == nil {
		return page
	}

	// this should never happen as the request is being made by the SDK
	panic(fmt.Sprintf("invalid per_page: %s", strPage))
}

func (prh *PaginatedReponseHandler) generateLinkHeader(
	w http.ResponseWriter,
	r *http.Request,
) {
	currentPage := prh.getCurrentPage(r)
	lastPage := len(prh.ResponsePages)

	buf := bytes.NewBufferString(`<?page=1>; rel="first",`)
	buf.WriteString(fmt.Sprintf(`<?page=%d>; rel="last",`, lastPage))

	if currentPage < lastPage {
		// when resp.NextPage == 0, it means no more pages
		// which is basically as not setting it in the response
		buf.WriteString(fmt.Sprintf(`<?page=%d>; rel="next",`, currentPage+1))
	}

	if currentPage > 1 {
		buf.WriteString(fmt.Sprintf(`<?page=%d>; rel="prev",`, currentPage-1))
	}

	w.Header().Add("Link", buf.String())
}

func (prh *PaginatedReponseHandler) generatePageHeaders(
	w http.ResponseWriter,
	r *http.Request,
) {
	currentPage := prh.getCurrentPage(r)
	lastPage := len(prh.ResponsePages)
	perPage := prh.getPerPage(r)

	if currentPage < lastPage {
		w.Header().Add("x-next-page", fmt.Sprintf("%d", currentPage+1))
	}

	if currentPage > 1 {
		w.Header().Add("x-prev-page", fmt.Sprintf("%d", currentPage-1))
	}
	w.Header().Add("x-page", fmt.Sprintf("%d", currentPage))
	w.Header().Add("x-per-page", fmt.Sprintf("%d", perPage))
	w.Header().Add("x-total-pages", fmt.Sprintf("%d", lastPage))
	// We don't set `x-total` header, because we don't want to convert
	// byte array to string just so we could count the items.
	// If needed, it can be don though at the expense of a few more
	// type conversions.
}

// ServeHTTP implementation of `http.Handler`
func (prh *PaginatedReponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prh.generateLinkHeader(w, r)
	prh.generatePageHeaders(w, r)
	w.Write(prh.ResponsePages[prh.getCurrentPage(r)-1])
}

// EnforceHostRoundTripper rewrites all requests with the given `Host`.
type EnforceHostRoundTripper struct {
	Host                 string
	UpstreamRoundTripper http.RoundTripper
}

// RoundTrip implementation of `http.RoundTripper`
func (efrt *EnforceHostRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	splitHost := strings.Split(efrt.Host, "://")
	r.URL.Scheme = splitHost[0]
	r.URL.Host = splitHost[1]

	return efrt.UpstreamRoundTripper.RoundTrip(r)
}

// NewMockedHTTPServer creates and configures a mocked Gitlab backend API
// returning url string to pass to gitlab's client constructor.
//
// Example:
//
//	url := NewMockedHTTPServer(
//		WithRequestMatch(
//			GetApiV4UsersById,
//			&gitlab.User{
//				ID:   5,
//				Name: "foobar",
//			},
//		),
//		WithRequestMatch(
//			GetApiV4Groups,
//			[]*gitlab.Group{
//				{
//					Name: "foobar123thisorgwasmocked",
//				},
//			},
//		),
//		WithRequestMatchHandler(
//			GetApiV4Projects,
//			http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
//				w.Write(MustMarshal([]gitlab.Project{
//					{
//						Name: "mocked-proj-1",
//					},
//					{
//						Name: "mocked-proj-2",
//					},
//				}))
//			}),
//		),
//	)
//	c, err := gitlab.NewClient("", gitlab.WithBaseURL(url))
func NewMockedHTTPServer(options ...MockBackendOption) string {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteError(
			w,
			http.StatusNotFound,
			fmt.Sprintf("mock response not found for %s", r.URL.Path),
		)
	})

	for _, o := range options {
		o(router)
	}

	mockServer := httptest.NewServer(router)

	return mockServer.URL
}
