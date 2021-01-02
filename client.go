package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

func New(ctx context.Context, conf *oauth2.Config, t *oauth2.Token, f OnTokenExchangedFunc) *http.Client {
	ts := &TokenSource{
		new: conf.TokenSource(ctx, t),
		t:   t,
		f:   f,
	}

	return oauth2.NewClient(ctx, ts)
}

func Endpoint(baseUrl string) oauth2.Endpoint {
	return oauth2.Endpoint{
		TokenURL: fmt.Sprintf("%s/oauth2/access_token", baseUrl),
	}
}
