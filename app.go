package main

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

// Greet returns a greeting for the given name
func (a *App) Confirm(title string, message string) bool {
	text, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "No",
	})

	if err != nil {
		return false
	}
	return text == "Yes"
}

type SelectFileResponse struct {
	Path string `json:"path"`
	Err  string `json:"error"`
}

func (a *App) SelectFile() SelectFileResponse {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return SelectFileResponse{
			Err: err.Error(),
		}
	}
	selectedPath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:           homeDir,
		Title:                      "Upload File",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: false,
	})

	if err != nil {
		return SelectFileResponse{
			Err: err.Error(),
		}
	}
	return SelectFileResponse{
		Path: selectedPath,
	}
}
