package twilio

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func TwilioCLI() schema.Executable {
	return schema.Executable{
		Runs:      []string{"twilio"},
		Name:      "Twilio CLI",
		DocsURL:   sdk.URL("https://twilio.com/docs/cli"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Credentials: []schema.CredentialType{
			APIKey(),
		},
	}
}
