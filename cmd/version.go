package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"

	"github.com/spf13/cobra"
)

// versionCmd returns the version of the application
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of the application",
	Long:  "Returns the version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		result := client.NewClient("", "").Version()
		fmt.Println(result)
	},
	Example: `k8ctl -l nyc version`,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
