package client

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/clawio/sdk"
)

type Credentials struct {
	AuthenticationServiceBaseURL string
	ClientID                     string
	ClientSecret                 string
}

type TokenStore interface {
	Get() string
	Set(token string) error
}

func NewFileTokenStore(file string) TokenStore {
	return &fileTokenStore{fn: file}
}

type fileTokenStore struct {
	mu sync.Mutex
	fn string
}

func (s *fileTokenStore) Get() string {
	data, _ := ioutil.ReadFile(s.fn)
	return string(data)
}
func (s *fileTokenStore) Set(token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return ioutil.WriteFile(s.fn, []byte(token), 0600)
}

func NewClientWithAuth(config *Credentials, source TokenStore) *http.Client {
	t := &transport{}
	t.config = config
	t.source = source
	t.base = http.DefaultTransport

	return &http.Client{
		Transport: t,
	}
}

type transport struct {
	config *Credentials
	source TokenStore
	base   http.RoundTripper
}

// RoundTrip authorizes and authenticates the request with an
// access token. If no token exists or token is expired,
// tries to refresh/fetch a new token.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// clone request to re-send it if auth fails.
	reqCopy := cloneRequest(req)
	token := t.source.Get()
	req.Header.Set("Authorization", "Bearer "+token)
	res, err := t.base.RoundTrip(req)
	if err != nil {
		return res, err
	}
	// if res.StatusCode == 401 we get a new token and repeat the request
	if res.StatusCode == http.StatusUnauthorized {
		s := sdk.New(&sdk.ServiceEndpoints{AuthServiceBaseURL: t.config.AuthenticationServiceBaseURL}, nil)
		if err != nil {
			return res, err
		}
		token, _, err := s.Auth.Token(t.config.ClientID, t.config.ClientSecret)
		if err != nil {
			return res, err
		}
		t.source.Set(token)
		reqCopy.Header.Set("Authorization", "Bearer "+token)
		return t.base.RoundTrip(reqCopy)
	}
	return res, err
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
