package secrets

import (
	"io/ioutil"
	"log"

	yaml "github.com/ghodss/yaml"
	"github.com/ibelikov/ghosctl/pkg/config"
	"github.com/variantdev/vals"
)

// Manifest represents yaml manifest structure
type Manifest struct {
	Secrets []Secret `json:"secrets"`
}

// Secret represents the structure of secret block
type Secret struct {
	Name          string   `json:"name"`
	Value         string   `json:"value"`
	SelectedRepos []string `json:"repos,omitempty"`
}

// Apply creates/updates the secrets listed in a given .yaml manifest
func Apply(config *config.Configuration, yamlpath string) {
	runtime, _ := vals.New(vals.Options{})

	manifest := LoadConfig(yamlpath)

	for _, secret := range manifest.Secrets {
		renderedSecret, err := runtime.Eval(map[string]interface{}{
			"value": secret.Value,
		})
		if err != nil {
			log.Fatalf("Error interpolating secret references in manifest: %v", err)
		}
		Create(config, secret.Name, renderedSecret["value"].(string), secret.SelectedRepos)
	}
}

// LoadConfig loads config file and returns Config object merged with defaults
func LoadConfig(path string) *Manifest {
	manifest := &Manifest{}
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read secrets manifest from %s: %v", path, err)
	}

	err = yaml.Unmarshal(configFile, manifest)
	if err != nil {
		log.Fatalf("Couldn't process secrets manifest %s: %v", path, err)
	}
	return manifest
}
