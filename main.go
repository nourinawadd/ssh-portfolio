package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/muesli/termenv"

	"github.com/nourinawadd/ssh-portfolio/ui"
)

func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)

	s, err := wish.NewServer(
		wish.WithAddress(":2323"),
		wish.WithHostKeyPath("/home/ubuntu/.ssh/id_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, _ := s.Pty()

				width := 80
				height := 24

				if pty.Window.Width > 0 {
					width = pty.Window.Width
				}
				if pty.Window.Height > 0 {
					height = pty.Window.Height
				}

				m := ui.NewModel(width, height)
				return m, []tea.ProgramOption{tea.WithAltScreen()}
			}),
		),
	)
	if err != nil {
		log.Fatal("could not create server:", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Println("SSH portfolio listening on :2323")
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("server error:", err)
		}
	}()

	<-done
	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("shutdown error:", err)
	}
}