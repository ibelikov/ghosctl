package secrets

import (
	"io/ioutil"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/ibelikov/org-secrets-manager/pkg/config"
)

// Secret represents yaml manifest structure
type Secret struct {
	Name          string  `json:"name"`
	Value         string  `json:"value"`
	SelectedRepos []int64 `json:"repos,omitempty"`
}

// Apply creates/updates the secrets listed in a given .yaml manifest
func Apply(config *config.Configuration, yamlpath string) {
	secrets := LoadConfig(yamlpath)
	for _, secret := range *secrets {
		Create(config, secret.Name, secret.Value, secret.SelectedRepos)
	}
}

// LoadConfig loads config file and returns Config object merged with defaults
func LoadConfig(path string) *[]Secret {
	secrets := &[]Secret{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read secrets manifest from %s: %v", path, err)
	}

	err = yaml.Unmarshal(configFile, secrets)
	if err != nil {
		log.Fatalf("Couldn't process secrets manifest %s: %v", path, err)
	}
	return secrets
}