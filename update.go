package main

import (
	"encoding/json"
	"fmt"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"net/http"
)

type Release struct {
	TagName string `json:"tag_name"`
}

// 检查更新
func checkForUpdate(currentVersion string) (bool, string, error) {
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

// 更新弹窗
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
