package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"path/filepath"
)

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
		Filters:          []runtime.FileFilter{{"Log Files", "*.log"}},
	})
	if err != nil {
		return "", fmt.Errorf("failed opening dialog - %s", err.Error())
	}
	return directoryPath, nil
}

// ReadDir finds all files in a directory and read them line by line
func (a *App) ReadDir(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only add files with the .log extension
		if filepath.Ext(path) == ".log" && !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil
	}
	return files
}

// ReadFile reads a file line by line and parse as JSON
func (a *App) ReadFile(file string) []string {
	var lines []string
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return lines
}
