package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	jobsCmd = &cobra.Command{
		Use:     "jobs",
		Short:   "Display job infomation",
		Long:    "Top level command for displaying job information in a namespace.",
		Example: `k8ctl jobs --help (for subcommands)`,
	}

	jobsSubCmdDescribe = &cobra.Command{
		Use:   "describe [flags] [JOB]",
		Short: "Display details of a job",
		Long:  "Displays details of a job in the cluster and namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			return runJobsDescribe(name, namespace)
		},
		Example: `k8ctl jobs describe --help
k8ctl jobs describe --cluster nyc --namespace dev myapp-job
k8ctl jobs describe -l nyc -n dev myapp-job`,
	}

	jobsSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List jobs",
		Long:  "List will display a list of all jobs in a namespace.",
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
			return runJobsList(namespace, format)
		},
		Example: `k8ctl jobs list --help
k8ctl jobs list --cluster nyc --namespace dev
k8ctl jobs list -l nyc -n dev
k8ctl jobs list -l nyc -n dev --format json
k8ctl jobs list -l nyc -n dev -f yaml`,
	}
)

func init() {
	RootCmd.AddCommand(jobsCmd)
	jobsCmd.AddCommand(jobsSubCmdDescribe)
	jobsCmd.AddCommand(jobsSubCmdList)

	jobsSubCmdDescribe.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")
	jobsSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	jobsSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to report. (required)")

	jobsSubCmdDescribe.MarkFlagRequired("namespace")
	jobsSubCmdList.MarkFlagRequired("namespace")
}

// Support functions to conduct the client call.

func runJobsDescribe(name string, namespace string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Job(name, namespace)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runJobsList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Jobs(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
