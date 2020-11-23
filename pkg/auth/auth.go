package auth

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/org-secrets-manager/pkg/config"
)

// GetClient performs authorization using GitHub App credentials and returns GitHub Client
func GetClient(config *config.EnvSettings) *github.Client {
	key, err := base64.URLEncoding.DecodeString(config.PrivateKey)
	if err != nil {
		log.Fatalf("Can't base64 decode GH_APP_PRIVATE_KEY env var: %v", err)
	}

	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID authenticating with installation ID.
	itr, err := ghinstallation.New(tr, config.AppID, config.InstallationID, key)
	if err != nil {
		log.Fatal(err)
	}

	// Use installation transport with github.com/google/go-github
	client := github.NewClient(&http.Client{Transport: itr})

	return client
}
