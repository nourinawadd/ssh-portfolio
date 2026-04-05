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
				m.scrollOffset = 0
			}
		case "left", "h", "shift+tab":
			if !m.detailOpen {
				m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
				m.cursor = 0
				m.scrollOffset = 0
			}

		// ── up: cursor in list tabs, scroll elsewhere ─────────────────────
		case "up", "k":
			if !m.detailOpen && m.isListTab() {
				if m.cursor > 0 {
					m.cursor--
				}
			} else {
				if m.scrollOffset > 0 {
					m.scrollOffset--
				}
			}

		// ── down: cursor in list tabs, scroll elsewhere ───────────────────
		case "down", "j":
			if !m.detailOpen && m.isListTab() {
				if m.cursor < m.maxCursor() {
					m.cursor++
				}
			} else {
				m.scrollOffset++ // clamped in view via applyScroll
			}

		// ── open detail ───────────────────────────────────────────────────
		case "enter", " ":
			if !m.detailOpen && m.isListTab() {
				m.detailOpen = true
				m.scrollOffset = 0
			}

		// ── close detail ──────────────────────────────────────────────────
		case "esc", "b", "backspace":
			if m.detailOpen {
				m.detailOpen = false
				m.scrollOffset = 0
			}
		}

	case tickMsg:
		m.frame++
		return m, tickCmd()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.cursor > m.maxCursor() {
			m.cursor = m.maxCursor()
		}
	}

	return m, nil
}

// isListTab returns true for tabs that use cursor-based navigation.
func (m Model) isListTab() bool {
	return m.activeTab == 1 || m.activeTab == 2
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