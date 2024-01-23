package game

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func findSaveFiles() ([]string, error) {
	dir, err := terminalTillerDir()
	if err != nil {
		return nil, err
	}

	result := []string{}
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading save files: %w", err)
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".data" {
			result = append(result, filepath.Join(dir, f.Name()))
		}
	}

	return result, nil
}

func terminalTillerDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}

	dir = filepath.Join(dir, ".terminal-tiller")

	err = os.MkdirAll(dir, 0750)
	if err != nil {
		return "", fmt.Errorf("error creating save file directory: %w", err)
	}

	return dir, nil
}

func (g *game) saveAndQuit() tea.Msg {
	data, err := g.farm.Marshal()
	if err != nil {
		// TODO: how to handle errors in bubbletea???
		panic(fmt.Sprintf("MARSHAL Error: %v", err))
	}

	err = os.WriteFile(g.filename, data, 0644)
	if err != nil {
		panic(fmt.Sprintf("WRITE Error: %v", err))
	}

	return tea.Quit()
}
