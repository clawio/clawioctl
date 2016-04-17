package client

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type Config struct {
	ClientID     string
	ClientSecret string
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

func NewClient(config *Config, source TokenStore) *http.Client {
	t := &transport{}
	t.config = config
	t.source = source
	t.base = http.DefaultTransport

	return &http.Client{
		Transport: t,
	}
}

type transport struct {
	config *Config
	source TokenStore
	base   http.RoundTripper
}

// RoundTrip authorizes and authenticates the request with an
// access token. If no token exists or token is expired,
// tries to refresh/fetch a new token.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token := t.source.Get()
	req.Header.Set("Token", token)
	res, err := t.base.RoundTrip(req)
	return res, err
}
