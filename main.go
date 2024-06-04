/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"log"
	"notes/cmd"
	_ "notes/cmd/category"
	"os"
	"path/filepath"
)

type config struct {
	NotesDirectory string `conf:""`
	Editor         string `conf:"default:vim"`
}

func main() {
	const prefix = "NOTES_CLI"
	var cfg config
	_, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if !errors.Is(err, conf.ErrHelpWanted) {
			fmt.Printf("parsing config: %s\n", err)
			return
		}
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("getting home directory: %v", err)
	}

	if cfg.NotesDirectory == "" {
		cfg.NotesDirectory = filepath.Join(home, "notes")
	}

	cmd.NotesDirectory = cfg.NotesDirectory
	cmd.NotesEditor = cfg.Editor
	cmd.NotesConfigDirectory = filepath.Join(home, ".config", "notes")
	cmd.NotesConfigFile = filepath.Join(home, ".config", "notes", "config.yaml")
	cmd.Execute()
}
