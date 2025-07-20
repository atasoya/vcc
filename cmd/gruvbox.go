package cmd

import (
	"fmt"
	"os"
	"strings"

	"vcc/utils"
	"vcc/utils/config"

	"github.com/spf13/cobra"
)

var gruvboxVariant string

var gruvboxCmd = &cobra.Command{
	Use:   "gruvbox",
	Short: "Apply the Gruvbox theme to Neovim",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsVariantEmpty(gruvboxVariant) {
			gruvboxVariant = utils.GetDefaultGruvboxVariant()
		}

		gruvboxVariants := utils.GruvboxVariants()

		if !gruvboxVariants.Has(gruvboxVariant) {
			fmt.Printf("❌ Invalid variant: %s\n", gruvboxVariant)
			fmt.Printf("✅ Valid variants: %s\n", strings.Join(gruvboxVariants.GetAll(), ", "))
			os.Exit(1)
		}

		fmt.Printf("🪵 Switching to Gruvbox with variant: %s\n", gruvboxVariant)

		err := config.WriteGruvboxConfig(gruvboxVariant)
		if err != nil {
			fmt.Println("❌ Failed to write file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Theme applied! Restart Neovim or run :Lazy sync if needed.")

	},
}

func init() {
	rootCmd.AddCommand(gruvboxCmd)

	gruvboxCmd.Flags().StringVarP(
		&gruvboxVariant,
		"variant", "v", "",
		"Gruvbox variant (dark, light, hard, soft, material)",
	)
}
