package editor

import (
	"os"
	"os/exec"
)

func OpenFileInEditor(editor, file string) error {
	writeCmd := exec.Command(editor, file)
	writeCmd.Stdin = os.Stdin
	writeCmd.Stdout = os.Stdout
	return writeCmd.Run()
}
