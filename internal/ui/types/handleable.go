package types

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Handlable interface {
	Key() string
	BecameActive() tea.Cmd
	BecameInactive() tea.Cmd

	Validate() error

	Update(msg tea.Msg) tea.Cmd
	ViewInput() string

	list.DefaultItem
}
