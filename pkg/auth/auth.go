package auth

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
)

var (
	githubAppID             = os.Getenv("GH_APP_ID")
	githubAppInstallationID = os.Getenv("GH_APP_INSTALLATION_ID")
	encodedRSAKey           = os.Getenv("GH_APP_PRIVATE_KEY")
)

// GetClient performs authorization using GitHub App credentials and returns GitHub Client
func GetClient() *github.Client {
	key, err := base64.URLEncoding.DecodeString(encodedRSAKey)
	if err != nil {
		log.Fatalf("Can't base64 decode GH_APP_PRIVATE_KEY env var: %v", err)
	}

	appID, err := strconv.ParseInt(githubAppID, 10, 64)
	if err != nil {
		log.Fatalf("Can't parse GH_APP_ID nev var: %v", err)
	}

	appInstallationID, err := strconv.ParseInt(githubAppInstallationID, 10, 64)
	if err != nil {
		log.Fatalf("Can't parse GH_APP_INSTALLATION_ID nev var: %v", err)
	}

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID authenticating with installation ID.
	itr, err := ghinstallation.New(tr, appID, appInstallationID, key)
	if err != nil {
		log.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	return client
}
