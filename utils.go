package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

// ConfigPath 配置文件路径
func ConfigPath(file string) string {
	current, _ := user.Current()
	homeDir := current.HomeDir
	configDir := path.Join(homeDir, ".config", "gptfusion")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		fmt.Println("Error creating config dir", err)
	}
	filePath := path.Join(configDir, file)
	fmt.Println(filePath)
	return filePath
}
