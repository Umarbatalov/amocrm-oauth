package oauth

import (
	"golang.org/x/oauth2"
	"sync"
)

// called when token refreshed
// so new refresh token can be persisted
type OnTokenExchangedFunc func(*oauth2.Token) error

type TokenSource struct {
	new oauth2.TokenSource
	mu  sync.Mutex // guards t
	t   *oauth2.Token
	f   OnTokenExchangedFunc
}

// Token returns the current token if it's still valid, else will
// refresh the current token (using r.Context for HTTP http information)
// and return the new one.
func (s *TokenSource) Token() (*oauth2.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.t.Valid() {
		return s.t, nil
	}
	t, err := s.new.Token()
	if err != nil {
		return nil, err
	}
	s.t = t
	return t, s.f(t)
}
