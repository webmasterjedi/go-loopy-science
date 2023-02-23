package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var userHomeDir, _ = os.UserHomeDir()
var journalDir string = userHomeDir + "\\Saved Games\\Frontier Developments\\Elite Dangerous"
var journalDirExists bool = false
var allBodies []string

type Entry struct {
	Event       string `json:"Event"`
	ScanType    string `json:"ScanType"`
	StarSystem  string `json:"StarSystem"`
	PlanetClass string `json:"PlanetClass"`
	StarType    string `json:"StarType"`
	Subclass    int    `json:"Subclass"`
	Luminosity  string `json:"Luminosity"`
}

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
	if a.checkJournalDir() {
		journalDirExists = true
		allBodies, err := processJournalDirectory(journalDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, body := range allBodies {
			fmt.Println(body)
		}
	}
}

// checkJournalDir checks if the journal directory exists
func (a *App) checkJournalDir() bool {
	if _, err := os.Stat(journalDir); os.IsNotExist(err) {
		return false
	}
	return true
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

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func parseJournalEntry(data []byte) (Entry, error) {
	var entry Entry
	err := json.Unmarshal(data, &entry)
	if err != nil {
		return entry, err
	}
	return entry, nil
}

func processJournalFile(path string) ([]string, error) {
	data, err := readFile(path)
	if err != nil {
		return nil, err
	}

	var bodies []string
	entries := strings.Split(string(data), "\n")
	for _, entry := range entries {
		if entry == "" {
			continue
		}
		entryJson, err := parseJournalEntry([]byte(entry))
		if err != nil {
			return nil, err
		}
		if entryJson.StarType != "" || entryJson.PlanetClass != "" {
			bodies = append(bodies, fmt.Sprintf("%+v\n", entryJson))
		}
	}

	sort.Strings(bodies)

	return bodies, nil
}

func processJournalDirectory(dir string) ([]string, error) {
	var allBodies []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".log" {
			fileBodies, err := processJournalFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			allBodies = append(allBodies, fileBodies...)
		}
	}

	return allBodies, nil
}
