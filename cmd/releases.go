package cmd

import (
	"fmt"

	"github.com/composer22/k8ctl/client"
	"github.com/spf13/cobra"
)

var (
	releasesCmd = &cobra.Command{
		Use:     "releases",
		Short:   "Display and manage helm releases",
		Long:    "Top level command for displaying and managing helm releases",
		Example: `k8ctl releases --help (for subcommands)`,
	}

	releasesSubCmdDelete = &cobra.Command{
		Use:   "delete [flags] [RELEASE]",
		Short: "Delete a release",
		Long:  "Deletes a helm release from the cluster.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			release := args[0]
			return runDelete(release)
		},
		Example: `k8ctl releases delete --help
k8ctl releases delete --cluster nyc myapp-dev
k8ctl releases delete -l nyc myapp-dev`,
	}

	releasesSubCmdDeploy = &cobra.Command{
		Use:   "deploy [flags] [CHART]",
		Short: "Deploy or refreshes a release",
		Long:  "Deploy will apply a helm chart into a namespace.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			release := args[0]
			var tag, namespace, memo string
			var err error
			if tag, err = cmd.Flags().GetString("tag"); err != nil {
				return err
			}
			if namespace, err = cmd.Flags().GetString("namespace"); err != nil {
				return err
			}
			if memo, err = cmd.Flags().GetString("memo"); err != nil {
				return err
			}

			return runDeploy(release, tag, namespace, memo)
		},
		Example: `k8ctl releases deploy --help
k8ctl releases deploy --cluster nyc --namespace dev --tag k8-1.0.0-1234 myapp-service
k8ctl releases deploy -l nyc -n dev -t k8-1.0.0-1234 --memo "a really good bug!" myapp-service`,
	}

	releasesSubCmdHistory = &cobra.Command{
		Use:   "history [flags] [RELEASE]",
		Short: "Display release history",
		Long:  "Displays the history of a release including previous releases and failed deploys.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			release := args[0]
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			return runHistory(release, format)
		},
		Example: `k8ctl releases history --help
k8ctl releases history --cluster nyc my-release-dev
k8ctl releases history -l nyc --format json my-release-dev
k8ctl releases history -l nyc -f yaml my-release-dev`,
	}

	releasesSubCmdList = &cobra.Command{
		Use:   "list [flags]",
		Short: "List releases",
		Long:  "List will display a list of all releases in a namespace.",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			var format string
			format, err = cmd.Flags().GetString("format")
			if err != nil {
				return err
			}

			return runList(namespace, format)
		},
		Example: `k8ctl release list --help
k8ctl release list --cluster nyc --namespace dev
k8ctl release list -l nyc -n dev --format json
k8ctl release list -l nyc -n dev -f yaml`,
	}

	releasesSubCmdRollback = &cobra.Command{
		Use:   "rollback [flags] [RELEASE]",
		Short: "Rollback a release",
		Long:  "Rollbacks a release to the previous revision or one of your choosing.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			release := args[0]
			revision, err := cmd.Flags().GetString("revision")
			if err != nil {
				return err
			}
			return runRollback(release, revision)
		},
		Example: `k8ctl releases rollback --help
k8ctl releases rollback --cluster nyc my-release-dev-003
k8ctl releases rollback -c nyc --revision my-release-dev-001 my-release-dev-003
k8ctl releases rollback -c nyc -r my-release-dev-001 my-release-dev-003`,
	}

	releasesSubCmdStatus = &cobra.Command{
		Use:   "status [flags] [RELEASE]",
		Short: "Display the status of a release",
		Long:  "Displays the status of a release, including deployments, services, ingresses etc.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			release := args[0]
			format, err := cmd.Flags().GetString("format")
			if err != nil {
				return err
			}
			return runStatus(release, format)
		},
		Example: `k8ctl releases status --help
k8ctl releases status --cluster nyc my-release-dev
k8ctl releases status -l nyc my-release-dev
k8ctl releases status -l nyc --format json my-release-dev
k8ctl releases status -l nyc -f yaml my-release-dev`,
	}
)

func init() {
	RootCmd.AddCommand(releasesCmd)
	releasesCmd.AddCommand(releasesSubCmdDelete)
	releasesCmd.AddCommand(releasesSubCmdDeploy)
	releasesCmd.AddCommand(releasesSubCmdHistory)
	releasesCmd.AddCommand(releasesSubCmdList)
	releasesCmd.AddCommand(releasesSubCmdRollback)
	releasesCmd.AddCommand(releasesSubCmdStatus)

	releasesSubCmdDeploy.Flags().StringP("tag", "t", "", "Docker image tag (required)")
	releasesSubCmdDeploy.Flags().StringP("namespace", "n", "", "Namespace to deploy to: dev, qa etc. (required)")
	releasesSubCmdDeploy.Flags().StringP("memo", "m", "", "Information to display in slack etc. (optional)")
	releasesSubCmdDeploy.MarkFlagRequired("tag")
	releasesSubCmdDeploy.MarkFlagRequired("namespace")

	releasesSubCmdHistory.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")

	releasesSubCmdList.Flags().StringP("namespace", "n", "", "Namespace to list to: dev, qa etc. (required)")
	releasesSubCmdList.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
	releasesSubCmdList.MarkFlagRequired("namespace")

	releasesSubCmdRollback.Flags().StringP("revision", "r", "0", "A previous release version")

	releasesSubCmdStatus.Flags().StringP("format", "f", "", "Format (optional: json|yaml)")
}

// Support functions to conduct the client call.

func runDelete(release string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Delete(release)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runDeploy(release string, tag string, namespace string, memo string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Deploy(release, tag, namespace, memo)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runHistory(release string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.History(release, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runList(namespace string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.List(namespace, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runRollback(release string, revision string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Rollback(release, revision)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}

func runStatus(release string, format string) error {
	cl := client.NewClient(clusterUrl, bearerToken)
	resp, err := cl.Status(release, format)
	if err != nil {
		return err
	}
	fmt.Println(resp.Message)
	return nil
}
