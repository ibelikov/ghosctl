package secrets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/org-secrets-manager/pkg/auth"
	"github.com/ibelikov/org-secrets-manager/pkg/config"
)

// List returns the list of GitHub Org Secrets
func List() {
	config := config.New()
	client := auth.GetClient(config)
	secrets, _, _ := client.Actions.ListOrgSecrets(context.Background(), config.Organization, &github.ListOptions{})

	for _, secret := range secrets.Secrets {
		prettyOutput, _ := json.MarshalIndent(secret, "", "  ")
		fmt.Printf("%s\n", string(prettyOutput))
	}
}
