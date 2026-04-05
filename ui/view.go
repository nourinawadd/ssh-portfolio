package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ─── Palette ──────────────────────────────────────────────────────────────────

var (
	clrPurple = lipgloss.Color("#7C3AED")
	clrGreen  = lipgloss.Color("#10B981")
	clrAmber  = lipgloss.Color("#F59E0B")
	clrDim    = lipgloss.Color("#6B7280")
	clrBright = lipgloss.Color("#F3F4F6")
	clrIndigo = lipgloss.Color("#818CF8")
	clrCyan   = lipgloss.Color("#22D3EE")
)

// ─── Base styles ─────────────────────────────────────────────────────────────

var (
	headerStyle = lipgloss.NewStyle().Foreground(clrGreen).Bold(true)
	subStyle    = lipgloss.NewStyle().Foreground(clrDim)

	activeTabSty = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(clrPurple).
			Padding(0, 2)

	inactiveTabSty = lipgloss.NewStyle().
			Foreground(clrDim).
			Padding(0, 2)

	// Main content box — purple border
	contentBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrPurple).
			Padding(1, 2)

	// Detail panel box — amber border
	detailBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrAmber).
			Padding(1, 2)

	sectionSty  = lipgloss.NewStyle().Foreground(clrGreen).Bold(true)
	labelSty    = lipgloss.NewStyle().Foreground(clrAmber).Bold(true)
	dimSty      = lipgloss.NewStyle().Foreground(clrDim)
	brightSty   = lipgloss.NewStyle().Foreground(clrBright).Bold(true)
	linkSty     = lipgloss.NewStyle().Foreground(clrCyan).Underline(true)
	cursorSty   = lipgloss.NewStyle().Foreground(clrPurple).Bold(true)
	footerSty   = lipgloss.NewStyle().Foreground(clrDim)

	tagSty = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(clrPurple).
		Padding(0, 1)

	catSty = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(clrIndigo).
		Padding(0, 1)

	spiralSty = lipgloss.NewStyle().Foreground(clrPurple)
)

// ─── Data ─────────────────────────────────────────────────────────────────────

type project struct {
	name    string
	tech    string
	short   string
	bullets []string
	github  string
}

type sideProject struct {
	name     string
	category string
	short    string
	bullets  []string
	link     string
	tags     []string
}

var projects = []project{
	{
		name:  "Tether Note",
		tech:  "Node.js  ·  MongoDB  ·  Express  ·  React",
		short: "Time-delayed note delivery web app",
		bullets: []string{
			"Scheduled delivery system letting users message their future selves",
			"Node.js cron jobs + MongoDB for time-based scheduling & delivery logic",
			"React frontend with Express/MongoDB backend for auth & note management",
		},
		github: "https://github.com/nourinawadd",
	},
	{
		name:  "Subscriptions Tracker API",
		tech:  "Node.js  ·  Express  ·  MongoDB",
		short: "RESTful API for subscription lifecycle management",
		bullets: []string{
			"15+ endpoints covering full subscription management operations",
			"JWT authentication with bcrypt hashing and role-based access control",
			"MongoDB/Mongoose schema across 3 collections with referential integrity",
		},
		github: "https://github.com/nourinawadd",
	},
	{
		name:  "Social Feed App",
		tech:  "Node.js  ·  MongoDB  ·  Express  ·  Angular",
		short: "Full-stack social media application",
		bullets: []string{
			"User registration, posts, likes, comments & real-time UI updates",
			"JWT authentication + RESTful API with Express/MongoDB backend",
			"Angular SPA with HttpClient for API calls & responsive Bootstrap UI",
		},
		github: "https://github.com/nourinawadd",
	},
}

// sideProjects — update with real game dev titles as needed
var sideProjects = []sideProject{
	{
		name:     "Indie Game Projects",
		category: "Game Dev",
		short:    "Personal game development experiments and prototypes",
		bullets: []string{
			"Various indie games built as personal passion projects",
			"Exploring mechanics, art direction, and interactive storytelling",
			"Source available on GitHub — check the link below",
		},
		link: "https://github.com/nourinawadd",
		tags: []string{"Game Dev", "Unity", "C#", "Indie"},
	},
	{
		name:     "Behance Portfolio",
		category: "Graphic Design",
		short:    "Visual design work: branding, UI, illustration",
		bullets: []string{
			"Branding, logo design, and visual identity projects",
			"UI/UX mockups and interface design explorations",
			"Digital illustrations and creative compositions",
		},
		link: "https://www.behance.net/nourinawadd",
		tags: []string{"Graphic Design", "Branding", "UI/UX", "Illustration"},
	},
}

// ─── OSC 8 hyperlink ─────────────────────────────────────────────────────────

// osc8 wraps text in an OSC 8 hyperlink escape sequence.
// In supporting terminals (iTerm2, WezTerm, etc.) the text becomes clickable.
// Falls back gracefully to plain styled text in unsupported terminals.
func osc8(url, text string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

// clickLink renders a styled, clickable hyperlink.
func clickLink(url string) string {
	return osc8(url, linkSty.Render(url))
}

// ─── Spiral animation ────────────────────────────────────────────────────────

// renderSpiral draws a compact animated spiral (rows×cols character grid).
func renderSpiral(frame int) string {
	const rows, cols = 12, 22
	var grid [rows][cols]rune
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	cx := float64(cols) / 2.0
	cy := float64(rows) / 2.0

	type ring struct {
		radius float64
		dots   int
		speed  float64
		char   rune
	}
	rings := []ring{
		{1.0, 5, 0.08, '·'},
		{2.0, 9, 0.05, '•'},
		{3.0, 14, 0.03, '●'},
		{4.3, 20, 0.02, '•'},
		{5.4, 27, 0.01, '·'},
	}

	for _, r := range rings {
		phase := float64(frame) * r.speed
		for d := 0; d < r.dots; d++ {
			angle := (float64(d)/float64(r.dots))*2*math.Pi + phase
			x := cx + r.radius*2.1*math.Cos(angle)
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
		sb.WriteString(string(row[:]) + "\n")
	}
	return sb.String()
}

// ─── Top-level view ──────────────────────────────────────────────────────────

func (m Model) View() string {
	if m.width < 30 || m.height < 8 {
		return lipgloss.NewStyle().Foreground(clrAmber).Render(
			"\n  [ terminal too small — please resize ]",
		)
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		m.renderHeader(),
		m.renderTabBar(),
		m.renderMiddle(),
		m.renderFooter(),
	)
}

// ─── Header ───────────────────────────────────────────────────────────────────

func (m Model) renderHeader() string {
	if m.width < 55 {
		// compact single-line header for narrow terminals
		return headerStyle.Render("Nourin Awad")
	}
	ascii := ` ▗▖  ▗▖ ▄▄▄  █  ▐▌ ▄▄▄ ▄ ▄▄▄▄       ▗▄▖ ▄   ▄ ▗▞▀▜▌▐▌▄
 ▐▛▚▖▐▌█   █ ▀▄▄▞▘█    ▄ █   █     ▐▌ ▐▌█ ▄ █ ▝▚▄▟▌▐▌  
 ▐▌ ▝▜▌▀▄▄▄▀      █    █ █   █     ▐▛▀▜▌█▄█▄█   ▗▞▀▜▌  
 ▐▌  ▐▌                █           ▐▌ ▐▌        ▝▚▄▟▌  `
	return headerStyle.Render(ascii)
}

// ─── Tab bar ──────────────────────────────────────────────────────────────────

func (m Model) renderTabBar() string {
	// Use short names when the terminal is narrow
	names := m.tabs
	if m.width < 65 {
		names = []string{"About", "Proj", "Side", "Skills", "Exp", "Contact"}
	}
	if m.width < 48 {
		names = []string{"Abt", "Prj", "Sid", "Skl", "Exp", "Cnt"}
	}

	tabs := make([]string, len(names))
	for i, name := range names {
		if i == m.activeTab {
			tabs[i] = activeTabSty.Render(name)
		} else {
			tabs[i] = inactiveTabSty.Render(name)
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

// ─── Middle (spiral + content) ────────────────────────────────────────────────

const sidebarW = 26 // spiral sidebar fixed width

func (m Model) renderMiddle() string {
	// Choose box style: amber border for detail panels, purple otherwise
	boxSty := contentBoxSty
	if m.detailOpen {
		boxSty = detailBoxSty
	}

	if m.width < 72 {
		// Narrow layout: no spiral, content fills full width
		inner := m.width - 8 // border (2) + padding (4) + margin (2)
		if inner < 10 {
			inner = 10
		}
		return boxSty.Width(m.width - 2).Render(m.renderContent(inner))
	}

	// Wide layout: spiral on left, content on right
	spiral := spiralSty.
		Width(sidebarW).
		Padding(1, 1).
		Render(renderSpiral(m.frame))

	contentW := m.width - sidebarW - 2
	inner := contentW - 8 // border + padding overhead
	if inner < 10 {
		inner = 10
	}

	content := boxSty.Width(contentW).Render(m.renderContent(inner))
	return lipgloss.JoinHorizontal(lipgloss.Top, spiral, content)
}

// ─── Content router ───────────────────────────────────────────────────────────

func (m Model) renderContent(w int) string {
	switch m.activeTab {
	case 0:
		return m.renderAbout(w)
	case 1:
		return m.renderProjects(w)
	case 2:
		return m.renderSideProjects(w)
	case 3:
		return m.renderSkills(w)
	case 4:
		return m.renderExperience(w)
	case 5:
		return m.renderContact(w)
	}
	return ""
}

// ─── About ────────────────────────────────────────────────────────────────────

func (m Model) renderAbout(w int) string {
	var b strings.Builder

	b.WriteString(sectionSty.Render("◆ About Me") + "\n\n")
	b.WriteString("  Backend developer from Mansoura, Egypt. I build scalable\n")
	b.WriteString("  systems and clean APIs, primarily with Node.js and .NET.\n\n")
	b.WriteString("  Studying Communications & Computer Engineering at\n")
	b.WriteString("  Mansoura University (GPA: 3.95 / 4.0).\n\n")

	b.WriteString(sectionSty.Render("◆ Education") + "\n\n")
	b.WriteString("  " + labelSty.Render("Mansoura University") + "\n")
	b.WriteString("  BE · Communications & Computer Engineering\n")
	b.WriteString("  " + dimSty.Render("Sep 2022 – Present") + "   " +
		lipgloss.NewStyle().Foreground(clrGreen).Render("GPA: 3.95 / 4.0") + "\n\n")

	b.WriteString(sectionSty.Render("◆ Extracurriculars") + "\n\n")
	b.WriteString("  " + labelSty.Render("IEEE Mansoura Student Branch") + "\n")
	b.WriteString("  Technical Director  " + dimSty.Render("Jun – Sep 2024") + "\n")
	b.WriteString("  Led national event: 880+ participants, 200+ teams,\n")
	b.WriteString("  managing system setup across 4 tracks.\n\n")

	b.WriteString(sectionSty.Render("◆ Languages") + "\n\n")
	b.WriteString("  " + tagSty.Render("English · Fluent") + "  " +
		tagSty.Render("Arabic · Native") + "\n")

	return b.String()
}

// ─── Projects ─────────────────────────────────────────────────────────────────

func (m Model) renderProjects(w int) string {
	if m.detailOpen {
		return m.renderProjectDetail(m.cursor)
	}

	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Technical Projects") + "\n")
	b.WriteString(dimSty.Render("  ↑↓ navigate   enter open   esc back") + "\n\n")

	for i, p := range projects {
		selected := i == m.cursor
		pfx := "  "
		if selected {
			pfx = cursorSty.Render("▶ ")
		}

		nameSty := lipgloss.NewStyle().Foreground(clrBright)
		if selected {
			nameSty = nameSty.Bold(true)
		}

		b.WriteString(pfx + nameSty.Render(p.name) + "\n")
		b.WriteString("    " + dimSty.Render(p.short) + "\n")
		b.WriteString("    " + lipgloss.NewStyle().Foreground(clrIndigo).Render(p.tech) + "\n\n")
	}

	return b.String()
}

func (m Model) renderProjectDetail(idx int) string {
	if idx < 0 || idx >= len(projects) {
		return ""
	}
	p := projects[idx]

	var b strings.Builder
	b.WriteString(dimSty.Render("esc · back to projects") + "\n\n")

	b.WriteString(labelSty.Render(p.name) + "\n")
	b.WriteString(lipgloss.NewStyle().Foreground(clrIndigo).Render(p.tech) + "\n\n")

	b.WriteString(sectionSty.Render("About") + "\n")
	b.WriteString("  " + p.short + "\n\n")

	b.WriteString(sectionSty.Render("Highlights") + "\n")
	for _, bullet := range p.bullets {
		b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
	}

	b.WriteString("\n" + sectionSty.Render("Repository") + "\n")
	b.WriteString("  " + clickLink(p.github) + "\n")
	b.WriteString("  " + dimSty.Render("ctrl+click to open in browser (supported in iTerm2, WezTerm, etc.)") + "\n")

	return b.String()
}

// ─── Side Projects ────────────────────────────────────────────────────────────

func (m Model) renderSideProjects(w int) string {
	if m.detailOpen {
		return m.renderSideProjectDetail(m.cursor)
	}

	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Side Projects") + "\n")
	b.WriteString(dimSty.Render("  ↑↓ navigate   enter explore   esc back") + "\n\n")

	for i, sp := range sideProjects {
		selected := i == m.cursor
		pfx := "  "
		if selected {
			pfx = cursorSty.Render("▶ ")
		}

		nameSty := lipgloss.NewStyle().Foreground(clrBright)
		if selected {
			nameSty = nameSty.Bold(true)
		}

		cat := catSty.Render(sp.category)
		b.WriteString(pfx + nameSty.Render(sp.name) + "  " + cat + "\n")
		b.WriteString("    " + dimSty.Render(sp.short) + "\n\n")
	}

	return b.String()
}

func (m Model) renderSideProjectDetail(idx int) string {
	if idx < 0 || idx >= len(sideProjects) {
		return ""
	}
	sp := sideProjects[idx]

	var b strings.Builder
	b.WriteString(dimSty.Render("esc · back to side projects") + "\n\n")

	b.WriteString(labelSty.Render(sp.name) + "  " + catSty.Render(sp.category) + "\n\n")

	b.WriteString(sectionSty.Render("About") + "\n")
	b.WriteString("  " + sp.short + "\n\n")

	b.WriteString(sectionSty.Render("Highlights") + "\n")
	for _, bullet := range sp.bullets {
		b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
	}

	b.WriteString("\n" + sectionSty.Render("Tags") + "\n  ")
	for _, tag := range sp.tags {
		b.WriteString(tagSty.Render(tag) + " ")
	}

	b.WriteString("\n\n" + sectionSty.Render("Link") + "\n")
	b.WriteString("  " + clickLink(sp.link) + "\n")
	b.WriteString("  " + dimSty.Render("ctrl+click to open in browser") + "\n")

	return b.String()
}

// ─── Skills ───────────────────────────────────────────────────────────────────

func (m Model) renderSkills(w int) string {
	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Technical Skills") + "\n\n")

	sections := []struct {
		label string
		items []string
	}{
		{
			"Languages",
			[]string{"C", "C#", "Python", "JavaScript", "SQL"},
		},
		{
			"Frameworks & Tools",
			[]string{".NET", "ASP.NET MVC", "Entity Framework", "Express.js", "Node.js", "Angular", "Mongoose", "Git", "Postman"},
		},
		{
			"Databases & Architecture",
			[]string{"SQL Server", "MongoDB", "RESTful API Design", "JWT Authentication", "MVC Architecture"},
		},
	}

	for _, s := range sections {
		b.WriteString("  " + labelSty.Render(s.label) + "\n  ")
		for _, item := range s.items {
			b.WriteString(tagSty.Render(item) + " ")
		}
		b.WriteString("\n\n")
	}

	return b.String()
}

// ─── Experience ───────────────────────────────────────────────────────────────

func (m Model) renderExperience(w int) string {
	type job struct {
		company  string
		role     string
		period   string
		location string
		bullets  []string
	}

	jobs := []job{
		{
			"Information Technology Institute (ITI)",
			".NET Full Stack Intern",
			"Jul – Sep 2025",
			"Mansoura, Egypt",
			[]string{
				"Built relational databases in SQL Server (stored procedures, joins, queries)",
				"Developed ASP.NET MVC web apps in C# with Razor views, controllers & routing",
				"Implemented Entity Framework (LINQ, migrations, CRUD) for DB integration",
			},
		},
		{
			"National Telecommunications Institute (NTI)",
			"MEAN Stack Intern",
			"Jun – Jul 2025",
			"Mansoura, Egypt",
			[]string{
				"Built RESTful APIs with JWT authentication and MongoDB CRUD via Mongoose",
				"Developed Angular SPAs with routing, reactive forms & state management",
				"Applied API security, testing & performance optimization best practices",
			},
		},
		{
			"Nile University",
			"Blockchain Research Intern",
			"Jul – Sep 2025",
			"Cairo, Egypt",
			[]string{
				"Researched scalability: Lightning Network, sharding, DAGs, layer-2 protocols",
				"Co-authored a comparative survey outlining findings & future research directions",
			},
		},
	}

	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Professional Experience") + "\n\n")

	for _, j := range jobs {
		b.WriteString("  " + labelSty.Render(j.company) + "\n")
		b.WriteString("  " + brightSty.Render(j.role) +
			"  " + dimSty.Render(j.period) + "\n")
		b.WriteString("  " + dimSty.Render(j.location) + "\n")
		for _, bullet := range j.bullets {
			b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
		}
		b.WriteString("\n")
	}

	return b.String()
}

// ─── Contact ──────────────────────────────────────────────────────────────────

func (m Model) renderContact(w int) string {
	type entry struct {
		label   string
		display string
		url     string
	}

	entries := []entry{
		{"Email", "nourinawad@gmail.com", "mailto:nourinawad@gmail.com"},
		{"GitHub", "https://github.com/nourinawadd", "https://github.com/nourinawadd"},
		{"LinkedIn", "https://linkedin.com/in/nourinawad", "https://linkedin.com/in/nourinawad"},
		{"Behance", "https://www.behance.net/nourinawadd", "https://www.behance.net/nourinawadd"},
	}

	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Get in Touch") + "\n\n")

	for _, e := range entries {
		label := labelSty.Render(fmt.Sprintf("%-10s", e.label+":"))
		link := osc8(e.url, linkSty.Render(e.display))
		b.WriteString("  " + label + "  " + link + "\n\n")
	}

	b.WriteString(dimSty.Render(
		"  ctrl+click any link above to open  ·  supported in iTerm2, WezTerm, Kitty, etc.",
	) + "\n")

	return b.String()
}

// ─── Footer ───────────────────────────────────────────────────────────────────

func (m Model) renderFooter() string {
	var hint string
	switch {
	case m.detailOpen:
		hint = "esc back   q quit"
	case m.activeTab == 1 || m.activeTab == 2:
		hint = "← → tabs   ↑↓ select   enter open   esc back   q quit"
	default:
		hint = "← → tabs   q quit"
	}
	return footerSty.Render("  " + hint)
}