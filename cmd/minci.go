package cmd

import (
	"errors"
	"flag"
	"strings"

	_ "github.com/golang/glog"
	"github.com/minio/minci/pkg/ci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "minci",
	Short: "minimalist ci system",
	RunE: func(c *cobra.Command, args []string) error {
		selfURL := viper.GetString("self-url")
		repo := viper.GetString("repository")
		githubID := viper.GetString("github-id")
		githubSecret := viper.GetString("github-secret")
		webhookSecret := viper.GetString("webhook-secret")
		port := viper.GetInt("port")

		if len(selfURL) == 0 {
			return errors.New("self-url cannot be empty")
		}
		if len(repo) == 0 {
			return errors.New("repository cannot be empty")
		}
		if len(githubID) == 0 {
			return errors.New("githubID cannot be empty")
		}
		if len(githubSecret) == 0 {
			return errors.New("githubSecret cannot be empty")
		}
		if len(webhookSecret) == 0 {
			return errors.New("webhookSecret cannot be empty")
		}

		return ci.StartCIServer(selfURL, repo, githubID, githubSecret, webhookSecret, port)
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	RootCmd.Flags().String("self-url", "", "url of this server")
	RootCmd.Flags().String("repository", "", "repository to run ci")
	RootCmd.Flags().String("github-id", "", "github application id")
	RootCmd.Flags().String("github-secret", "", "github application secret")
	RootCmd.Flags().String("webhook-secret", "", "github application webhook secret")
	RootCmd.Flags().Int("port", 8080, "CI server port")

	// parse the go default flagset to get flags for glog and other packages in future
	RootCmd.Flags().AddGoFlagSet(flag.CommandLine)

	// defaulting this to true so that logs are printed to console
	flag.Set("logtostderr", "true")

	//suppress the incorrect prefix in glog output
	flag.CommandLine.Parse([]string{})

	viper.BindPFlags(RootCmd.Flags())
}
