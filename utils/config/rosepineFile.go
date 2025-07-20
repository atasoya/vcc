package config

import (
	"fmt"
	"os"
)

func WriteRosePineConfig(variant string, customPath string) error {
	configPath := GetConfigPath(customPath)

	configContent := fmt.Sprintf(`return {
 {
    "rose-pine/neovim",
    name = "rose-pine",
    priority = 1000,
    config = function()
      require("rose-pine").setup({ variant = "%s" })
      vim.cmd.colorscheme("rose-pine")
    end,
  },
}`, variant)

	return os.WriteFile(configPath, []byte(configContent), 0644)
}
