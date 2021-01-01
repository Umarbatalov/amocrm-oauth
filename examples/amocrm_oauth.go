package main

import (
	"context"
	"encoding/json"
	"fmt"
	oauth "github.com/Umarbatalov/amocrm-oauth"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"os"
)

var tokenFile = "token.json"

type Account struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getTokenFromStore() (*oauth2.Token, error) {
	f, _ := ioutil.ReadFile(tokenFile)
	token := &oauth2.Token{}

	err := json.Unmarshal(f, token)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func storeNewToken(t *oauth2.Token) error {
	j, _ := json.Marshal(t)

	return ioutil.WriteFile(tokenFile, j, 0644)
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

	token, err := getTokenFromStore()

	if err != nil {
		log.Println("Token from file not found, try create")

		var authCode string
		if _, err = fmt.Scan(&authCode); err != nil {
			log.Fatal(err)
		}

		// Exchange token with auth_code
		token, err = conf.Exchange(ctx, authCode)

		if err != nil {
			log.Fatal(err)
		}

		err = storeNewToken(token)

		if err != nil {
			log.Fatal(err)
		}
	}

	src := oauth.NewTokenSource(ctx, conf, token, storeNewToken)
	client := oauth2.NewClient(ctx, src)

	res, err := client.Get(baseUrl + "/api/v4/account")

	if err != nil {
		log.Fatal(err)
	}

	account := &Account{}

	err = json.NewDecoder(res.Body).Decode(account)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("Account: %+v", *account))
}
