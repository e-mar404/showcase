package main

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/e-mar404/showcase/internal/config"
	"github.com/e-mar404/showcase/internal/pages"
)

const (
	host = "localhost"
	port = "42069"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)

	if err != nil {
		log.Error("Could not create server")
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}	
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
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

	return pages.InitialModel(cfg), []tea.ProgramOption{tea.WithAltScreen()}
}

