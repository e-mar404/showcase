package main 

import (
	"encoding/json"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/e-mar404/showcase/internal/config"
	"github.com/e-mar404/showcase/internal/pages"
)

func main() {
	file, err := os.Open(".showcase.json")
	if err != nil {
		log.Fatal(err)
    os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	var cfg config.Config 
	if err = decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
    os.Exit(1)
	}

  p := tea.NewProgram(pages.NewHomePage(cfg))
  if _, err = p.Run(); err != nil {
		log.Fatal(err)
    os.Exit(1)
  }
}

