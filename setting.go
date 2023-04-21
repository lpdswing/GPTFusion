package main

import (
	"encoding/json"
	"fmt"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type Setting struct {
	Mode              string `json:"mode"`
	AlwaysOnTop       bool   `json:"always_on_top"`
	HideWindowOnClose bool   `json:"hide_window_on_close"`
	RememberLastPage  bool   `json:"remember_last_page"`
	LastPage          string `json:"last_page"`
}

// WindowMode 窗口模式
func (app *App) WindowMode() {
	wruntime.WindowSetSize(app.ctx, 1024, 768)
}

// SideMode 侧边栏模式
func (app *App) SideMode() {
	wruntime.WindowSetPosition(app.ctx, 0, 25)
	wruntime.WindowSetSize(app.ctx, 450, 768)
}

// ReadSetting 读配置文件
func (app *App) ReadSetting() Setting {
	file := ConfigPath("setting.json")
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		data = []byte("{\"mode\":\"1\",\"always_on_top\":false,\"hide_window_on_close\":false,\"remember_last_page\":true,\"last_page\":\"https://chatbot.theb.ai/\"}")
	}
	var setting Setting
	err = json.Unmarshal(data, &setting)
	if err != nil {
		return Setting{}
	}
	return setting
}

// WriteSetting 写配置文件
func (app *App) WriteSetting(setting Setting) {
	file := ConfigPath("setting.json")
	data, err := json.Marshal(setting)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("Error Writing file", err)
	}
	app.reload(setting)
}

// 重新加载窗口设置
func (app *App) reload(setting Setting) {
	if setting.Mode == "1" {
		app.WindowMode()
	} else {
		app.SideMode()
	}
	wruntime.WindowSetAlwaysOnTop(app.ctx, setting.AlwaysOnTop)
}

// WriteLastPage 记录最后登录的页面
func (app *App) WriteLastPage(url string) {
	setting := app.ReadSetting()
	if setting.RememberLastPage {
		setting.LastPage = url
		app.WriteSetting(setting)
	}
}
