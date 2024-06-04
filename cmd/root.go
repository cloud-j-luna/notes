/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "notes",
	Short: "A simple note taking app",
	Long: `Notes is a note taking app built to be simple to use and work across different systems.
			It supports templates for categories of notes so you can jump on a meeting and start taking notes right away.`,
}

var NotesDirectory string
var NotesEditor string
var NotesConfigDirectory string
var NotesConfigFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.PersistentFlags().StringVar(&NotesDirectory, "directory", "", "Path to notes directory, default is $NOTES_CLI_NOTES_DIRECTORY or $HOME/notes")
}
