package config

import (
	"log"
	"os"
	"strconv"
)

// EnvSettings describes all of the environment settings.
type EnvSettings struct {
	// AppID is the GitHub APP ID.
	AppID int64
	// InstallationID is the GitHub App Installation ID.
	InstallationID int64
	// Organization is the name of GitHub Organization.
	Organization string
	// PrivateKey is base64-encoded RSA private key of GitHub App.
	PrivateKey string
}

// New parses env vars to EnvSettings.
func New() *EnvSettings {
	env := &EnvSettings{
		AppID:          envInt64("GH_APP_ID"),
		InstallationID: envInt64("GH_APP_INSTALLATION_ID"),
		Organization:   os.Getenv("GH_ORG"),
		PrivateKey:     os.Getenv("GH_APP_PRIVATE_KEY"),
	}

	return env
}

func envInt64(name string) int64 {
	ret, err := strconv.ParseInt(os.Getenv(name), 10, 64)
	if err != nil {
		log.Fatalf("Can't parse %s env var: %v", name, err)
	}
	return ret
}
