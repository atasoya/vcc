package config

import "os"

func GetConfigPath(customPath string) string {
	if customPath != "" {
		return customPath
	}
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
