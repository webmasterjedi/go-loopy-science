package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"goloopyscience/loopy/db"
	"log"
	"os"
)

var userHomeDir, _ = os.UserHomeDir()
var journalDir string = userHomeDir + "\\Saved Games\\Frontier Developments\\Elite Dangerous\\test"
var journalDirExists bool = false

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := db.CreateTables()
	if err != nil {
		log.Fatal(err)
	}
	if a.checkJournalDir() {
		journalDirExists = true
	}
}

// checkJournalDir checks if the journal directory exists
func (a *App) checkJournalDir() bool {
	if _, err := os.Stat(journalDir); os.IsNotExist(err) {
		journalDirExists = false
		return journalDirExists
	}
	journalDirExists = true
	return journalDirExists
}

// OpenDirDialog function to choose a directory with OpenDirectoryDialog
func (a *App) OpenDirDialog() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	directoryPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory: dirname + "\\Saved Games\\Frontier Developments\\Elite Dangerous",
		Title:            "Choose Elite Dangerous Logs Directory",
		Filters:          []runtime.FileFilter{{DisplayName: "Log Files", Pattern: "*.log"}},
	})
	if err != nil {
		return "", fmt.Errorf("failed opening dialog - %s", err.Error())
	}
	return directoryPath, nil
}
