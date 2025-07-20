package cmd

import (
	"fmt"
	"os"
	"strings"

	"vcc/utils"
	"vcc/utils/config"

	"github.com/spf13/cobra"
)

var rosepineVariant string

var rosepineCmd = &cobra.Command{
	Use:   "rosepine",
	Short: "Apply the Rose Pine theme to Neovim",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsVariantEmpty(rosepineVariant) {
			rosepineVariant = utils.GetDefaultRosePineVariant()
		}

		rosePineVariants := utils.RosePineVariants()

		if !rosePineVariants.Has(rosepineVariant) {
			fmt.Printf("❌ Invalid variant: %s\n", rosepineVariant)
			fmt.Printf("✅ Valid variants: %s\n", strings.Join(rosePineVariants.GetAll(), ", "))
			os.Exit(1)
		}

		fmt.Printf("🌹 Switching to Rose Pine with variant: %s\n", rosepineVariant)

		err := config.WriteRosePineConfig(rosepineVariant)
		if err != nil {
			fmt.Println("❌ Failed to write file:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Theme applied! Restart Neovim or run :Lazy sync if needed.")

	},
}

func init() {
	rootCmd.AddCommand(rosepineCmd)

	rosepineCmd.Flags().StringVarP(
		&rosepineVariant,
		"variant", "v", "",
		"Rose Pine variant (main, moon, dawn)",
	)
}
