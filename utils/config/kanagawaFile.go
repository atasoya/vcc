package config

import (
	"fmt"
	"os"
)

func WriteKanagawaConfig(variant string) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"

	configContent := fmt.Sprintf(`return {
 {
    "rebelot/kanagawa.nvim",
    name = "kanagawa",
    priority = 1000,
    config = function()
      require("kanagawa").setup()
      vim.cmd.colorscheme("kanagawa-%s")
    end,
  },
}`, variant)

	return os.WriteFile(configPath, []byte(configContent), 0644)
}

func GetKanagawaConfigPath() string {
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
