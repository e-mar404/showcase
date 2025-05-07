package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type config struct {
	UserName string `json:"user_name"`
	IntroText string `json:"intro_text"`
	ProjectList []string `json:"project_list"`
}

type homeModel struct {
	userName string
	introText string
	projectList []string
}

func main() {
	file, err := os.Open(".showcase.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	var config config
	if err = decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}

func homeModelFromConfig(cfg config) homeModel {
	return homeModel {
		userName: cfg.UserName,
		introText: cfg.IntroText,
		projectList: cfg.ProjectList,
	}
}

func (hm homeModel) Init() tea.Cmd {
	return nil
}

func (hm homeModel) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (hm homeModel) View() string {
	return ""
}

