package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
	"os"
	"os/user"
	"path"
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
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
	wruntime.WindowExecJS(app.ctx, "window.location.href='https://yiyan.baidu.com/';")
	app.updateDialog(false)
}

type PlatForm struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

func (app *App) ReadMenu() []PlatForm {
	filePath := configPath("menu.json")

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
				Label: "自定义Demo",
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

func (app *App) EditMenu(platorms []PlatForm) {
	filePath := configPath("menu.json")
	content, err := json.Marshal(platorms)
	if err != nil {
		fmt.Println("Error marshalling json", err)
	}
	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
	}
	fmt.Println("Updated file")
	app.updateCustomMenu()
}

func (app *App) WriteHome(url string) {
	filePath := configPath("home.txt")
	data := []byte(url)
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
	}
}

func (app *App) updateCustomMenu() {
	_menu := app.initMenu()
	wruntime.MenuSetApplicationMenu(app.ctx, _menu)
	wruntime.MenuUpdateApplicationMenu(app.ctx)
	wruntime.WindowReload(app.ctx)
}

func (app *App) initMenu() *menu.Menu {
	trayMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		trayMenu.Append(menu.AppMenu())
		trayMenu.Append(menu.EditMenu())
	}
	// 内置聊天平台
	platforms := trayMenu.AddSubmenu("AI聊天平台")
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
	// 内置绘画平台
	aiDraw := trayMenu.AddSubmenu("AI绘画平台")
	aiDraw.AddText("文心一格", nil, func(cd *menu.CallbackData) {
		wruntime.WindowExecJS(app.ctx, "window.location.replace('https://yige.baidu.com/');")
	})

	custom := trayMenu.AddSubmenu("自定义平台")
	custom_menu_data := app.ReadMenu()
	fmt.Println(custom_menu_data)
	for _, p := range custom_menu_data {
		// go的for循环陷阱
		temp := p
		custom.Append(&menu.MenuItem{
			Label: temp.Label,
			Type:  menu.TextType,
			Click: func(cd *menu.CallbackData) {
				jscode := fmt.Sprintf("window.location.replace('%s');", temp.Url)
				wruntime.WindowExecJS(app.ctx, jscode)
			},
		})
	}
	// 工具
	setting := trayMenu.AddSubmenu("设置")
	setting.AddText("打开设置", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
		home := configPath("home.txt")
		url, err := os.ReadFile(home)
		if err != nil {
			fmt.Println("Error reading file", err)
			url = []byte("wails://wails/")
		}
		data := string(url)
		fmt.Println(data)
		wruntime.WindowExecJS(app.ctx, fmt.Sprintf("window.location.replace('%s');", data))
		wruntime.WindowReload(app.ctx)
	})
	setting.AddText("侧边栏模式", keys.CmdOrCtrl("s"), func(cd *menu.CallbackData) {
		app.SideMode()
	})
	setting.AddText("窗口模式", keys.CmdOrCtrl("w"), func(cd *menu.CallbackData) {
		app.WindowMode()
	})

	about := trayMenu.AddSubmenu("帮助")
	about.AddText("关于我们", nil, func(cd *menu.CallbackData) {
		wruntime.MessageDialog(app.ctx, wruntime.MessageDialogOptions{
			Title:   "关于我们",
			Message: "GPTFusion " + Version + "\n\n" + "作者：lpdswing\n\n" + "请关注微信公众号：Go学习日记",
		})
	})
	about.AddText("前往Github", nil, func(cd *menu.CallbackData) {
		wruntime.BrowserOpenURL(app.ctx, "https://github.com/lpdswing/chatgpt")
	})
	about.AddText("检查更新", nil, func(cd *menu.CallbackData) {
		// 检查更新
		app.updateDialog(true)
	})
	return trayMenu
}

func configPath(file string) string {
	user, _ := user.Current()
	homeDir := user.HomeDir
	configDir := path.Join(homeDir, ".config", "gptfusion")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		fmt.Println("Error creating config dir", err)
	}
	filePath := path.Join(configDir, file)
	fmt.Println(filePath)
	return filePath
}

func checkForUpdate(currentVersion string) (bool, string, error) {
	// 检查更新
	url := "https://api.github.com/repos/lpdswing/chatgpt/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return false, "", err
	}
	fmt.Println(release.TagName)
	return release.TagName != currentVersion, release.TagName, nil
}

type Release struct {
	TagName string `json:"tag_name"`
}

func (app *App) updateDialog(show bool) {
	ok, LatestVersion, err := checkForUpdate(Version)
	if err != nil {
		fmt.Println(err)
	}
	// 和当前版本不一致
	if ok {
		wruntime.MessageDialog(app.ctx, wruntime.MessageDialogOptions{
			Title:   "软件更新",
			Message: "您的当前版本为" + Version + "\n" + "最新版本为" + LatestVersion + "\n" + "请前往Github下载最新版本",
		})
	} else {
		if show {
			wruntime.MessageDialog(app.ctx, wruntime.MessageDialogOptions{
				Title:   "软件更新",
				Message: "您的当前版本为" + Version + "\n" + "已经是最新版本",
			})
		}
	}
}

func (app *App) GetVersion() string {
	return Version
}

func (app *App) WindowMode() {
	wruntime.WindowSetSize(app.ctx, 1024, 768)
}

func (app *App) SideMode() {
	wruntime.WindowSetPosition(app.ctx, 0, 25)
	wruntime.WindowSetSize(app.ctx, 400, 768)
	wruntime.WindowSetAlwaysOnTop(app.ctx, true)
}

type Setting struct {
	Mode              string `json:"mode"`
	AlwaysOnTop       bool   `json:"always_on_top"`
	HideWindowOnClose bool   `json:"hide_window_on_close"`
}

func (app *App) ReadSetting() Setting {
	file := configPath("setting.json")
	url, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file", err)
		url = []byte("{\"mode\":\"1\",\"always_on_top\":false,\"hide_window_on_close\":true}")
	}
	var setting Setting
	json.Unmarshal(url, &setting)
	fmt.Println(setting)
	return setting
}

func (app *App) WriteSetting(setting Setting) {
	file := configPath("setting.json")
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

func (app *App) reload(setting Setting) {
	if setting.Mode == "1" {
		app.WindowMode()
	} else {
		app.SideMode()
	}
	wruntime.WindowSetAlwaysOnTop(app.ctx, setting.AlwaysOnTop)
}
