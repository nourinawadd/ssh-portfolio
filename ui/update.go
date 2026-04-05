package ui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// ── quit ──────────────────────────────────────────────────────────────
		case "q", "ctrl+c":
			return m, tea.Quit

		// ── tab navigation (blocked while detail open) ─────────────────────
		case "right", "l", "tab":
			if !m.detailOpen {
				m.activeTab = (m.activeTab + 1) % len(m.tabs)
				m.cursor = 0
			}
		case "left", "h", "shift+tab":
			if !m.detailOpen {
				m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
				m.cursor = 0
			}

		// ── item navigation ───────────────────────────────────────────────
		case "up", "k":
			if !m.detailOpen && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if !m.detailOpen {
				if m.cursor < m.maxCursor() {
					m.cursor++
				}
			}

		// ── open detail ───────────────────────────────────────────────────
		case "enter", " ":
			if !m.detailOpen && (m.activeTab == 1 || m.activeTab == 2) {
				m.detailOpen = true
			}

		// ── close detail ──────────────────────────────────────────────────
		case "esc", "b", "backspace":
			m.detailOpen = false
		}

	case tickMsg:
		m.frame++
		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Safety: clamp cursor after resize
		if m.cursor > m.maxCursor() {
			m.cursor = m.maxCursor()
		}
	}

	return m, nil
}

// maxCursor returns the highest valid cursor index for the active tab.
func (m Model) maxCursor() int {
	switch m.activeTab {
	case 1:
		return len(projects) - 1
	case 2:
		return len(sideProjects) - 1
	default:
		return 0
	}
}