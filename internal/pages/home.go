package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/showcase/internal/config"
)

func NewHomePage(cfg config.Config) HomePage{
	return HomePage{
    userName: cfg.UserName,
    introText: cfg.IntroText,
    projectList: cfg.ProjectList,
		selectedProject: 0,
	}
}

func (hp HomePage) Init() tea.Cmd {
    return nil
}

func (hp HomePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
					return hp, tea.Quit

				case "k", "up":
					if hp.selectedProject > 0 {
						hp.selectedProject--
					}

				case "j":
					if hp.selectedProject < len(hp.projectList)-1 {
						hp.selectedProject++
					}
        }
    }
    return hp, nil
}

func (hp HomePage) View() string {
  s := fmt.Sprintf("%v's Showcase\n\n", hp.userName)
  s += fmt.Sprintf("%v\n\n", hp.introText)

	for i, project := range hp.projectList {
		if hp.selectedProject == i {
			s += "> "
		} else {
			s += "- "
		}
		s += fmt.Sprintf("%v\n", project.Name)
	}

  s += "\n(press 'q' to quit)"

  return s 

}
