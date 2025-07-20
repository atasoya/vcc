package cmd

import (
	"fmt"
	"os"
	"strings"

	"vcc/utils"
	"vcc/utils/config"

	"github.com/spf13/cobra"
)

var tokyoVariant string
var tokyoConfigPath string

var tokyonightCmd = &cobra.Command{
	Use:   "tokyonight",
	Short: "Apply the Tokyo Night theme to Neovim",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsVariantEmpty(tokyoVariant) {
			tokyoVariant = utils.GetDefaultTokyoNightVariant()
		}

		tokyoNightVariants := utils.TokyoNightVariants()

		if !tokyoNightVariants.Has(tokyoVariant) {
			fmt.Printf("❌ Invalid variant: %s\n", tokyoVariant)
			fmt.Printf("✅ Valid variants: %s\n", strings.Join(tokyoNightVariants.GetAll(), ", "))
			os.Exit(1)
		}

		fmt.Printf("🌃 Switching to Tokyo Night with variant: %s\n", tokyoVariant)

		err := config.WriteTokyoNightConfig(tokyoVariant, tokyoConfigPath)
		if err != nil {
			fmt.Println("❌ Failed to write file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Theme applied! Restart Neovim or run :Lazy sync if needed.")

	},
}

func init() {
	rootCmd.AddCommand(tokyonightCmd)

	tokyonightCmd.Flags().StringVarP(
		&tokyoVariant,
		"variant", "v", "",
		"Tokyo Night variant (storm, night, day, moon)",
	)

	tokyonightCmd.Flags().StringVarP(
		&tokyoConfigPath,
		"path", "p", "",
		"Custom path for the colorscheme config file",
	)
}
