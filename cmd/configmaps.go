package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	configmapsCmd = &cobra.Command{
		Use:     "configmaps",
		Short:   "Display configmap infomation",
		Long:    "Top level command for displaying configmap information in a namespace.",
		Example: `k8ctl configmaps --help (for subcommands)`,
	}

	configmapsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [CONFIGMAP-NAME]",
		Short: "Display details of a configmap",
		Long:  "Displays details of a configmap in a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runConfigmapsDescribe(name, namespace)
		},
		Example: `k8ctl configmaps describe --help
k8ctl configmaps describe --cluster nyc --namespace dev myapp-configmap
k8ctl configmaps describe -l nyc -n dev myapp-configmap`,
	}

	configmapsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List configmaps",
		Long:  "List will display a list of all configmaps in a namespace.",
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
			return runConfigmapsList(namespace, format)
		},
		Example: `k8ctl configmaps list --help
k8ctl configmaps list --cluster nyc --namespace dev
k8ctl configmaps list -l nyc -n dev`,
	}
)

func init() {
	RootCmd.AddCommand(configmapsCmd)
	configmapsCmd.AddCommand(configmapsSubCmdDescribe)
	configmapsCmd.AddCommand(configmapsSubCmdList)

	configmapsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	configmapsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	configmapsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	configmapsSubCmdDescribe.MarkFlagRequired("namespace")
	configmapsSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runConfigmapsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Configmap(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runConfigmapsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Configmaps(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
