package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

// GetAppClient performs authorization using GitHub App credentials and returns GitHub Client
func GetAppClient(appID int64, installationID int64, privateKey []byte) *github.Client {
	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID authenticating with installation ID.
	itr, err := ghinstallation.New(tr, appID, installationID, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	return client
}

// GetTokenClient performs authorization using GitHub Personal Token and returns GitHub Client
func GetTokenClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	return client
}
