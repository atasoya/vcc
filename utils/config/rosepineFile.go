package config

import (
	"fmt"
	"os"
)

func WriteRosePineConfig(variant string) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"

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

func GetRosePineConfigPath() string {
	return os.Getenv("HOME") + "/.config/nvim/lua/plugins/colorscheme.lua"
}
