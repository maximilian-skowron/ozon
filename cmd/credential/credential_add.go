package credential

import "github.com/spf13/cobra"

// CredentialCmd is used to create and manage credentials
// and standard values for the different platforms
var credentialAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Command to add authentification details.",
}
