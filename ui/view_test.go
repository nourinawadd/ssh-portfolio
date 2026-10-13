package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestMainProjectsMatchPortfolioList(t *testing.T) {
	want := []string{"Anchor", "nourin-drive", "SSH Portfolio", "Tether Note", "HOLD"}
	if len(projects) != len(want) {
		t.Fatalf("got %d projects, want %d", len(projects), len(want))
	}
	for i, name := range want {
		if projects[i].name != name {
			t.Errorf("projects[%d].name = %q, want %q", i, projects[i].name, name)
		}
	}
}

func TestAboutShowsAllCVExtracurriculars(t *testing.T) {
	content := NewModel(100, 40).renderAbout(80)
	want := []string{
		"Women in Games",
		"IEEE Mansoura Student Branch · Victoris 3.0",
		"Backend Committee Member",
		"CATReloaded",
		"Game Development Competition Coordinator",
		"Multimedia Volunteer",
		"Embedded Systems Volunteer",
		"Code Steps 4.0",
	}
	for _, text := range want {
		if !strings.Contains(content, text) {
			t.Errorf("About extracurricular content missing %q", text)
		}
	}
}

func TestAboutShowsEducationCourseworkAndLanguages(t *testing.T) {
	content := NewModel(100, 40).renderAbout(80)
	want := []string{
		"Mansoura University",
		"GPA: 3.95 / 4.0",
		"Data Structures",
		"Mansoura College International Schools",
		"IGCSE · 99%",
		"Arabic · Native",
		"English · Full Professional Proficiency",
		"German · A2",
	}
	for _, text := range want {
		if !strings.Contains(content, text) {
			t.Errorf("About content missing %q", text)
		}
	}
}

func TestAboutParagraphsWrapOnWordsWithoutForcedSentenceBreaks(t *testing.T) {
	content := NewModel(140, 40).renderAbout(120)
	for _, phrase := range []string{
		"Software engineer based in Mansoura, Egypt. I build scalable systems and clean APIs, primarily with Node.js.",
		"Studying Communications & Computer Engineering at Mansoura University.",
		"Coursework: Data Structures, Operating Systems, Distributed Systems, AI, Calculus",
	} {
		if !strings.Contains(content, phrase) {
			t.Errorf("wide About paragraph was unexpectedly broken: %q", phrase)
		}
	}
}

func TestSkillsShowAllCVCertifications(t *testing.T) {
	content := NewModel(100, 40).renderSkills(80)
	want := []string{
		"Certifications",
		"Backend Development and APIs",
		"Relational Database",
		"Foundational C# with Microsoft",
		"Project Management Foundations",
		"Time Management Fundamentals",
		"Python Programming",
		"freeCodeCamp",
		"LinkedIn",
		"MaharaTech",
	}
	for _, text := range want {
		if !strings.Contains(content, text) {
			t.Errorf("Skills content missing %q", text)
		}
	}
}

func TestExperienceMatchesCVEntries(t *testing.T) {
	content := NewModel(100, 40).renderExperience(80)
	want := []string{
		"Information Technology Institute (ITI)",
		"Jul – Aug 2025",
		"National Telecommunications Institute (NTI)",
		"Jun – Jul 2025",
		"Nile University",
		"Blockchain Alternatives Research Intern",
		"Aug – Sep 2025",
		"Lightning Network",
	}
	for _, text := range want {
		if !strings.Contains(content, text) {
			t.Errorf("Experience content missing %q", text)
		}
	}
}

func TestGameProjectsAreNestedUnderIndieGames(t *testing.T) {
	if len(sideProjects) != 2 {
		t.Fatalf("got %d top-level side projects, want 2", len(sideProjects))
	}
	if sideProjects[0].name != "Indie Game Projects" {
		t.Fatalf("first side project = %q, want Indie Game Projects", sideProjects[0].name)
	}
	if sideProjects[1].name != "Behance Portfolio" {
		t.Fatalf("second side project = %q, want Behance Portfolio", sideProjects[1].name)
	}

	wantGames := []string{"Bomb Voyage", "AstraCipher", "Suhoor Rush", "Alias", "Puppetry"}
	games := sideProjects[0].games
	if len(games) != len(wantGames) {
		t.Fatalf("got %d games, want %d", len(games), len(wantGames))
	}
	for i, name := range wantGames {
		if games[i].name != name {
			t.Errorf("games[%d].name = %q, want %q", i, games[i].name, name)
		}
	}
}

func TestNestedGameNavigationAndBackStack(t *testing.T) {
	m := NewModel(100, 40)
	m.activeTab = 2

	if got := m.View(); strings.Contains(got, "Bomb Voyage") {
		t.Fatal("game entries should not appear in the top-level Side Projects list")
	}

	m = updateModel(m, tea.KeyMsg{Type: tea.KeyEnter})
	if !m.gameListOpen || m.detailOpen {
		t.Fatal("opening Indie Game Projects should show the nested game list")
	}
	if got := m.View(); !strings.Contains(got, "Bomb Voyage") || !strings.Contains(got, "Alias") {
		t.Fatal("nested game list should show all individual games")
	}

	m = updateModel(m, tea.KeyMsg{Type: tea.KeyRight})
	if m.activeTab != 2 {
		t.Fatal("tab navigation should remain within the nested game list")
	}

	m = updateModel(m, tea.KeyMsg{Type: tea.KeyDown})
	if m.gameCursor != 1 {
		t.Fatalf("game cursor = %d, want 1", m.gameCursor)
	}
	m = updateModel(m, tea.KeyMsg{Type: tea.KeyEnter})
	if !m.detailOpen {
		t.Fatal("opening a game should show its detail")
	}
	if got := m.View(); !strings.Contains(got, "AstraCipher") ||
		!strings.Contains(got, "Godot 4.4") ||
		!strings.Contains(got, "https://sifrstudios.itch.io/astracipher") {
		t.Fatal("game detail should show its name, stack, and itch link")
	}

	m = updateModel(m, tea.KeyMsg{Type: tea.KeyEsc})
	if m.detailOpen || !m.gameListOpen || m.gameCursor != 1 || m.cursor != 0 {
		t.Fatal("back from game detail should restore the selected game and parent")
	}
	m = updateModel(m, tea.KeyMsg{Type: tea.KeyEsc})
	if m.gameListOpen || m.cursor != 0 {
		t.Fatal("back from the game list should restore the parent selection")
	}
	if got := m.View(); strings.Contains(got, "Bomb Voyage") || !strings.Contains(got, "Behance Portfolio") {
		t.Fatal("back should return to the top-level Side Projects list")
	}
}

func TestLastMainProjectRemainsVisibleWhenSelected(t *testing.T) {
	m := NewModel(100, 30)
	m.activeTab = 1

	for i := 0; i < len(projects)-1; i++ {
		m = updateModel(m, tea.KeyMsg{Type: tea.KeyDown})
	}

	if m.cursor != len(projects)-1 {
		t.Fatalf("project cursor = %d, want %d", m.cursor, len(projects)-1)
	}
	if got := m.View(); !strings.Contains(got, "HOLD") {
		t.Fatal("selected last project should remain visible in the list")
	}
}

func TestLastGameRemainsVisibleWhenSelected(t *testing.T) {
	m := NewModel(61, 30)
	m.activeTab = 2
	m = updateModel(m, tea.KeyMsg{Type: tea.KeyEnter})

	for i := 0; i < len(sideProjects[0].games)-1; i++ {
		m = updateModel(m, tea.KeyMsg{Type: tea.KeyDown})
	}

	if m.gameCursor != len(sideProjects[0].games)-1 {
		t.Fatalf("game cursor = %d, want %d", m.gameCursor, len(sideProjects[0].games)-1)
	}
	if got := m.View(); !strings.Contains(got, "Alias") {
		t.Fatal("selected last game should remain visible in the list")
	}
}

func updateModel(m Model, msg tea.Msg) Model {
	updated, _ := m.Update(msg)
	return updated.(Model)
}
