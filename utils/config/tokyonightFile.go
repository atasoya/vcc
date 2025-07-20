package config

import (
	"fmt"
	"os"
)

func WriteTokyoNightConfig(variant string) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"

	configContent := fmt.Sprintf(`return {
 {
    "folke/tokyonight.nvim",
    name = "tokyonight",
    priority = 1000,
    config = function()
      require("tokyonight").setup({ style = "%s" })
      vim.cmd.colorscheme("tokyonight")
    end,
  },
}`, variant)

	return os.WriteFile(configPath, []byte(configContent), 0644)
}

func GetTokyoNightConfigPath() string {
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
