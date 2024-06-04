/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	Categories map[string]Category `yaml:"categories"`
}

type Category struct {
	Template string `yaml:"template"`
}

const defaultMetadata = `---
title: {{ .Name }}
category: {{ .Category }}
tags:
{{range .Tags}}
	- {{.}}
{{end}}
---
`

type Metadata struct {
	Name     string
	Category string
	Tags     []string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your note taking journey",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		initialConfiguration := Configuration{
			Categories: map[string]Category{},
		}

		d, err := yaml.Marshal(&initialConfiguration)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		if err := os.Mkdir(NotesDirectory, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(NotesConfigDirectory, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(NotesConfigFile, d, 0644)
		if err != nil {
			log.Fatalf("writing configuration file: %v", err)
		}

		err = os.WriteFile(filepath.Join(NotesConfigDirectory, "metadata.md"), []byte(defaultMetadata), 0644)
		if err != nil {
			log.Fatalf("writing metadata file: %v", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
