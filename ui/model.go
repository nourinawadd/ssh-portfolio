package ui

import (
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    width     int
    height    int
    activeTab int
    tabs      []string
    cursor    int
    frame     int
}

func NewModel(width, height int) Model {
    return Model{
        width:     width,
        height:    height,
        activeTab: 0,
        tabs:      []string{"About", "Projects", "Skills", "Contact"},
        cursor:    0,
        frame:     0,
    }
}

func (m Model) Init() tea.Cmd {
    return tickCmd()
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
    return tea.Tick(time.Millisecond*80, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}