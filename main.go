package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mhaatha/gua-cli/internal/config"
	appError "github.com/mhaatha/gua-cli/internal/errors"
	"github.com/mhaatha/gua-cli/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error when calling LoadConfig: %v\n", appError.ErrCannotLoadEnv)
		os.Exit(1)
	}

	fetchDataService := service.NewFetchDataService(cfg)

	p := tea.NewProgram(initialModel(fetchDataService))
	if _, err := p.Run(); err != nil {
		log.Printf("error when running the program: %v\n", err)
		os.Exit(1)
	}
}

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	service   service.FetchDataService
	err       error
}

func initialModel(fetchDataService service.FetchDataService) model {
	ti := textinput.New()
	ti.Placeholder = "username"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	return model{
		textInput: ti,
		service:   fetchDataService,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			username := m.textInput.Value()

			return m, checkUserCmd(m.service, username)
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Ups, error is occurred: %v", m.err)
	}

	return fmt.Sprintf(
		"Which GitHub user do you want to check?\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func checkUserCmd(service service.FetchDataService, username string) tea.Cmd {
	return func() tea.Msg {
		err := service.GetUsername(username)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}
