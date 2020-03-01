// Package ozon root command for the cobra cli
package ozon

import (
	"github.com/Ozon-Project/ozon/cmd/create"
	"github.com/Ozon-Project/ozon/cmd/credential"
	"github.com/Ozon-Project/ozon/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ozon",
	Short: "Ozon will generate multiple projects across your variouse softwares.",
	Long: `Ozon will generate multiple projects across your variouse softwares.
It will utilize your confiugred credentials.
			
To get startet create credentials for supportet platforms with:
	ozon cred ...`,
}

// Execute executes the root cobra command ozon
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(credential.CredentialCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(create.CreateCmd)
}
