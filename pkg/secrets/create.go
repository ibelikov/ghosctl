package secrets

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	sodium "github.com/GoKillers/libsodium-go/cryptobox"
	"github.com/google/go-github/v32/github"
	"github.com/ibelikov/ghosctl/pkg/config"
)

// Create org secret with given name and value
func Create(config *config.Configuration, name string, value string, repos []string) *github.Response {
	key, _, err := config.Client.Actions.GetOrgPublicKey(context.Background(), config.Organization)
	if err != nil {
		log.Fatalf("Can't get Org public key: %v", err)
	}
	encryptedSecret, err := encryptSecretWithPublicKey(key, name, value)
	if err != nil {
		log.Fatalf("Can't create encrypted secret: %v", err)
	}

	// Change visibility and pass selected repos list if needed
	if len(repos) > 0 {
		repoIDs := make([]int64, len(repos))
		for i, repoName := range repos {
			repo, _, err := config.Client.Repositories.Get(context.Background(), config.Organization, repoName)
			if err != nil {
				log.Fatalf("Can't get repository ID for %s: %v", repoName, err)
			}
			repoIDs[i] = *repo.ID
		}
		encryptedSecret.Visibility = "selected"
		encryptedSecret.SelectedRepositoryIDs = repoIDs
	}

	resp, err := config.Client.Actions.CreateOrUpdateOrgSecret(context.Background(), config.Organization, encryptedSecret)
	if err != nil {
		log.Fatalf("Can't create Org secret: %v", err)
	}

	return resp
}

func encryptSecretWithPublicKey(publicKey *github.PublicKey, secretName string, secretValue string) (*github.EncryptedSecret, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey.GetKey())
	if err != nil {
		return nil, fmt.Errorf("base64.StdEncoding.DecodeString was unable to decode public key: %v", err)
	}

	secretBytes := []byte(secretValue)
	encryptedBytes, exit := sodium.CryptoBoxSeal(secretBytes, decodedPublicKey)
	if exit != 0 {
		return nil, errors.New("sodium.CryptoBoxSeal exited with non zero exit code")
	}

	encryptedString := base64.StdEncoding.EncodeToString(encryptedBytes)
	keyID := publicKey.GetKeyID()
	encryptedSecret := &github.EncryptedSecret{
		Name:  secretName,
		KeyID: keyID,
		// Set default visibility to "private", we're not targeting public projects here
		Visibility:     "private",
		EncryptedValue: encryptedString,
	}
	return encryptedSecret, nil
}
