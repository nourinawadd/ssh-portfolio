package ui

import (
    "fmt"
    "strings"

    "github.com/charmbracelet/lipgloss"
)

var (
    primaryColor   = lipgloss.Color("#7C3AED")
    secondaryColor = lipgloss.Color("#10B981") 
    dimColor       = lipgloss.Color("#6B7280")

    activeTabStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(primaryColor).
        Padding(0, 2)

    inactiveTabStyle = lipgloss.NewStyle().
        Foreground(dimColor).
        Padding(0, 2)

    contentStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(primaryColor).
        Padding(1, 2)

    headerStyle = lipgloss.NewStyle().
        Foreground(secondaryColor).
        Bold(true)
)

func (m Model) View() string {
    var sections []string
    sections = append(sections, m.renderHeader())
    sections = append(sections, m.renderTabs())
    sections = append(sections, m.renderContent())
    sections = append(sections, m.renderFooter())
    return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m Model) renderHeader() string {
    ascii := `
▗▖  ▗▖ ▄▄▄  █  ▐▌ ▄▄▄ ▄ ▄▄▄▄       ▗▄▖ ▄   ▄ ▗▞▀▜▌▐▌▄ 
▐▛▚▖▐▌█   █ ▀▄▄▞▘█    ▄ █   █     ▐▌ ▐▌█ ▄ █ ▝▚▄▟▌▐▌  
▐▌ ▝▜▌▀▄▄▄▀      █    █ █   █     ▐▛▀▜▌█▄█▄█   ▗▞▀▜▌  
▐▌  ▐▌                █           ▐▌ ▐▌        ▝▚▄▟▌  
                                                                                                            
`                                                       
    return headerStyle.Render(ascii)
}

func (m Model) renderTabs() string {
    var tabs []string
    for i, tab := range m.tabs {
        if i == m.activeTab {
            tabs = append(tabs, activeTabStyle.Render(tab))
        } else {
            tabs = append(tabs, inactiveTabStyle.Render(tab))
        }
    }
    return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m Model) renderContent() string {
    var content string
    switch m.activeTab {
    case 0:
        content = m.renderAbout()
    case 1:
        content = m.renderProjects()
    case 2:
        content = m.renderSkills()
    case 3:
        content = m.renderContact()
    }
    return contentStyle.Width(m.width - 4).Render(content)
}

func (m Model) renderAbout() string {
    return fmt.Sprintf(`
  some random about text
  im from egypt
  i dont want to stay here 
  #getmeOUT
`)
}

func (m Model) renderProjects() string {
    projects := []struct{ name, desc, url string }{
        {"my-app", "A cool React app I built", "github.com/you/my-app"},
        {"another-thing", "Does something awesome", "github.com/you/another-thing"},
    }

    var lines []string
    for i, p := range projects {
        cursor := "  "
        if i == m.cursor {
            cursor = "▶ "
        }
        lines = append(lines, fmt.Sprintf("%s%s\n    %s\n    %s\n",
            cursor, 
            lipgloss.NewStyle().Bold(true).Render(p.name),
            p.desc,
            lipgloss.NewStyle().Foreground(dimColor).Render(p.url),
        ))
    }
    return strings.Join(lines, "\n")
}

func (m Model) renderSkills() string {
    return `
  Languages:   Go · JavaScript · TypeScript · Python
  Frontend:    React · Next.js · Tailwind
  Backend:     Node.js · Express · PostgreSQL
  Tools:       Git · Docker · Linux
`
}

func (m Model) renderContact() string {
    return `
  GitHub:    github.com/nourinawadd
  LinkedIn:  linkedin.com/in/nourinawad
  Email:     nourinawad@gmail.com
`
}

func (m Model) renderFooter() string {
    style := lipgloss.NewStyle().Foreground(dimColor)
    return style.Render("  ← → tabs   ↑ ↓ navigate   q quit")
}