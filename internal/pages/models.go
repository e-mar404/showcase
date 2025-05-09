package pages

import "github.com/e-mar404/showcase/internal/config"

type HomePage struct {
  userName string
  introText string
  projectList []config.Project
	selectedProject int
}
