package ui

import (
	"github.com/Kochanac/chonf/internal/ui/types"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type ConfigList struct {
	list list.Model
}

func (m ConfigList) Init() tea.Cmd {
	return nil
}

func (m ConfigList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ConfigList) View() string {
	return docStyle.Render(m.list.View())
}

func New() ConfigList {
	items := []list.Item{
		types.NewString(types.Item{Key: "Raspberry Pi’s", Desc: "I have ’em all over my house"}, ""),
		types.NewString(types.Item{Key: "Nutella", Desc: "It's good on toast"}, ""),
		types.NewString(types.Item{Key: "Bitter melon", Desc: "It cools you down"}, ""),
		types.NewString(types.Item{Key: "Nice socks", Desc: "And by that I mean socks without holes"}, "ffdsa"),
		types.NewInt(types.Item{Key: "Kek", Desc: "Only int here"}, "123"),
		types.NewBool(types.Item{Key: "Kek", Desc: "Only int here"}, false),
		types.NewBool(types.Item{Key: "Kek", Desc: "Only int here"}, true),
	}

	m := ConfigList{
		list: list.New(items, NewListDelegate(), 0, 0),
	}
	m.list.KeyMap.Quit = key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	)

	m.list.Title = "\"Трахер\" parasite config"
	return m
}
