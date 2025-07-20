package cmd

import (
	"fmt"
	"os"
	"strings"

	"vcc/utils"
	"vcc/utils/config"

	"github.com/spf13/cobra"
)

var kanagawaVariant string

var kanagawaCmd = &cobra.Command{
	Use:   "kanagawa",
	Short: "Apply the Kanagawa theme to Neovim",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsVariantEmpty(kanagawaVariant) {
			kanagawaVariant = utils.GetDefaultKanagawaVariant()
		}

		kanagawaVariants := utils.KanagawaVariants()

		if !kanagawaVariants.Has(kanagawaVariant) {
			fmt.Printf("❌ Invalid variant: %s\n", kanagawaVariant)
			fmt.Printf("✅ Valid variants: %s\n", strings.Join(kanagawaVariants.GetAll(), ", "))
			os.Exit(1)
		}

		fmt.Printf("🌊 Switching to Kanagawa with variant: %s\n", kanagawaVariant)

		err := config.WriteKanagawaConfig(kanagawaVariant)
		if err != nil {
			fmt.Println("❌ Failed to write file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Theme applied! Restart Neovim or run :Lazy sync if needed.")

	},
}

func init() {
	rootCmd.AddCommand(kanagawaCmd)

	kanagawaCmd.Flags().StringVarP(
		&kanagawaVariant,
		"variant", "v", "",
		"Kanagawa variant (wave, dragon, lotus)",
	)
}
