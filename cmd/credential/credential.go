// Package credential handler for all credential terms
package credential

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CredentialCmd is used to create and manage credentials
// and standard values for the different platforms
var CredentialCmd = &cobra.Command{
	Use:   "cred",
	Short: "Command to manage your credentials and connected platforms.",
	Long: `The credential (cred) command provides functionality to create and delete 
your authentification information and preferences for the project generation.

To get started, create a new authentification:
	ozon cred add ..`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Credential Command.")
	},
}

func init() {
	CredentialCmd.AddCommand(credentialAddCmd)
}
