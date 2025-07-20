package config

import (
	"fmt"
	"os"
)

func WriteCatppuccinConfig(variant string) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"

	configContent := fmt.Sprintf(`return {
 {
    "catppuccin/nvim",
    name = "catppuccin",
    priority = 1000,
    config = function()
      require("catppuccin").setup({ flavour = "%s" })
      vim.cmd.colorscheme("catppuccin")
    end,
  },
}`, variant)

	return os.WriteFile(configPath, []byte(configContent), 0644)
}

func GetCatppuccinConfigPath() string {
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
