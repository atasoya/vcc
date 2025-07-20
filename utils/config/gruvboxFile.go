package config

import (
	"fmt"
	"os"
)

func WriteGruvboxConfig(variant string, customPath string) error {
	configPath := GetConfigPath(customPath)

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
