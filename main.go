package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	peer "github.com/muka/peerjs-go"
)

/* ==[TYPES]== */

type appModel struct {
	viewId int

	client *peer.Peer
	peers  []peer.Connection
}

type TeaView func(m appModel) string

type viewModel struct {
	id   int
	key  string
	view TeaView
}

/* ==[CONSTANTS]== */

const (
	APP_SETTINGS     = iota
	CONTACTS         = iota
	CONTACT_SETTINGS = iota
	CHAT             = iota
	CALL             = iota
)

/* ==[FUNCS]== */

func mapViews(views []viewModel) map[int]viewModel {
	var mappedViews = map[int]viewModel{}
	for _, view := range views {
		mappedViews[view.id] = view
	}
	return mappedViews
}

/* ==[APP]== */

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		var key = msg.String()

		for _, view := range views {
			if key == view.key {
				m.viewId = view.id
				return m, nil
			}
		}

		switch key {

		case "ctrl+c", "q":
			return m, tea.Quit

		}
	}

	return m, nil
}

var views = mapViews([]viewModel{

	{
		id:  APP_SETTINGS,
		key: "s",
		view: func(m appModel) string {
			return "App Settings"
		},
	},

	{
		id:  CONTACTS,
		key: "c",
		view: func(m appModel) string {
			return "Contacts"
		},
	},
})

func (m appModel) View() string {
	return views[m.viewId].view(m)
}

func main() {
	p := tea.NewProgram(appModel{
		viewId: CONTACTS,

		client: nil,
		peers:  make([]peer.Connection, 1),
	})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
