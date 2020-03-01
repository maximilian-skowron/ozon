// Package create handle functions to create projects
package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CreateCmd command to create multiple projects.
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Projects",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("You need to provide at least a project name.")
		} else if len(args) >= 2 {
			fmt.Println("Too many names defined. For aditional configuration use flags.")
		} else {
			for _, element := range args {
				fmt.Println(element)
			}
		}
	},
}
