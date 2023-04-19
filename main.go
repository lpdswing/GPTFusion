package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const Version = "v0.5.0"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GPT聚合版",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             app.initMenu(),
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		// 无边框
		Frameless:         false,
		HideWindowOnClose: app.ReadSetting().HideWindowOnClose,
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "GPT聚合版",
				Message: "版本号(" + Version + ")\n\n" + "GPT聚合版是一个聚合了多个GPT聊天机器人的聊天工具.\n\n" + "© 2023 by lpdswing.\n" + "All Rights Reserved.",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
