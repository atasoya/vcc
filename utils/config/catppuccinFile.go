package config

import (
	"fmt"
	"os"
)

func WriteCatppuccinConfig(variant string, customPath string) error {
	configPath := GetConfigPath(customPath)

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
