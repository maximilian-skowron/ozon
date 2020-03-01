// Package credential handler for all credential terms
package credential

import "github.com/spf13/cobra"

// CredentialCmd is used to create and manage credentials
// and standard values for the different platforms
var CredentialCmd = &cobra.Command{
	Use:   "cred",
	Short: "cred",
}
