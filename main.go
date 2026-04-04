package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/wish"
    "github.com/charmbracelet/wish/bubbletea"
    "github.com/charmbracelet/ssh"
)

const (
    host = "0.0.0.0"
    port = "22"
)

func main() {
    s, err := wish.NewServer(
        wish.WithAddress(host+":"+port),
        wish.WithHostKeyPath(".ssh/id_ed25519"),
        wish.WithMiddleware(
            bubbletea.Middleware(teaHandler),
        ),
    )
    if err != nil {
        log.Fatal("Could not create server:", err)
    }
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    log.Printf("SSH server listening on %s:%s", host, port)
    go func() {
        if err := s.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    <-done
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    s.Shutdown(ctx)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
    pty, _, _ := s.Pty()
    width := pty.Window.Width
    height := pty.Window.Height

    m := ui.NewModel(width, height)

    return m, []tea.ProgramOption{
        tea.WithAltScreen(),
    }
}