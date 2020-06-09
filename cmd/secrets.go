package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	secretsCmd = &cobra.Command{
		Use:     "secrets",
		Short:   "Display secret infomation",
		Long:    "Top level command for displaying secret information in a namespace.",
		Example: `k8ctl secrets --help (for subcommands)`,
	}

	secretsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [SECRET]",
		Short: "Display details of a secret",
		Long:  "Displays details of a secret in a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runSecretsDescribe(name, namespace)
		},
		Example: `k8ctl secrets describe --help
k8ctl secrets describe --cluster nyc --namespace dev myapp-secret-123
k8ctl secrets describe -l nyc -n dev myapp-secret-123`,
	}

	secretsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List secrets",
		Long:  "List will display a list of all secrets in a cluster and namespace.",
		Args:  cobra.MaximumNArgs(0),
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
			return runSecretsList(namespace, format)
		},
		Example: `k8ctl secrets list --help
k8ctl secrets list --cluster nonprod --namespace dev
k8ctl secrets list -l nyc -n dev
k8ctl secrets list -l nyc -n dev --format json
k8ctl secrets list -l nyc -n dev -f yaml`,
	}
)

func init() {
	RootCmd.AddCommand(secretsCmd)
	secretsCmd.AddCommand(secretsSubCmdDescribe)
	secretsCmd.AddCommand(secretsSubCmdList)

	secretsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	secretsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	secretsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	secretsSubCmdDescribe.MarkFlagRequired("namespace")
	secretsSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runSecretsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Secret(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runSecretsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Secrets(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
