package secrets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v32/github"
)

// List returns the list of GitHub Org Secrets
func List(client *github.Client, org string) {
	secrets, _, _ := client.Actions.ListOrgSecrets(context.Background(), org, &github.ListOptions{})

	for _, secret := range secrets.Secrets {
		prettyOutput, _ := json.MarshalIndent(secret, "", "  ")
		fmt.Printf("%s\n", string(prettyOutput))
	}
}
