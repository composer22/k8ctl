package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/composer22/k8ctl/client"
)

// guideCmd returns extra help to the user
var guideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Usage guide for the application",
	Long:  "Returns extra help to the user including valid namespaces and resource names.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return printGuide()
	},
	Example: `k8ctl guide`,
}

func init() {
	RootCmd.AddCommand(guideCmd)
}

// Prints out the response message with the guide.
func printGuide() error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Guide()
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
