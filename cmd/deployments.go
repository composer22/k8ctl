package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	deploymentsCmd = &cobra.Command{
		Use:     "deployments",
		Short:   "Display and restart deployments",
		Long:    "Top level command for displaying or restarting a deployment in a namespace.",
		Example: `k8ctl deployments --help (for subcommands)`,
	}

	deploymentsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [DEPLOYMENT]",
		Short: "Display details of a deployment",
		Long:  "Displays details of a deployment in a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runDeploymentsDescribe(name, namespace)
		},
		Example: `k8ctl deployments describe --help
k8ctl deployments describe --cluster nyc --namespace dev myapp-deployment
k8ctl deployments describe -l nyc -n dev myapp-deployment`,
	}

	deploymentsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List deployments",
		Long:  "List will display a list of all deployments in a namespace.",
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
			return runDeploymentsList(namespace, format)
		},
		Example: `k8ctl deployments list --help
k8ctl deployments list --cluster nyc --namespace dev
k8ctl deployments list -l nyc -n dev
k8ctl deployments list -l nyc -n dev --format json
k8ctl deployments list -l nyc -n dev -f yaml`,
	}

	deploymentsSubCmdRestart = &cobra.Command{
		Use:   "restart [flags] [DEPLOYMENT]",
		Short: "Restart pods under a deployment",
		Long:  "Restart will restart all pods under a deployment in a namespace.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runDeploymentsRestart(name, namespace)
		},
		Example: `k8ctl deployments restart --help
k8ctl deployments restart --cluster nyc --namespace dev myapp-deployment
k8ctl deployments restart -l nyc -n dev myapp-deployment`,
	}
)

func init() {
	RootCmd.AddCommand(deploymentsCmd)
	deploymentsCmd.AddCommand(deploymentsSubCmdDescribe)
	deploymentsCmd.AddCommand(deploymentsSubCmdList)
	deploymentsCmd.AddCommand(deploymentsSubCmdRestart)

	deploymentsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	deploymentsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	deploymentsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")

	deploymentsSubCmdRestart.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	deploymentsSubCmdDescribe.MarkFlagRequired("namespace")
	deploymentsSubCmdList.MarkFlagRequired("namespace")
	deploymentsSubCmdRestart.MarkFlagRequired("namespace")
}

func runDeploymentsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Deployment(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runDeploymentsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Deployments(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runDeploymentsRestart(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.DeploymentRestart(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
