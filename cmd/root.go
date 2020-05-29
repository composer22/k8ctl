package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Used globally for all commands
var (
	cfgFile     string
	cluster     string // Which server to use on what cluster?
	format      string // text, json, or yaml.
	bearerToken string // api token for the user access to the server
	clusterUrl  string // endpoint to the server in the cluster of choice.
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "k8ctl",
	Short: "Manage and deploy applications in a K8 cluster",
	Long:  "A command line client for deploying and managing applications and releases in a cluster/namespace.",
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.k8ctl.yaml)")
	RootCmd.PersistentFlags().StringVarP(&cluster, "cluster", "l", "", "Cluster to access (mandatory)")

	// Get values from config file.
	RootCmd.MarkFlagRequired("cluster")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

// initConfig reads in config file.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}
		viper.AddConfigPath(home) // adding home directory as first search path.
		viper.AddConfigPath(".")
		viper.SetConfigName(".k8ctl") // name of config file (without extension).
	}
	viper.AutomaticEnv()
	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Cannot find configuration file.\n\nERR: %s\n", err.Error())
		os.Exit(0)
	}

	// Sample code to keep around because it's cool.
	// clusters := viper.Get("clusters")
	// if rec, ok := clusters.(map[string]interface{}); ok {
	// 	for key, val := range rec {
	// 		fmt.Printf("Cluster %s:\n", key)
	// 		if subrec, ok := val.(map[string]interface{}); ok {
	// 			for key, val := range subrec {
	// 				fmt.Printf("%s:%s\n", key, val)
	// 			}
	// 		}
	// 	}
	// }

	// Retrieve the Cluster bearer token and url from the config based on the
	// cluster param.

	bearerToken = viper.GetString(fmt.Sprintf("clusters.%s.%s", cluster, "auth_token"))
	clusterUrl = viper.GetString(fmt.Sprintf("clusters.%s.%s", cluster, "url"))

	// Cluster not found.
	if clusterUrl == "" {
		fmt.Printf("Cluster name %s not found\n", cluster)
		os.Exit(0)
	}
}
