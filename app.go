package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"runtime"
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

type PlatForm struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

func (a *App) ReadMenu() []PlatForm {
	filePath := "menu.json"

	platforms := []PlatForm{}

	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("File exists")
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file", err)
		}
		err = json.Unmarshal(content, &platforms)
		if err != nil {
			fmt.Println("Error unmarshalling json", err)
		}

	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		platforms = []PlatForm{
			{
				Id:    "1",
				Label: "Google",
				Url:   "https://www.google.com",
			},
		}
		content, err := json.Marshal(platforms)
		if err != nil {
			fmt.Println("Error marshalling json", err)
		}
		err = os.WriteFile(filePath, content, 0644)
		if err != nil {
			fmt.Println("Error writing file", err)
		}
		fmt.Println("File created and written")
	} else {
		fmt.Println("Error reading file", err)
	}

	return platforms

}

func (a *App) EditMenu(platorms []PlatForm) {
	filePath := "menu.json"
	content, err := json.Marshal(platorms)
	if err != nil {
		fmt.Println("Error marshalling json", err)
	}
	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
	}
	fmt.Println("Updated file")
	a.updateCustomMenu()
}

func (a *App) updateCustomMenu() {
	wruntime.MenuSetApplicationMenu(a.ctx, a.initMenu())
	wruntime.MenuUpdateApplicationMenu(a.ctx)
	wruntime.WindowReload(a.ctx)
}

func (app *App) initMenu() *menu.Menu {
	var trayMenu *menu.Menu
	trayMenu = menu.NewMenu()
	if runtime.GOOS == "darwin" {
		trayMenu.Append(menu.AppMenu())
	}
	platforms := trayMenu.AddSubmenu("平台选择")
	platforms.AddText("文心一言(百度)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://yiyan.baidu.com/');")
	})
	platforms.AddText("通义千问(阿里)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://tongyi.aliyun.com/');")
	})
	platforms.AddText("NewBing(微软)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://www.bing.com/new');")
	})
	platforms.AddText("Bard(谷歌)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://bard.google.com/');")
	})
	platforms.AddSeparator()
	free := platforms.AddSubmenu("ChatGPT(免费)")
	free.AddText("ChatGPT(免费线路1)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://freegpt.one/');")
	})
	free.AddText("ChatGPT(免费线路2)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chatbot.theb.ai/');")
	})
	free.AddText("ChatGPT(免费线路3)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chatgpt-35-turbo.com/');")
	})
	free.AddText("ChatGPT(限额版)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chat.okis.dev/zh-CN?mode=chat');")
	})
	platforms.AddText("ChatGPT(官方版)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://chat.openai.com/chat');")
	})
	platforms.AddSeparator()
	platforms.AddText("POE(多平台)", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://poe.com/');")
	})
	custom := trayMenu.AddSubmenu("自定义")
	custom_menu_data := app.ReadMenu()
	fmt.Println(custom_menu_data)
	for _, p := range custom_menu_data {
		custom.AddText(p.Label, nil, func(cd *menu.CallbackData) {
			jscode := fmt.Sprintf("window.location.replace('%s');", p.Url)
			fmt.Println(jscode)
			wruntime.WindowExecJS(app.ctx, jscode)
		})
	}
	// 工具
	platformEdit := trayMenu.AddSubmenu("工具")
	platformEdit.AddText("平台管理", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('/');")
	})

	about := trayMenu.AddSubmenu("关于")
	about.AddText("访问Github", nil, func(cd *menu.CallbackData) {
		wruntime.BrowserOpenURL(app.ctx, "https://github.com/lpdswing/chatgpt")
	})
	return trayMenu
}
