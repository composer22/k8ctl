package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	ingressesCmd = &cobra.Command{
		Use:     "ingresses",
		Short:   "Display ingress infomation",
		Long:    "Top level command for displaying ingress information in a namespace.",
		Example: `k8ctl ingresses --help (for subcommands)`,
	}

	ingressesSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [INGRESS]",
		Short: "Display details of a ingress",
		Long:  "Displays details of a ingress in a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runIngressesDescribe(name, namespace)
		},
		Example: `k8ctl ingresses describe --help
k8ctl ingresses describe --cluster nyc --namespace dev myapp-ingress
k8ctl ingresses describe -l nyc -n dev myapp-ingress`,
	}

	ingressesSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List ingresses",
		Long:  "List will display a list of all ingresses in a namespace.",
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
			return runIngressesList(namespace, format)
		},
		Example: `k8ctl ingresses list --help
k8ctl ingresses list --cluster nyc --namespace dev
k8ctl ingresses list -l nyc -n dev
k8ctl ingresses list -l nyc -n dev --format json
k8ctl ingresses list -l nyc -n dev -f yaml`,
	}
)

func init() {
	RootCmd.AddCommand(ingressesCmd)
	ingressesCmd.AddCommand(ingressesSubCmdDescribe)
	ingressesCmd.AddCommand(ingressesSubCmdList)

	ingressesSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	ingressesSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	ingressesSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	ingressesSubCmdDescribe.MarkFlagRequired("namespace")
	ingressesSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runIngressesDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Ingress(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runIngressesList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Ingresses(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
