package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// HTTP Request TUI.

type model struct {
	status  int    // HTTP response
	err     error  // Possible error
	url     string // URL string
	spinner spinner.Model
}

// checkServer is a `Cmd` that makes a request to a server
// and returns the result as a `Msg`
//
// `Cmd`s are functions that perform some I/O and then return a `Msg`.
// Checking the time, ticking a timer, reading from a disk, etc
// are all I/O and should be run through commands.
func checkServer(url string) tea.Cmd {
	return func() tea.Msg {
		// Create an HTTP client and make a GET request.
		c := &http.Client{Timeout: 10 * time.Second}
		res, err := c.Get(url)

		if err != nil {
			// There was an error making our request. Wrap the error we received
			// in a message and return it.
			return errMsg{err}
		}
		// We received a response from the server. Return the HTTP status code
		// as a message.
		return statusMsg(res.StatusCode)
	}
}

type statusMsg int

type errMsg struct{ err error }

// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }

// We don't call the function
// the Bubble Tea runtime will do that when the time is right.
func (m *model) Init() tea.Cmd {
	m.spinner = spinner.New()
	m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	m.spinner.Spinner = spinner.Line

	return tea.Batch(
		checkServer(m.url),
		m.spinner.Tick,
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		// The server returned a status message. Save it to our model. Also
		// tell the Bubble Tea runtime we want to exit because we have nothing
		// else to do. We'll still be able to render a final view with our
		// status message.
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		// There was an error. Note it in the model. And tell the runtime
		// we're done and want to quit.
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		// Ctrl+c exits. Even with short running program it's good to have
		// a quit key, just in case your logic is off. Users will be very
		// annoyed if they can't exit.
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
		if msg.Type == tea.KeyEsc {
			return &model{
				err: errors.New("escape pressed"),
			}, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	default:
		return m, nil
	}

	// If we happen to get any other messages, don't do anything.
	return m, nil
}

// View is very straightforward. We look at the current model
// and build a string accordingly
func (m *model) View() string {
	// If there's an error, print it out and don't do anything else.
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v\n", m.err)
	}

	// Tell the user we're doing something.
	s := fmt.Sprintf("%s Checking %s ... ", m.spinner.View(), m.url)

	// When the server responds with a status, add it to the current line.
	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	// Send off whatever we came up with above for rendering.
	return "\n" + s + "\n\n"
}

func initialModel() *model {
	return &model{
		url: "https://charm.sh/",
	}
}

func main() {
	if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
