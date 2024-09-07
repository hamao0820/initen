package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "initen",
	Short: "A generator for Ebiten projects",
	Long: `Ebiten is a dead simple 2D game library for Go.
This application is a tool to generate the needed files
to quickly create a Ebiten project.`,
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
