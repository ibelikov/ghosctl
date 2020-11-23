package config

import (
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
	client := auth.GetAppClient(
		envInt64("GH_APP_ID"),
		envInt64("GH_APP_INSTALLATION_ID"),
		os.Getenv("GH_APP_PRIVATE_KEY"),
	)

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
