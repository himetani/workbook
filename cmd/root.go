package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wb",
	Short: "wb",
	Long:  `wb`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var (
		consumerKey string
		verbose     bool
	)

	var authCmd = &cobra.Command{
		Use:   "auth [flags]",
		Short: "Get authorized information of pocket",
		Long:  `Get authorized information of pocket`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var logger *log.Logger = log.New(ioutil.Discard, "Info: ", log.LstdFlags)
			if verbose {
				logger = log.New(os.Stdout, "Info: ", log.LstdFlags)
			}

			return RunAuth(consumerKey, logger)
		},
	}

	authCmd.Flags().StringVarP(&consumerKey, "consumerKey", "c", "", "consumer key")
	authCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose log message")
	authCmd.MarkFlagRequired("consumerKey")

	rootCmd.AddCommand(authCmd)
}
