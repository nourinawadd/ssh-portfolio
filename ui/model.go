package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width        int
	height       int
	activeTab    int
	tabs         []string
	cursor       int
	frame        int
	detailOpen   bool
	scrollOffset int
}

func NewModel(width, height int) Model {
	return Model{
		width:  width,
		height: height,
		tabs:   []string{"About", "Projects", "Side Projects", "Skills", "Experience", "Contact"},
	}
}

func (m Model) Init() tea.Cmd { return tickCmd() }

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(80*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}