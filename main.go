package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

type config struct {
	UserName string `json:"user_name"`
	IntroText string `json:"intro_text"`
	ProjectList []project `json:"project_list"`
}

type model struct {
  userName string
  introText string
  projectList []project
	selectedProject int
}

type project struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func main() {
	file, err := os.Open(".showcase.json")
	if err != nil {
		log.Fatal(err)
    os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	var cfg config
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
    os.Exit(1)
	}

  p := tea.NewProgram(initialModel(cfg))
  if _, err = p.Run(); err != nil {
		log.Fatal(err)
    os.Exit(1)
  }
}

func initialModel(cfg config) model {
	return model{
    userName: cfg.UserName,
    introText: cfg.IntroText,
    projectList: cfg.ProjectList,
		selectedProject: 0,
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
				case "enter":
					open(m.projectList[m.selectedProject].Url)

        case "ctrl+c", "q":
					return m, tea.Quit

				case "k", "up":
					if m.selectedProject > 0 {
						m.selectedProject--
					}

				case "j":
					if m.selectedProject < len(m.projectList)-1 {
						m.selectedProject++
					}
        }
    }
    return m, nil
}

func (m model) View() string {
  s := fmt.Sprintf("%v's Showcase\n\n", m.userName)
  s += fmt.Sprintf("%v\n\n", m.introText)

	for i, project := range m.projectList {
		if m.selectedProject == i {
			s += "> "
		} else {
			s += "- "
		}
		s += fmt.Sprintf("%v\n", project.Name)
	}

  s += "\npress 'q' to quit."

  return s 

}

func open(url string) error {
	var cmd string
	var args []string

		switch runtime.GOOS{
		case "windows":
			cmd = "cmd"
			args = []string{"/c", "start"}
		case "darwin":
			cmd = "open"
		default:
			cmd = "xdg-open"
		}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

