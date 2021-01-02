package oauth

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

func NewClient(ctx context.Context, conf *oauth2.Config, t *oauth2.Token, f OnTokenExchangedFunc) *http.Client {
	ts := &TokenSource{
		new: conf.TokenSource(ctx, t),
		t:   t,
		f:   f,
	}

	return oauth2.NewClient(ctx, ts)
}
