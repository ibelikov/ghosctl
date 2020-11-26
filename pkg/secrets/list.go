package secrets

import (
	"context"
	"log"

	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/ghosctl/pkg/config"
)

// List returns the list of GitHub Org Secrets
func List(config *config.Configuration) *github.Secrets {
	secrets, _, err := config.Client.Actions.ListOrgSecrets(context.Background(), config.Organization, &github.ListOptions{})

	if err != nil {
		log.Fatalf("Failed to list Organization Secrets: %v", err)
	}

	return secrets
}
