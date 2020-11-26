package secrets

import (
	"context"
	"log"

	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/ghosctl/pkg/config"
)

// Get returns GitHub Org Secret
func Get(config *config.Configuration, name string) *github.Secret {
	secret, resp, err := config.Client.Actions.GetOrgSecret(context.Background(), config.Organization, name)

	if err != nil {
		log.Printf("%v", resp.Request.Body)
		log.Fatalf("Failed to get Organization Secret: %v", err)
	}

	return secret
}
