package cmd

import (
	"fmt"
	"os"
	"strings"

	"vcc/utils"
	"vcc/utils/config"

	"github.com/spf13/cobra"
)

var variant string

// catppuccinCmd represents the catppuccin command
var catppuccinCmd = &cobra.Command{
	Use:   "catppuccin",
	Short: "Apply the Catppuccin theme to Neovim",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsVariantEmpty(variant) {
			variant = utils.GetDefaultCatppuccinVariant()
		}

		catppuccinVariants := utils.CatppuccinVariants()

		if !catppuccinVariants.Has(variant) {
			fmt.Printf("❌ Invalid variant: %s\n", variant)
			fmt.Printf("✅ Valid variants: %s\n", strings.Join(catppuccinVariants.GetAll(), ", "))
			os.Exit(1)
		}

		fmt.Printf("🌈 Switching to Catppuccin with variant: %s\n", variant)

		err := config.WriteCatppuccinConfig(variant)
		if err != nil {
			fmt.Println("❌ Failed to write file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Theme applied! Restart Neovim or run :Lazy sync if needed.")

	},
}

func init() {
	rootCmd.AddCommand(catppuccinCmd)

	catppuccinCmd.Flags().StringVarP(
		&variant,
		"variant", "v", "",
		"Catppuccin variant (latte, frappe, macchiato, mocha)",
	)
}
