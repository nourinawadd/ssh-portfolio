package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nourinawadd/ssh-portfolio/ui"
)

func main() {
	p := tea.NewProgram(
		ui.NewModel(80, 24),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
