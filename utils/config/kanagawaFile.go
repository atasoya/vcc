package config

import (
	"fmt"
	"os"
)

func WriteKanagawaConfig(variant string, customPath string) error {
	configPath := GetConfigPath(customPath)

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
