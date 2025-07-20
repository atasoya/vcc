package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vcc",
	Short: "neoVim color Scheme Changer (vcc) is a CLI tool to streamline color scheme changing process for nvim.",
	Long: `vcc is a CLI tool written in go to change color scheme in nvim with lazyvim.
This CLI tool automates the nvim color scheme changing process.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
