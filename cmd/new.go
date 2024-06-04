/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"notes/internal/editor"
	"notes/internal/file"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new note",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if file.Exists(filepath.Join(NotesDirectory, args[0])) {
			log.Fatalf("note already exists")
		}

		macros := Macros{CreatedDate: time.Now()}

		metadataContent, err := os.ReadFile(filepath.Join(NotesConfigDirectory, "metadata.md"))
		if err != nil {
			log.Fatalf("reading metadata file")
		}

		metadata := Metadata{
			Name:     args[0],
			Category: CategoryName,
			Tags:     Tags,
		}

		templateContent := templateMetadata(string(metadataContent), metadata) + getCategoryTemplate(CategoryName)

		dir, _ := filepath.Split(filepath.Join(NotesDirectory, args[0]))

		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile(filepath.Join(NotesDirectory, args[0]), os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		tmpl, err := template.New("note").Parse(templateContent)
		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(f, macros)
		if err != nil {
			panic(err)
		}

		if editor.OpenFileInEditor(NotesEditor, filepath.Join(NotesDirectory, args[0])) != nil {
			log.Fatal("could not open note")
		} else {
		}
	},
}

var CategoryName string
var Tags []string

type Macros struct {
	CreatedDate time.Time
}

func init() {
	RootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&CategoryName, "category", "c", "", "Category of the note")
	newCmd.Flags().StringSliceVarP(&Tags, "tags", "t", []string{}, "Tags associated with the note")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getCategoryTemplate(category string) string {
	data, err := os.ReadFile(filepath.Join(NotesConfigDirectory, "templates", category))
	if err != nil {
		return ""
	}

	return string(data)
}

func templateMetadata(metadataTemplate string, metadata Metadata) string {
	var str strings.Builder

	tmpl, err := template.New("note").Parse(metadataTemplate)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&str, metadata)
	if err != nil {
		panic(err)
	}

	return str.String()
}
