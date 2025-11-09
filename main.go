package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Shopping list TUI.

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do list items are selected
}

// initialModel return our initial model
// however, we could just as easily define the initial model
// as a variabel elsewhere, too.
func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	// return nil means "no I/O right now, please."
	return nil
}

// Will be called when "things happen."
// Its job is to look at what has happened
// and return an updated model in response.
//
// It can also return a `Cmd` to make more things happen.
//
// In this case, when a user pressed the down arrow
// Update's job is to notice that the down arrow was pressed
// and move the cursor accordingly (or not).
//
// The "something happened" comes in the form of a `Msg`, which can be any type.
// Messages are the result of some I/O that took place
// such as a keypress, timer tick, or a response from a server.
//
// We usually figure out which type of `Msg` we received with a type switch
// but you could also use a type assertion.
//
// For now we'll just deal with tea.KeyMsg messages, which are automatically
// sent to the update function when keys are pressed.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actually key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// The selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

// View is the simplest method.
//
// We look at the model in its current state and use it to return a `string`.
// That string is our UI!
//
// Since the view describes the entire UI of the application
// you don't have to worry about redrawing logic and stuff like that.
// Bubble Tea takes care of it for you.
func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q or esc to quit.\n"

	return s
}
