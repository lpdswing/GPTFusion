package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"net/http"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

const Version = "v0.6.0"

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GPTFusion",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: customHandler(),
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
				Title:   "GPTFusion",
				Message: "版本号(" + Version + ")\n\n" + "GPTFusion是一个聚合了多个GPT聊天机器人的聊天工具.\n\n" + "© 2023 by lpdswing.\n" + "All Rights Reserved.",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func customHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("My-Header", "test")
	})
}
