package main

import (
	"embed"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	var trayMenu *menu.Menu
	trayMenu = menu.NewMenu()
	if runtime.GOOS == "darwin" {
		trayMenu.Append(menu.AppMenu())
	}
	platforms := trayMenu.AddSubmenu("平台选择")
	platforms.AddText("文心一言", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://yiyan.baidu.com/');")
	})
	platforms.AddText("ChatGPT(免费线路1)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://freegpt.one/');")
	})
	platforms.AddText("ChatGPT(免费线路2)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chatbot.theb.ai/');")
	})
	platforms.AddText("ChatGPT(免费线路3)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chatgpt-35-turbo.com/');")
	})
	platforms.AddText("ChatGPT(官方版)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chat.openai.com/chat');")
	})
	platforms.AddText("ChatGPT(限额版)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chat.okis.dev/zh-CN?mode=chat');")
	})
	platforms.AddText("POE", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://poe.com/');")
	})
	platforms.AddText("NewBing", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://www.bing.com/new');")
	})
	platforms.AddText("Bard", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://bard.google.com/');")
	})

	about := trayMenu.AddSubmenu("关于")
	about.AddText("访问Github", nil, func(cd *menu.CallbackData) {
		wruntime.BrowserOpenURL(app.ctx, "https://github.com/lpdswing/chatgpt")
	})

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "GPT聚合版",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             trayMenu,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Mac: &mac.Options{
			About: &mac.AboutInfo{
				Title:   "GPT聚合版",
				Message: "GPT聚合版是一个聚合了多个GPT聊天机器人的聊天工具.\n" + "© 2023 by lpdswing. All Rights Reserved.",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
