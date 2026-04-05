package ui

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ─── Palette ──────────────────────────────────────────────────────────────────

var (
	// dusty sky-blue – name, active borders, accent elements
	clrAccent = lipgloss.Color("#7BADC0")
	// pale blue-grey – section headers  (◆ headings)
	clrSection = lipgloss.Color("#B8C8D4")
	// muted blue-grey – labels, tech stacks
	clrLabel = lipgloss.Color("#9AACB8")
	// medium grey – dim / secondary text
	clrDim = lipgloss.Color("#6A7580")
	// near-white with cool tint – primary readable text
	clrBright = lipgloss.Color("#DCE4EA")
	// softer dusty blue – secondary accent (detail border, category badges)
	clrSecond = lipgloss.Color("#89AABA")
	// light dusty blue – hyperlinks
	clrLink = lipgloss.Color("#9CC0D0")
)

// ─── Base styles ─────────────────────────────────────────────────────────────

var (
	headerStyle = lipgloss.NewStyle().Foreground(clrAccent).Bold(true)
	subStyle    = lipgloss.NewStyle().Foreground(clrDim) //nolint:unused

	activeTabSty = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#4D7A96")).
			Padding(0, 2)

	inactiveTabSty = lipgloss.NewStyle().
			Foreground(clrDim).
			Padding(0, 2)

	// Main content box – accent border
	contentBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrAccent).
			Padding(1, 2)

	// Detail panel box – secondary border
	detailBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrSecond).
			Padding(1, 2)

	sectionSty = lipgloss.NewStyle().Foreground(clrSection).Bold(true)
	labelSty   = lipgloss.NewStyle().Foreground(clrLabel).Bold(true)
	dimSty     = lipgloss.NewStyle().Foreground(clrDim)
	brightSty  = lipgloss.NewStyle().Foreground(clrBright).Bold(true)
	linkSty    = lipgloss.NewStyle().Foreground(clrLink).Underline(true)
	cursorSty  = lipgloss.NewStyle().Foreground(clrAccent).Bold(true)
	footerSty  = lipgloss.NewStyle().Foreground(clrDim)

	tagSty = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8EEF2")).
		Background(lipgloss.Color("#4D7A96")).
		Padding(0, 1)

	catSty = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8EEF2")).
		Background(lipgloss.Color("#5D8EA6")).
		Padding(0, 1)

	spiralSty = lipgloss.NewStyle().Foreground(clrAccent)
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

func osc8(url, text string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

func clickLink(url string) string {
	return osc8(url, linkSty.Render(url))
}

// ─── Spiral animation (dynamic size + proportional rings) ────────────────────

// renderSpiral draws an animated concentric-ring spiral scaled to fit rows×cols.
func renderSpiral(frame, rows, cols int) string {
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = make([]rune, cols)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	cx := float64(cols) / 2.0
	cy := float64(rows) / 2.0

	// Maximum radius that fits within the grid.
	// Horizontal positions are scaled ×2.1 to compensate for character aspect ratio.
	maxR := math.Min(cy*0.88, (cx/2.1)*0.88)
	if maxR < 1 {
		maxR = 1
	}

	type ring struct {
		frac  float64 // radius as fraction of maxR
		dots  int
		speed float64
		char  rune
	}

	// Six rings spread from centre to edge; scale with maxR automatically.
	rings := []ring{
		{0.15, 6,  0.080, '·'},
		{0.32, 11, 0.050, '•'},
		{0.50, 18, 0.030, '●'},
		{0.66, 26, 0.020, '•'},
		{0.82, 36, 0.013, '·'},
		{0.97, 48, 0.008, '·'},
	}

	for _, r := range rings {
		radius := r.frac * maxR
		phase := float64(frame) * r.speed
		for d := 0; d < r.dots; d++ {
			angle := (float64(d)/float64(r.dots))*2*math.Pi + phase
			x := cx + radius*2.1*math.Cos(angle)
			y := cy + radius*math.Sin(angle)
			xi := int(math.Round(x))
			yi := int(math.Round(y))
			if yi >= 0 && yi < rows && xi >= 0 && xi < cols {
				grid[yi][xi] = r.char
			}
		}
	}

	var sb strings.Builder
	for i, row := range grid {
		sb.WriteString(string(row))
		if i < rows-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// ─── Scroll helper ────────────────────────────────────────────────────────────

// applyScroll slices content to a window of maxLines starting at offset.
// If content is shorter than maxLines the result is padded with empty lines so
// the surrounding box always renders at a consistent height.
func applyScroll(content string, offset, maxLines int) string {
	lines := strings.Split(content, "\n")
	total := len(lines)

	if offset < 0 {
		offset = 0
	}
	if offset >= total {
		offset = total - 1
	}

	end := offset + maxLines
	if end > total {
		end = total
	}

	sliced := append([]string(nil), lines[offset:end]...)

	// Pad to maxLines so the box height stays constant while scrolling.
	for len(sliced) < maxLines {
		sliced = append(sliced, "")
	}

	return strings.Join(sliced, "\n")
}

// ─── Available content lines ──────────────────────────────────────────────────

// contentAvailLines returns how many lines of inner content the box can display.
func (m Model) contentAvailLines() int {
	headerH := 5 // 4-line ASCII art + 1 blank breathing line
	if m.width < 55 {
		headerH = 1
	}
	// Overhead: header + 3 JoinVertical separators + tabbar(1) + footer(1)
	//           + box border top+bottom(2) + box padding top+bottom(2)
	avail := m.height - headerH - 3 - 1 - 1 - 2 - 2
	if avail < 5 {
		avail = 5
	}
	return avail
}

// ─── Top-level view ──────────────────────────────────────────────────────────

func (m Model) View() string {
	if m.width < 61 || m.height < 30 {
		return lipgloss.NewStyle().Foreground(clrLabel).Render(
			"\n  [ terminal too small — please resize to at least 61×30 ]",
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
		return headerStyle.Render("Nourin Awad")
	}
	ascii := ` ▗▖  ▗▖ ▄▄▄  █  ▐▌ ▄▄▄ ▄ ▄▄▄▄       ▗▄▖ ▄   ▄ ▗▞▀▜▌▐▌▄
 ▐▛▚▖▐▌█   █ ▀▄▄▞▘█    ▄ █   █     ▐▌ ▐▌█ ▄ █ ▝▚▄▟▌▐▌  
 ▐▌ ▝▜▌▀▄▄▄▀      █    █ █   █     ▐▛▀▜▌█▄█▄█   ▗▞▀▜▌  
 ▐▌  ▐▌                █           ▐▌ ▐▌        ▝▚▄▟▌  `
	return headerStyle.Render(ascii) + "\n"
}

// ─── Tab bar ──────────────────────────────────────────────────────────────────

func (m Model) renderTabBar() string {
	names := m.tabs
	if m.width < 70 {
		names = []string{"About", "Proj", "Side", "Skills", "Exp", "Contact"}
	}
	if m.width < 54 {
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

const sidebarW = 26 // spiral sidebar fixed character width

func (m Model) renderMiddle() string {
	boxSty := contentBoxSty
	if m.detailOpen {
		boxSty = detailBoxSty
	}

	avail := m.contentAvailLines()

	if m.width < 72 {
		// Narrow layout: no spiral, content fills full width.
		inner := m.width - 8
		if inner < 10 {
			inner = 10
		}
		full := m.renderContent(inner)
		shown := applyScroll(full, m.scrollOffset, avail)
		return boxSty.Width(m.width - 2).Render(shown)
	}

	// Wide layout: spiral on left, scrollable content on right.
	contentW := m.width - sidebarW - 2
	inner := contentW - 8
	if inner < 10 {
		inner = 10
	}

	full := m.renderContent(inner)
	shown := applyScroll(full, m.scrollOffset, avail)

	// Make spiral exactly as tall as the content box.
	//   content box total height = avail + border(2) + padding(2) = avail + 4
	//   spiral panel total height = spiralRows + padding top+bottom(2)
	//   → spiralRows = avail + 2
	spiralRows := avail + 2
	if spiralRows < 8 {
		spiralRows = 8
	}
	const spiralCols = 22 // fits within sidebarW(26) with padding(1,1)

	spiral := spiralSty.
		Width(sidebarW).
		Padding(1, 1).
		Render(renderSpiral(m.frame, spiralRows, spiralCols))

	content := boxSty.Width(contentW).Render(shown)
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
		lipgloss.NewStyle().Foreground(clrSection).Render("GPA: 3.95 / 4.0") + "\n\n")

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
		b.WriteString("    " + lipgloss.NewStyle().Foreground(clrSecond).Render(p.tech) + "\n\n")
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
	b.WriteString(lipgloss.NewStyle().Foreground(clrSecond).Render(p.tech) + "\n\n")

	b.WriteString(sectionSty.Render("About") + "\n")
	b.WriteString("  " + p.short + "\n\n")

	b.WriteString(sectionSty.Render("Highlights") + "\n")
	for _, bullet := range p.bullets {
		b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
	}

	b.WriteString("\n" + sectionSty.Render("Repository") + "\n")
	b.WriteString("  " + clickLink(p.github) + "\n")
	b.WriteString("  " + dimSty.Render("ctrl+click to open in browser (iTerm2, WezTerm, etc.)") + "\n")

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
		"  ctrl+click any link to open  ·  iTerm2, WezTerm, Kitty, etc.",
	) + "\n")

	return b.String()
}

// ─── Footer ───────────────────────────────────────────────────────────────────

func (m Model) renderFooter() string {
	var hint string
	switch {
	case m.detailOpen:
		hint = "↑↓ scroll   esc back   q quit"
	case m.activeTab == 1 || m.activeTab == 2:
		hint = "← → tabs   ↑↓ select   enter open   q quit"
	default:
		hint = "← → tabs   ↑↓ scroll   q quit"
	}

	suffix := ""
	if m.scrollOffset > 0 {
		suffix = fmt.Sprintf("   [+%d]", m.scrollOffset)
	}

	return footerSty.Render("  " + hint + suffix)
}