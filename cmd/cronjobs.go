package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	cronjobsCmd = &cobra.Command{
		Use:     "cronjobs",
		Short:   "Display cronjob infomation",
		Long:    "Top level command for displaying cronjob information in a namespace.",
		Example: `k8ctl cronjobs --help (for subcommands)`,
	}

	cronjobsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [CRONJOBNAME]",
		Short: "Display details of a cronjob",
		Long:  "Displays details of a cronjob in a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runCronjobsDescribe(name, namespace)
		},
		Example: `k8ctl cronjobs describe --help
k8ctl cronjobs describe --cluster nyc --namespace dev myapp-cronjob
k8ctl cronjobs describe -l nyc -n dev myapp-cronjob`,
	}

	cronjobsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List cronjobs",
		Long:  "List will display a list of all cronjobs in a namespace.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var namespace, format string
			var err error
			if namespace, err = cmd.Flags().GetString("namespace"); err != nil {
				return err
			}
			format, err = cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			return runCronjobsList(namespace, format)
		},
		Example: `k8ctl cronjobs list --help
k8ctl cronjobs list --cluster nyc --namespace dev
k8ctl cronjobs list -l nyc -n dev`,
	}
)

func init() {
	RootCmd.AddCommand(cronjobsCmd)
	cronjobsCmd.AddCommand(cronjobsSubCmdDescribe)
	cronjobsCmd.AddCommand(cronjobsSubCmdList)

	cronjobsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	cronjobsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	cronjobsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	cronjobsSubCmdDescribe.MarkFlagRequired("namespace")
	cronjobsSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runCronjobsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Cronjob(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runCronjobsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Cronjobs(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
