package oauth

import (
	"fmt"
	"golang.org/x/oauth2"
)

func NewConfig(clientId, clientSecret, redirectUrl, accountUrl string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     endpoint(accountUrl),
	}
}

func endpoint(baseUrl string) oauth2.Endpoint {
	return oauth2.Endpoint{
		TokenURL: fmt.Sprintf("%s/oauth2/access_token", baseUrl),
	}
}
