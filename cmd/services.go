package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	servicesCmd = &cobra.Command{
		Use:     "services",
		Short:   "Display service infomation",
		Long:    "Top level command for displaying service information in a namespace.",
		Example: `k8ctl services --help (for subcommands)`,
	}

	servicesSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [SERVICE]",
		Short: "Display details of a service",
		Long:  "Displays details of a service in anamespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runServicesDescribe(name, namespace)
		},
		Example: `k8ctl services describe --help
k8ctl services describe --cluster nyc --namespace dev myapp-service-123
k8ctl services describe -l nyc -n dev myapp-service-123`,
	}

	servicesSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List services",
		Long:  "List will display a list of all services in a cluster and namespace.",
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
			return runServicesList(namespace, format)
		},
		Example: `k8ctl services list --help
k8ctl services list --cluster nonprod --namespace dev
k8ctl services list -l nyc -n dev
k8ctl services list -l nyc -n dev --format json
k8ctl services list -l nyc -n dev -f yaml`,
	}
)

func init() {
	RootCmd.AddCommand(servicesCmd)
	servicesCmd.AddCommand(servicesSubCmdDescribe)
	servicesCmd.AddCommand(servicesSubCmdList)

	servicesSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	servicesSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	servicesSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	servicesSubCmdDescribe.MarkFlagRequired("namespace")
	servicesSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runServicesDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Service(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runServicesList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Services(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
