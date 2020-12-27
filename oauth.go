package oauth

import (
	"context"
	"fmt"
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
// refresh the current token (using r.Context for HTTP client information)
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

func NewTokenSource(ctx context.Context, conf *oauth2.Config, t *oauth2.Token, f OnTokenExchangedFunc) *TokenSource {
	return &TokenSource{
		new: conf.TokenSource(ctx, t),
		t:   t,
		f:   f,
	}
}

// clientId - integration id
// clientSecret - secret key
// baseUrl - account url, example: https://rkachko.amocrm.ru
func NewConfig(clientId, clientSecret, redirectUrl, baseUrl string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s/oauth2/access_token", baseUrl),
		},
	}
}
