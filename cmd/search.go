/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Description string
var Code string

var selectedFlagCount int

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search emoji you want",
	Long: `
		With this subcommand you can search the emoji you want to use
		in your commit. You can search by code ' -c :memo: ' or by
		description ' -d "update" '. 
		You don't need the write all the description or code to search,
		the subcommand have autocomplete.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if selectedFlagCount != 1 {
			fmt.Println("You should use only one flag")
			return
		}

		fmt.Println("search called")
		if Description != "" {
			fmt.Printf("input: %v\n", Description)
			return
		}

		if Code != "" {
			fmt.Printf("input: %v\n", Code)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&Description, "description", "d", "", "Search by emoji description")
	searchCmd.Flags().StringVarP(&Code, "code", "c", "", "Search by emoji code")

	searchCmd.Flags().SetInterspersed(false) // Evita que Cobra parse argumentos después de flags

	searchCmd.Flags().Visit(func(f *pflag.Flag) {
		if f.Value.String() != "" {
			selectedFlagCount++
		}
	})
}
