package ui

import (
    "fmt"
    "math"
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
var spiralStyle = lipgloss.NewStyle().
    Foreground(lipgloss.Color("#7C3AED")).
    Padding(1, 2).
    Width(26)

func renderSpiral(frame int) string {
    rows := 13
    cols := 24
    // Use a float grid to track intensity
    grid := make([][]rune, rows)
    for i := range grid {
        grid[i] = make([]rune, cols)
        for j := range grid[i] {
            grid[i][j] = ' '
        }
    }

    cx := float64(cols) / 2.0
    cy := float64(rows) / 2.0

    type ring struct {
        radius  float64
        dots    int
        speed   float64
        char    rune
    }

    rings := []ring{
        {1.5, 6,  0.08, '·'},
        {3.0, 10, 0.05, '•'},
        {4.5, 16, 0.03, '●'},
        {6.0, 22, 0.02, '•'},
        {7.5, 28, 0.01, '·'},
    }

    for _, r := range rings {
        phase := float64(frame) * r.speed
        for d := 0; d < r.dots; d++ {
            angle := (float64(d)/float64(r.dots))*2*math.Pi + phase
            x := cx + r.radius*2.2*math.Cos(angle)
            y := cy + r.radius*math.Sin(angle)

            xi := int(math.Round(x))
            yi := int(math.Round(y))

            if yi >= 0 && yi < rows && xi >= 0 && xi < cols {
                grid[yi][xi] = r.char
            }
        }
    }

    var sb strings.Builder
    for _, row := range grid {
        sb.WriteString(string(row))
        sb.WriteRune('\n')
    }
    return sb.String()
}

func (m Model) View() string {
    header := m.renderHeader()
    tabs   := m.renderTabs()

    spiral  := spiralStyle.Render(renderSpiral(m.frame))
    content := contentStyle.Width(m.width - 30).Render(m.renderContent())

    middle := lipgloss.JoinHorizontal(lipgloss.Top, spiral, content)
    footer := m.renderFooter()

    return lipgloss.JoinVertical(lipgloss.Left, header, tabs, middle, footer)
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