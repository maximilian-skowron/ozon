// Package version handler for all version terms
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCmd provides the command to show the current version
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of ozon.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ozon version: 0.1.0")
	},
}
