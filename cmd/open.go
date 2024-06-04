package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"notes/internal/editor"
	"path/filepath"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note for editing",
	Long:  `Open a note for editing using your predefined editor`,
	Run: func(cmd *cobra.Command, args []string) {
		if editor.OpenFileInEditor(NotesEditor, filepath.Join(NotesDirectory, args[0])) != nil {
			log.Fatal("could not open note")
		} else {
		}
	},
}

func init() {
	RootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
