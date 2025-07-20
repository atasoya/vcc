package config

import (
	"fmt"
	"os"
)

func WriteGruvboxConfig(variant string) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"

	configContent := fmt.Sprintf(`return {
 {
    "ellisonleao/gruvbox.nvim",
    name = "gruvbox",
    priority = 1000,
    config = function()
      require("gruvbox").setup({ contrast = "%s" })
      vim.cmd.colorscheme("gruvbox")
    end,
  },
}`, variant)

	return os.WriteFile(configPath, []byte(configContent), 0644)
}

func GetGruvboxConfigPath() string {
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
