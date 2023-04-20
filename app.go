package main

import (
	"context"
	"fmt"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
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
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	wruntime.WindowExecJS(app.ctx, "window.location.href='https://chatbot.theb.ai/';")
	app.updateDialog(false)
}

func (app *App) WriteHome(url string) {
	filePath := ConfigPath("home.txt")
	data := []byte(url)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
	}
}

func (app *App) GetVersion() string {
	return Version
}
