module github.com/ibelikov/ghosctl

go 1.15

require (
	github.com/GoKillers/libsodium-go v0.0.0-20171022220152-dd733721c3cb
	github.com/bradleyfalzon/ghinstallation v1.1.1
	github.com/ghodss/yaml v1.0.0
	github.com/google/go-github/v32 v32.1.0
	github.com/spf13/cobra v1.1.1
	github.com/variantdev/vals v0.11.0
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
)

replace github.com/variantdev/vals => github.com/tnaroska/vals v0.11.1-0.20201119005136-88fa11462978
