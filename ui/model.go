package ui
import tea "github.com/charmbracelet/bubbletea"
type Model struct {
    width      int
    height     int
    activeTab  int      
    tabs       []string 
    cursor     int  
}

func NewModel(width, height int) Model {
    return Model{
        width:     width,
        height:    height,
        activeTab: 0,
        tabs:      []string{"About", "Projects", "Skills", "Contact"},
        cursor:    0,
    }
}
func (m Model) Init() tea.Cmd {
    return nil
}