/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package category

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	cmd2 "notes/cmd"
	"notes/internal/editor"
	"os"
	"path/filepath"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templatesDirectory := filepath.Join(cmd2.NotesConfigDirectory, "templates")

		if err := os.MkdirAll(templatesDirectory, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if editor.OpenFileInEditor(cmd2.NotesEditor, filepath.Join(templatesDirectory, args[0])) != nil {
			log.Fatal("could not open editor")
		} else {
			addTemplateToCategoryConfiguration(filepath.Join(templatesDirectory, args[0]), args[0])
		}
	},
}

func init() {
	categoryCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addTemplateToCategoryConfiguration(templatePath, category string) {
	cfg := cmd2.Configuration{}

	data, err := os.ReadFile(cmd2.NotesConfigFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("unmarshal config file: %v", err)
	}

	cfg.Categories[category] = cmd2.Category{
		Template: templatePath,
	}

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = os.WriteFile(cmd2.NotesConfigFile, d, 0644)
	if err != nil {
		log.Fatalf("writing configuration file: %v", err)
	}
}
