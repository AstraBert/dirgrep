package cmd

import (
	"fmt"
	"os"

	"github.com/rvfet/rich-go"
	"github.com/spf13/cobra"
)

var pattern string
var directory string
var contextWindow int
var recursive bool
var toSkip []string
var showHelp bool

var rootCmd = &cobra.Command{
	Use:   "dirgrep",
	Short: "Simple and intuitive CLI tool to perform grep operations directory-wise",
	Long:  "dirgrep is a simple and intuitive CLI tool that can perform grep operations within a specific diectory (recursively or not). Powered by concurrent Go, with love.",
	Run: func(cmd *cobra.Command, args []string) {
		if showHelp {
			cmd.Help()
		} else if pattern == "" {
			fmt.Println("Missing required option: `pattern`")
			cmd.Help()
		} else {
			mp, err := GrepMany(pattern, directory, recursive, toSkip, contextWindow)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Oops. An error while executing dirgrep '%s'\n", err)
				os.Exit(1)
			} else {
				rich.Print("[bold green]MATCHES[/]\n")
				for k := range mp {
					for _, v := range mp[k] {
						rich.Print("File: [bold white]" + k + "[/]\n")
						rich.Print(v)
						fmt.Println()
					}
				}
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing dirgrep '%s'\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceVarP(&toSkip, "skip", "s", nil, "One or more sub-directories to skip. Can be used multiple times, can be used with comma-separated values. Defaults to an empty list.")
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Pattern to search for within the given directory. Required.")
	rootCmd.Flags().StringVarP(&directory, "directory", "d", ".", "The directory to search for the pattern in. Defaults to the current working directory if not specified.")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Whether or not to search for files to grep recursively. Defaults to false if not used")
	rootCmd.Flags().IntVarP(&contextWindow, "context", "c", 0, "The context to add to the matches (number of charachters). Defaults to 0 if not used")
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Show the help message and exit.")

	rootCmd.MarkFlagRequired("pattern")
}
