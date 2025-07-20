package config

import (
	"fmt"
	"os"
)

func WriteTokyoNightConfig(variant string, customPath string) error {
	configPath := GetConfigPath(customPath)

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
