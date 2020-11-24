package config

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"

	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/org-secrets-manager/pkg/auth"
)

// Configuration injects the dependencies for CLI actions
type Configuration struct {
	Client       *github.Client
	Organization string
}

// New injects configuration that all actions share
func New() *Configuration {
	privateKey := os.Getenv("GH_APP_PRIVATE_KEY")
	token := os.Getenv("GH_TOKEN")
	client := github.NewClient(nil)

	if privateKey != "" {
		key, err := base64.URLEncoding.DecodeString(privateKey)
		if err != nil {
			log.Fatalf("Can't base64 decode GH_APP_PRIVATE_KEY env var: %v", err)
		}
		client = auth.GetAppClient(
			envInt64("GH_APP_ID"),
			envInt64("GH_APP_INSTALLATION_ID"),
			key,
		)
	} else if token != "" {
		client = auth.GetTokenClient(token)
	} else {
		log.Fatalln("GitHub auth credentials missing." +
			"Either GH_TOKEN or GH_APP_PRIVATE_KEY, GH_APP_ID, GH_APP_INSTALLATION_ID env vars must be set.")
	}
	config := &Configuration{
		Client:       client,
		Organization: os.Getenv("GH_ORG"),
	}

	return config
}

func envInt64(name string) int64 {
	ret, err := strconv.ParseInt(os.Getenv(name), 10, 64)
	if err != nil {
		log.Fatalf("Can't parse %s env var: %v", name, err)
	}
	return ret
}
