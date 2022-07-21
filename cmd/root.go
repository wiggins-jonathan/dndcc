package cmd

import (
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dndcc",
	Short: "A DnD 5e Character Creator",
	Long:  "dndcc is a 5th Edition Dungeons & Dragons Character Creator",
	RunE: func(cmd *cobra.Command, args []string) error {
		app := tea.NewProgram(ui.InitialModel())
		if err := app.Start(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
