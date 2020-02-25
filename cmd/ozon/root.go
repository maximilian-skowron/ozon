// Package ozon
// root command for the cobra cli
package ozon

import( 
	"github.com/spf13/cobra"
	"github.com/Ozon-Project/ozon/cmd/credential"
)

var rootCmd = &cobra.Command{
	Use: "ozon",
	Short: "Ozon short",
	Long: `Ozon long`,
}

// Executes executes the root cobra command ozon
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(credential.CredentialCmd)
}