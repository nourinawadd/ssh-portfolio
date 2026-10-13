package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	clrAccent  = lipgloss.Color("#7BADC0")
	clrSection = lipgloss.Color("#B8C8D4")
	clrLabel   = lipgloss.Color("#9AACB8")
	clrDim     = lipgloss.Color("#6A7580")
	clrBright  = lipgloss.Color("#DCE4EA")
	clrSecond  = lipgloss.Color("#89AABA")
	clrLink    = lipgloss.Color("#9CC0D0")
	clrBg      = lipgloss.Color("#000000")
)

var (
	headerStyle = lipgloss.NewStyle().Foreground(clrAccent).Background(clrBg).Bold(true)

	activeTabSty = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#4D7A96")).
			Padding(0, 2)

	inactiveTabSty = lipgloss.NewStyle().
			Foreground(clrDim).
			Background(clrBg).
			Padding(0, 2)

	contentBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrAccent).
			Background(clrBg).
			Padding(1, 2)

	detailBoxSty = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(clrSecond).
			Background(clrBg).
			Padding(1, 2)

	sectionSty = lipgloss.NewStyle().Foreground(clrSection).Background(clrBg).Bold(true)
	labelSty   = lipgloss.NewStyle().Foreground(clrLabel).Background(clrBg).Bold(true)
	dimSty     = lipgloss.NewStyle().Foreground(clrDim).Background(clrBg)
	brightSty  = lipgloss.NewStyle().Foreground(clrBright).Background(clrBg).Bold(true)
	linkSty    = lipgloss.NewStyle().Foreground(clrLink).Background(clrBg).Underline(true)
	cursorSty  = lipgloss.NewStyle().Foreground(clrAccent).Background(clrBg).Bold(true)
	footerSty  = lipgloss.NewStyle().Foreground(clrDim).Background(clrBg)
	tagSty     = lipgloss.NewStyle().Foreground(clrSecond).Background(clrBg)
	catSty     = lipgloss.NewStyle().Foreground(clrAccent).Background(clrBg).Bold(true)
	starSty    = lipgloss.NewStyle().Foreground(clrAccent).Background(clrBg)
)

type project struct {
	name    string
	tech    string
	short   string
	bullets []string
	live    string
	github  string
}

type sideProject struct {
	name     string
	category string
	short    string
	bullets  []string
	link     string
	github   string
	ig       string
	tags     []string
	games    []sideProject
}

var projects = []project{
	{
		name:  "Anchor",
		tech:  "React Native  ·  Node  ·  MongoDB",
		short: "NFC focus app with app blocking and AI insights",
		bullets: []string{
			"React Native focus sessions with NFC check-ins, streaks, and goals",
			"Blocks distracting apps with iOS Screen Time and Live Activities",
			"Node/Express + MongoDB backend with AI-generated focus insights",
		},
		live:   "https://anchor-iesq.onrender.com/",
		github: "https://github.com/nourinawadd/anchor",
	},
	{
		name:  "nourin-drive",
		tech:  "Next.js  ·  React  ·  Node  ·  MongoDB",
		short: "Amiga Workbench-style personal portfolio",
		bullets: []string{
			"An interactive personal portfolio styled as an Amiga Workbench desktop",
			"Built with Next.js and React, with Node.js and MongoDB in the stack",
		},
		live: "https://nourin.is-a.dev/",
	},
	{
		name:  "SSH Portfolio",
		tech:  "Go  ·  Bubble Tea  ·  Wish",
		short: "Terminal portfolio served over SSH",
		bullets: []string{
			"Interactive terminal portfolio built with Bubble Tea and Wish",
			"Connect with: ssh nourin.is-a.dev -p 2323",
		},
		live:   "https://nourin.is-a.dev/terminal",
		github: "https://github.com/nourinawadd/ssh-portfolio",
	},
	{
		name:  "Tether Note",
		tech:  "Node  ·  MongoDB  ·  React",
		short: "Cron-scheduled notes for your future self",
		bullets: []string{
			"Scheduled note delivery for messages to your future self",
			"Node.js cron jobs and MongoDB automate time-based delivery",
			"JWT session management for note operations",
		},
		live: "https://tethernote.vercel.app/",
	},
	{
		name:  "HOLD",
		tech:  ".NET 10  ·  Blazor Server  ·  EF Core  ·  PostgreSQL",
		short: "Wishlist with a wait-before-buy timer",
		bullets: []string{
			"Save products and set a waiting period before buying",
			"Blazor Server app with EF Core and PostgreSQL",
			"Product scraper extracts item details from pasted links",
		},
		live:   "https://hold.nourin.me/",
		github: "https://github.com/nourinawadd/hold",
	},
}

var sideProjects = []sideProject{
	{
		name:     "Indie Game Projects",
		category: "Game Dev",
		short:    "Open to browse five individual indie games and game-jam projects",
		bullets: []string{
			"Browse each game for its description, technology, and links",
		},
		games: []sideProject{
			{
				name:     "Bomb Voyage",
				category: "Game Dev",
				short:    "Four-player bomb minigames in a shrinking arena",
				bullets: []string{
					"Four-player party game made for the VICTORIS 5.0 jam",
					"Solve minigames on your bomb to buy time while dodging opponents",
					"The arena shrinks as the match goes on",
				},
				link:   "https://abdelrahman-hosny.itch.io/bomb-voyage",
				github: "https://github.com/Abdelrahmanhr/Bomb-Voyage",
				tags:   []string{"Godot"},
			},
			{
				name:     "AstraCipher",
				category: "Game Dev",
				short:    "Top-down platformer focused on dark-web awareness",
				bullets: []string{
					"2D story-driven top-down platformer for dark-web awareness",
					"Uses scene instancing, event-driven architecture, and AI-driven NPCs",
					"Includes state machines and pathfinding",
				},
				link: "https://sifrstudios.itch.io/astracipher",
				tags: []string{"Godot 4.4", "GDScript", "WebAssembly"},
			},
			{
				name:     "Suhoor Rush",
				category: "Game Dev",
				short:    "Rapid item-collecting web game with a narrative",
				bullets: []string{
					"Built in 3 days for a 'Collect Them All' jam",
					"UI Toolkit menus, Scriptable Objects, and C# events/delegates",
					"Object pooling, sprite atlasing, and WebGL optimization",
				},
				link: "https://sifrstudios.itch.io/suhoor-rush",
				tags: []string{"Unity", "C#", "WebGL"},
			},
			{
				name:     "Alias",
				category: "Game Dev",
				short:    "Detective game: narrow celebrity suspects by clues",
				bullets: []string{
					"Read witness reports about strange incidents at high-profile events",
					"Filter celebrity suspects with clues about appearance and background",
					"Build a shortlist and report the culprit to solve each case",
				},
				link: "https://sifrstudios.itch.io/alias",
				tags: []string{"Unity"},
			},
			{
				name:     "Puppetry",
				category: "Game Dev",
				short:    "Power-themed puzzle about coordinating puppets",
				bullets: []string{
					"2D puzzle for a power-themed jam, built in 9 days",
					"Coordinate multiple puppets and interact with the environment",
					"Timed disintegration mechanics",
				},
				link: "https://sifrstudios.itch.io/puppetry",
				tags: []string{"Godot 4.3"},
			},
		},
	},
	{
		name:     "Behance Portfolio",
		category: "Graphic Design",
		short:    "Visual design work: branding, illustration, animation",
		bullets: []string{
			"Branding, logo design, and visual identity projects",
			"3D modelling and animation projects",
			"Digital illustrations and creative compositions",
		},
		link: "https://www.behance.net/nourinawadd",
		ig:   "https://www.instagram.com/diarydump.jpg/",
		tags: []string{"Graphic Design", "Branding", "UI/UX", "Illustration"},
	},
}

func osc8(url, text string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

func clickLink(label, url string) string {
	return osc8(url, linkSty.Render(label))
}

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
	for len(sliced) < maxLines {
		sliced = append(sliced, "")
	}
	return strings.Join(sliced, "\n")
}

func (m Model) contentAvailLines() int {
	headerH := 5
	if m.width < 55 {
		headerH = 1
	}
	avail := m.height - headerH - 3 - 1 - 1 - 2 - 2
	if avail < 5 {
		avail = 5
	}
	return avail
}

func (m Model) View() string {
	if m.width < 61 || m.height < 30 {
		content := lipgloss.NewStyle().Foreground(clrLabel).Background(clrBg).Render(
			"\n  [ terminal too small — please resize to at least 61×30 ]",
		)
		return m.placeOnBlack(content)
	}
	content := lipgloss.JoinVertical(lipgloss.Left,
		m.renderHeader(),
		m.renderTabBar(),
		m.renderMiddle(),
		m.renderFooter(),
	)
	return m.placeOnBlack(content)
}

func (m Model) placeOnBlack(content string) string {
	content = lipgloss.NewStyle().Background(clrBg).Render(content)
	return lipgloss.Place(
		m.width, m.height, lipgloss.Left, lipgloss.Top, content,
		lipgloss.WithWhitespaceBackground(clrBg),
	)
}

func (m Model) renderHeader() string {
	if m.width < 55 {
		return headerStyle.Render("Nourin Awad")
	}
	ascii := ` ▗▖  ▗▖ ▄▄▄  █  ▐▌ ▄▄▄ ▄ ▄▄▄▄       ▗▄▖ ▄   ▄ ▗▞▀▜▌▐▌
 ▐▛▚▖▐▌█   █ ▀▄▄▞▘█    ▄ █   █     ▐▌ ▐▌█ ▄ █ ▝▚▄▟▌▐▌  
 ▐▌ ▝▜▌▀▄▄▄▀      █    █ █   █     ▐▛▀▜▌█▄█▄█   ▗▞▀▜▌  
 ▐▌  ▐▌                █           ▐▌ ▐▌        ▝▚▄▟▌  `
	return headerStyle.Render(ascii) + "\n"
}

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

const sidebarW = 26

func (m Model) renderMiddle() string {
	boxSty := contentBoxSty
	if m.detailOpen {
		boxSty = detailBoxSty
	}

	avail := m.contentAvailLines()

	if m.width < 72 {
		inner := m.width - 8
		if inner < 10 {
			inner = 10
		}
		full := m.renderContent(inner)
		shown := applyScroll(full, m.visibleListOffset(avail), avail)
		return boxSty.Width(m.width - 2).Render(shown)
	}

	contentW := m.width - sidebarW - 2
	inner := contentW - 8
	if inner < 10 {
		inner = 10
	}

	full := m.renderContent(inner)
	shown := applyScroll(full, m.visibleListOffset(avail), avail)

	starRows := avail + 2
	if starRows < 8 {
		starRows = 8
	}
	const starCols = 22

	star := starSty.
		Width(sidebarW).
		Padding(1, 1).
		Render(RenderStar(m.frame, starRows, starCols))

	content := boxSty.Width(contentW).Render(shown)
	return lipgloss.JoinHorizontal(lipgloss.Top, star, content)
}

func (m Model) visibleListOffset(avail int) int {
	offset := m.scrollOffset
	if m.detailOpen || !m.isListTab() {
		return offset
	}

	line := 3 + 3*m.cursor
	if m.activeTab == 1 {
		line = 3 + 4*m.cursor
	} else if m.gameListOpen {
		line = 3 + 3*m.gameCursor
	}

	if line < offset {
		return line
	}
	if line >= offset+avail {
		return line - avail + 1
	}
	return offset
}

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

func (m Model) renderAbout(w int) string {
	var b strings.Builder

	b.WriteString(sectionSty.Render("◆ About Me") + "\n\n")
	writeParagraph(&b, w, "Software engineer based in Mansoura, Egypt. I build scalable systems and clean APIs, primarily with Node.js.")
	b.WriteString("\n")
	writeParagraph(&b, w, "Studying Communications & Computer Engineering at Mansoura University.")
	b.WriteString("\n")

	b.WriteString(sectionSty.Render("◆ Education") + "\n\n")
	b.WriteString("  " + labelSty.Render("Mansoura University") + "\n")
	b.WriteString("  BE · Communications & Computer Engineering\n")
	b.WriteString("  " + dimSty.Render("Sep 2022 – Present") + "   " +
		lipgloss.NewStyle().Foreground(clrSection).Background(clrBg).Render("GPA: 3.95 / 4.0") + "\n")
	writeParagraph(&b, w, "Coursework: Data Structures, Operating Systems, Distributed Systems, AI, Calculus")
	b.WriteString("\n")
	b.WriteString("  " + labelSty.Render("Mansoura College International Schools") + "\n")
	b.WriteString("  IGCSE · 99%  " + dimSty.Render("Sep 2020 – Sep 2022") + "\n\n")

	b.WriteString(sectionSty.Render("◆ Extracurriculars") + "\n\n")
	extracurriculars := []struct {
		organization string
		role         string
		period       string
		highlights   []string
	}{
		{
			organization: "Women in Games",
			role:         "Ambassador",
			period:       "Apr 2026 – Present",
		},
		{
			organization: "IEEE Mansoura Student Branch · Victoris 3.0",
			role:         "Technical Director",
			period:       "Jun – Sep 2024",
			highlights: []string{
				"Led a national event: 880+ participants, 200+ teams",
				"Coordinated technical delivery across four tracks",
			},
		},
		{
			organization: "IEEE Mansoura Computer Society Chapter",
			role:         "Backend Committee Member",
			period:       "Dec 2024 – Present",
			highlights: []string{
				"Built MERN app APIs, database, and server-side logic",
				"Recognized as one of the chapter's most active volunteers",
			},
		},
		{
			organization: "CATReloaded",
			role:         "Game Developer",
			period:       "Dec 2024 – Present",
			highlights: []string{
				"Applied Unity 2D/3D, mechanics, physics, and UI",
				"Built prototypes in Unity/Godot during 48-hour game jams",
			},
		},
		{
			organization: "IEEE Mansoura Student Branch",
			role:         "Game Development Competition Coordinator",
			period:       "May – Sep 2025",
		},
		{
			organization: "IEEE Mansoura Student Branch",
			role:         "Multimedia Volunteer",
			period:       "Dec 2024 – Oct 2025",
			highlights: []string{
				"Led the Advanced Graphic Design Team for Career Insider 6.0",
				"Recognized as one of the branch's most active volunteers",
			},
		},
		{
			organization: "IEEE Mansoura Computer Society Chapter",
			role:         "Embedded Systems Volunteer",
			period:       "Feb – Dec 2024",
			highlights: []string{
				"Co-designed and taught six Code Steps 4.0 sessions",
				"Created activities for children aged 9–14 and led the final competition",
			},
		},
	}
	for _, item := range extracurriculars {
		b.WriteString("  " + labelSty.Render(item.organization) + "\n")
		b.WriteString("  " + brightSty.Render(item.role) + "  " + dimSty.Render(item.period) + "\n")
		for _, highlight := range item.highlights {
			b.WriteString("  " + cursorSty.Render("▸") + " " + highlight + "\n")
		}
		b.WriteString("\n")
	}

	b.WriteString(sectionSty.Render("◆ Languages") + "\n\n")
	sep := dimSty.Render("  ·  ")
	b.WriteString("  " + tagSty.Render("Arabic · Native") + sep +
		tagSty.Render("English · Full Professional Proficiency") + sep +
		tagSty.Render("German · A2") + "\n")

	return b.String()
}

func writeParagraph(b *strings.Builder, width int, text string) {
	lineWidth := width - 2
	if lineWidth < 1 {
		lineWidth = 1
	}
	lines := strings.Split(lipgloss.NewStyle().Width(lineWidth).Background(clrBg).Render(text), "\n")
	for _, line := range lines {
		b.WriteString("  " + line + "\n")
	}
}

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

		nameSty := lipgloss.NewStyle().Foreground(clrBright).Background(clrBg)
		if selected {
			nameSty = nameSty.Bold(true)
		}

		b.WriteString(pfx + nameSty.Render(p.name) + "\n")
		b.WriteString("    " + dimSty.Render(p.short) + "\n")
		b.WriteString("    " + lipgloss.NewStyle().Foreground(clrSecond).Background(clrBg).Render(p.tech) + "\n\n")
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
	b.WriteString(lipgloss.NewStyle().Foreground(clrSecond).Background(clrBg).Render(p.tech) + "\n\n")

	b.WriteString(sectionSty.Render("About") + "\n")
	b.WriteString("  " + p.short + "\n\n")

	b.WriteString(sectionSty.Render("Highlights") + "\n")
	for _, bullet := range p.bullets {
		b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
	}

	b.WriteString("\n" + sectionSty.Render("Links") + "\n")
	if p.live != "" {
		b.WriteString("  " + labelSty.Render("Live:  ") + clickLink(p.live, p.live) + "\n")
	}
	if p.github != "" {
		b.WriteString("  " + labelSty.Render("Repo:  ") + clickLink(p.github, p.github) + "\n")
	}
	b.WriteString("  " + dimSty.Render("ctrl+click to open  ·  iTerm2, WezTerm, Kitty") + "\n")

	return b.String()
}

func (m Model) renderSideProjects(w int) string {
	if m.gameListOpen {
		if m.detailOpen {
			return m.renderGameDetail(m.gameCursor)
		}
		return m.renderGameList(w)
	}

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

		nameSty := lipgloss.NewStyle().Foreground(clrBright).Background(clrBg)
		if selected {
			nameSty = nameSty.Bold(true)
		}

		cat := catSty.Render(sp.category)
		b.WriteString(pfx + nameSty.Render(sp.name) + "  " + cat + "\n")
		b.WriteString("    " + dimSty.Render(sp.short) + "\n\n")
	}

	return b.String()
}

func (m Model) renderGameList(w int) string {
	games := m.currentGames()
	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Indie Game Projects") + "\n")
	b.WriteString(dimSty.Render("  ↑↓ select   enter open   esc back") + "\n\n")

	for i, game := range games {
		selected := i == m.gameCursor
		pfx := "  "
		if selected {
			pfx = cursorSty.Render("▶ ")
		}

		nameSty := lipgloss.NewStyle().Foreground(clrBright).Background(clrBg)
		if selected {
			nameSty = nameSty.Bold(true)
		}

		b.WriteString(pfx + nameSty.Render(game.name) + "\n")
		b.WriteString("    " + dimSty.Render(game.short) + "\n\n")
	}

	return b.String()
}

func (m Model) renderSideProjectDetail(idx int) string {
	if idx < 0 || idx >= len(sideProjects) {
		return ""
	}
	sp := sideProjects[idx]
	return m.renderSideProjectDetails(sp, "esc · back to side projects")
}

func (m Model) renderGameDetail(idx int) string {
	games := m.currentGames()
	if idx < 0 || idx >= len(games) {
		return ""
	}
	return m.renderSideProjectDetails(games[idx], "esc · back to games")
}

func (m Model) renderSideProjectDetails(sp sideProject, backHint string) string {
	var b strings.Builder
	b.WriteString(dimSty.Render(backHint) + "\n\n")

	b.WriteString(labelSty.Render(sp.name) + "  " + catSty.Render(sp.category) + "\n\n")

	b.WriteString(sectionSty.Render("About") + "\n")
	b.WriteString("  " + sp.short + "\n\n")

	b.WriteString(sectionSty.Render("Highlights") + "\n")
	for _, bullet := range sp.bullets {
		b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
	}

	sep := dimSty.Render("  ·  ")
	var parts []string
	for _, tag := range sp.tags {
		parts = append(parts, tagSty.Render(tag))
	}
	b.WriteString("\n" + sectionSty.Render("Tags") + "\n")
	b.WriteString("  " + strings.Join(parts, sep) + "\n")

	b.WriteString("\n" + sectionSty.Render("Links") + "\n")
	if sp.link != "" {
		b.WriteString("  " + labelSty.Render("Site:  ") + clickLink(sp.link, sp.link) + "\n")
	}
	if sp.github != "" {
		b.WriteString("  " + labelSty.Render("Repo:  ") + clickLink(sp.github, sp.github) + "\n")
	}
	if sp.ig != "" {
		b.WriteString("  " + labelSty.Render("Insta: ") + clickLink(sp.ig, sp.ig) + "\n")
	}
	b.WriteString("  " + dimSty.Render("ctrl+click to open  ·  iTerm2, WezTerm, Kitty") + "\n")

	return b.String()
}

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

	sep := dimSty.Render("  ·  ")
	for _, s := range sections {
		b.WriteString("  " + labelSty.Render(s.label) + "\n")
		var parts []string
		for _, item := range s.items {
			parts = append(parts, tagSty.Render(item))
		}
		b.WriteString("  " + strings.Join(parts, sep) + "\n\n")
	}

	b.WriteString("  " + labelSty.Render("Certifications") + "\n")
	certifications := []struct {
		name   string
		issuer string
	}{
		{"Backend Development and APIs", "freeCodeCamp"},
		{"Relational Database", "freeCodeCamp"},
		{"Foundational C# with Microsoft", "freeCodeCamp"},
		{"Project Management Foundations", "LinkedIn"},
		{"Time Management Fundamentals", "LinkedIn"},
		{"Python Programming", "MaharaTech"},
	}
	for _, cert := range certifications {
		b.WriteString("  " + tagSty.Render(cert.name) + dimSty.Render(" · "+cert.issuer) + "\n")
	}

	return b.String()
}

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
			"Jul – Aug 2025",
			"Mansoura, Egypt",
			[]string{
				"Designed SQL Server schemas, stored procedures, and multi-table joins",
				"Built ASP.NET MVC apps in C# with Razor views, controllers, and routing",
				"Integrated databases with Entity Framework, LINQ, migrations, and CRUD",
			},
		},
		{
			"National Telecommunications Institute (NTI)",
			"MEAN Stack Intern",
			"Jun – Jul 2025",
			"Mansoura, Egypt",
			[]string{
				"Built REST APIs with JWT auth and MongoDB/Mongoose CRUD",
				"Developed Angular SPAs with routing, reactive forms, and state management",
				"Applied API security, testing, and performance optimization",
			},
		},
		{
			"Nile University",
			"Blockchain Alternatives Research Intern",
			"Aug – Sep 2025",
			"Cairo, Egypt",
			[]string{
				"Researched scalability alternatives: Lightning Network, sharding, DAGs, and sidechains",
				"Studied layer-two protocols and blockchain scalability approaches",
			},
		},
	}

	var b strings.Builder
	b.WriteString(sectionSty.Render("◆ Professional Experience") + "\n\n")

	for _, j := range jobs {
		b.WriteString("  " + labelSty.Render(j.company) + "\n")
		b.WriteString("  " + brightSty.Render(j.role) + "  " + dimSty.Render(j.period) + "\n")
		b.WriteString("  " + dimSty.Render(j.location) + "\n")
		for _, bullet := range j.bullets {
			b.WriteString("  " + cursorSty.Render("▸") + " " + bullet + "\n")
		}
		b.WriteString("\n")
	}

	return b.String()
}

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

func (m Model) renderFooter() string {
	var hint string
	switch {
	case m.detailOpen:
		hint = "↑↓ scroll   esc back   q quit"
	case m.gameListOpen:
		hint = "↑↓ select   enter open   esc back   q quit"
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
