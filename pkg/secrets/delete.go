package secrets

import (
	"context"
	"log"

	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/ghosctl/pkg/config"
)

// Delete GitHub Org Secret
func Delete(config *config.Configuration, name string) *github.Response {
	resp, err := config.Client.Actions.DeleteOrgSecret(context.Background(), config.Organization, name)
	if err != nil {
		log.Fatalf("Failed to delete Organization Secret %s: %v", name, err)
	}

	return resp
}
