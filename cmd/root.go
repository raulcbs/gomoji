/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/raulcbs/gomoji/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ListEmojis bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomoji",
	Short: "Elevate your Git commits with 'gomoji' – seamlessly add expressive emojis to your commit messages using this CLI.",
	Long: `Introducing 'gomoji': Your effortless solution for enhancing Git commit messages. 
	This CLI seamlessly integrates into your workflow, allowing you to enrich your commits with expressive emojis from gitmoji.dev.
	Elevate collaboration and code tracking while adding a touch of creativity to your development process.
	Try 'gomoji' now and revolutionize the way you communicate through commits.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ListEmojis {
			emojis, _ := pkg.ListEmojisAvaliable()
			for _, emoji := range emojis {
				fmt.Printf("%v - %v - %v\n\n", emoji.Icon, emoji.Code, emoji.Name)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gomoji.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&ListEmojis, "list-emojis", "l", false, "List emojis avaliable")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gomoji" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gomoji")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
