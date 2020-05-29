package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	podsCmd = &cobra.Command{
		Use:     "pods",
		Short:   "Display pod infomation",
		Long:    "Top level command for displaying pod information in a namespace.",
		Example: `k8ctl pods --help (for subcommands)`,
	}

	podsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [POD]",
		Short: "Display details of a pod",
		Long:  "Displays details of a pod in anamespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runPodsDescribe(name, namespace)
		},
		Example: `k8ctl pods describe --help
k8ctl pods describe --cluster nyc --namespace dev myapp-pod-123
k8ctl pods describe -l nyc -n dev myapp-pod-123`,
	}

	podsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List pods",
		Long:  "List will display a list of all pods in a cluster and namespace.",
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
			return runPodsList(namespace, format)
		},
		Example: `k8ctl pods list --help
k8ctl pods list --cluster nyc --namespace dev
k8ctl pods list -l nyc -n dev
k8ctl pods list -l nyc -n dev --format json
k8ctl pods list -l nyc -n dev -f yaml`,
	}
)

func init() {
	RootCmd.AddCommand(podsCmd)
	podsCmd.AddCommand(podsSubCmdDescribe)
	podsCmd.AddCommand(podsSubCmdList)

	podsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	podsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	podsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	podsSubCmdDescribe.MarkFlagRequired("namespace")
	podsSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runPodsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Pod(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runPodsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Pods(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
