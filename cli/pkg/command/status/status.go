package status

import (
	"fmt"
	"nipo-tunnel/pkg/config"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Command creates and returns the status cobra subcommand to check active tunnels.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Check tunnel status",
		Run: func(cmd *cobra.Command, args []string) {
			states, err := config.GetActiveStates()
			if err != nil || len(states) == 0 {
				fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("No active tunnels found."))
				return
			}

			titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).MarginBottom(1)
			pidStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("87")).Bold(true)
			nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
			urlStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Underline(true)
			dimStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

			fmt.Println(titleStyle.Render("Active Nipo Tunnels"))

			for _, state := range states {
				fmt.Printf("Instance PID: %s\n", pidStyle.Render(fmt.Sprintf("%d", state.PID)))
				for _, t := range state.Tunnels {
					route := fmt.Sprintf("localhost:%d", t.Port)
					fmt.Printf("  %s %s %s %s\n",
						nameStyle.Render(fmt.Sprintf("[%-4s]", t.Name)),
						dimStyle.Render(route),
						dimStyle.Render("->"),
						urlStyle.Render(t.URL),
					)
				}
				fmt.Println()
			}
		},
	}

	return cmd
}
