package pages

import (
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/e-mar404/showcase/internal/config"
)

const (
	repoBaseDir = ".repos"
)

type ProjectPage struct {
	state state 
	session ssh.Session
	repoDir string
	repoURL string
	repoName string
}

type state int

const (
	LOADING state = iota
	FAILED
	LOADED	
)

func NewProjectPage(project config.Project, s ssh.Session) ProjectPage {
	if err := os.RemoveAll(repoBaseDir); err != nil {
		log.Errorf("Error cleaning up repo dir: %v", err)
		return ProjectPage{
			state: FAILED,
			session: s,
		}
	}

	repoDir := filepath.Join(repoBaseDir, project.Name)
	cmd := "git"
	args := []string{"clone", project.Url, repoDir}
	command := exec.Command(cmd, args...)
	log.Info(command)	
	if err := command.Start(); err != nil {
		log.Errorf("Unable to get repo from url: %s, err: %v", project.Url, err)
		return ProjectPage{
			state: FAILED,
		}
	}
	
	return ProjectPage{
		state: LOADED,
		repoDir: repoDir,
		repoURL: project.Url,
		repoName: project.Name,
	}
}

func (lp ProjectPage) Init() tea.Cmd {
	return nil
}

func (lp ProjectPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return lp, tea.Quit
			}
	}
	return lp, nil
}

func (lp ProjectPage) View() string {
	if lp.state == FAILED{
		return "Unable to load repo"
	}

	return "Help menu"
}
