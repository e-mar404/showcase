package pages

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/e-mar404/showcase/internal/config"
)

type model struct {
  status status
  userName string
  introText string
  projectList []config.Project
	selectedProject int
  projectsLoaded int
  progress progress.Model
}

type loadedRepo struct {}
type finishedDownloadingRepos struct {}
type downloadRepos struct {}

type status int

const (
  LOADING status = iota
  FAILED
  LOADED
)

const (
  repoBaseDir = ".repos"
  padding = 2
  maxWidth = 80
)

func InitialModel(cfg config.Config) model{
	return model{
    status: LOADING,
    userName: cfg.UserName,
    introText: cfg.IntroText,
    projectList: cfg.ProjectList,
		selectedProject: 0,
    projectsLoaded: 0,
    progress: progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C")),
	}
}

func (m model) Init() tea.Cmd {
  if err := os.RemoveAll(repoBaseDir); err != nil {
    log.Errorf("Error cleaning up repo dir: %v", err)
    return tea.Quit
  }

  return func() tea.Msg { return downloadRepos{} }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  log.Info("running update")
  switch msg := msg.(type) {
  case downloadRepos:
    return m, downloadRepo(m)

  case loadedRepo:
    m.projectsLoaded++
    if m.projectsLoaded == len(m.projectList){
      log.Info("finished downloading repos")
      return m, func () tea.Msg { return finishedDownloadingRepos{} }
    }
    return m, downloadRepo(m)

  case finishedDownloadingRepos:
    m.status = LOADED

  case tea.KeyMsg:
    switch msg.String() {
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
  log.Info("running view")
  if m.status == LOADING {
    percentDone := float64(m.projectsLoaded)/float64(len(m.projectList))
    pad := strings.Repeat(" ", padding)
    return pad + "setting up projects\n\n" + pad + m.progress.ViewAs(percentDone)
  }

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

  s += "\n(press 'q' to quit)"

  return s 

}

func downloadRepo(m model) tea.Cmd {
  return func() tea.Msg {
    log.Info("running download repo")
    
    project := m.projectList[m.projectsLoaded]
    cmd := "git"
    repoDir := filepath.Join(repoBaseDir, project.Name)
    args := []string{"clone", project.Url, repoDir} 
    command := exec.Command(cmd, args...)
    _, err := command.Output()
    if err != nil {
      log.Errorf("Unable to get repo from url: %s, err: %v", project.Url, err)
      return tea.Quit
    }

    return loadedRepo{}
  }
}
