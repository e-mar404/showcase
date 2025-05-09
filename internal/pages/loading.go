package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/showcase/internal/config"
)

type LoadingPage struct {}

func NewLoadingPage(_project config.Project) LoadingPage {
	return LoadingPage{}
}

func (lp LoadingPage) Init() tea.Cmd {
	return nil
}

func (lp LoadingPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return lp, tea.Quit
			}
	}
	return lp, nil
}

func (lp LoadingPage) View() string {
	return "initializing project...\n\n" 
}
