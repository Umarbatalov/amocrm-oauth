package main

import (
	"context"
	"encoding/json"
	"fmt"
	oauth "github.com/Umarbatalov/amocrm-oauth"
	"golang.org/x/oauth2"
	"log"
	"os"
)

func StoreNewToken(t *oauth2.Token) error {
	// persist token
	return nil // or error
}

type AccountAPI struct {
	Id   int
	Name string
}

func main() {
	// $ set -a
	// $ . ./examples/.env
	// $ set +a
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectUrl := os.Getenv("REDIRECT_URL")
	baseUrl := os.Getenv("BASE_URL")

	ctx := context.Background()
	conf := oauth.NewConfig(clientId, clientSecret, redirectUrl, baseUrl)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	src := oauth.NewTokenSource(ctx, conf, tok, StoreNewToken)
	client := oauth2.NewClient(ctx, src)
	_ = client

	res, _ := client.Get(baseUrl + "/api/v4/account")

	account := &AccountAPI{}

	err = json.NewDecoder(res.Body).Decode(account)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Account: %+v", *account))
}
