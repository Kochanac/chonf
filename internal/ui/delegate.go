package ui

import (
	"io"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Kochanac/chonf/internal/ui/keys"
	"github.com/Kochanac/chonf/internal/ui/types"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

var _ list.ItemDelegate = &ListDelegate{}

type ListDelegate struct {
	activeItem      types.Handlable
	defaultDelegate list.DefaultDelegate
}

func NewListDelegate() *ListDelegate {
	return &ListDelegate{
		activeItem:      nil,
		defaultDelegate: list.NewDefaultDelegate(),
	}
}

// Height implements list.ItemDelegate.
func (s *ListDelegate) Height() int {
	return 3
}

// Render implements list.ItemDelegate.
func (s *ListDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	it, ok := item.(types.Handlable)
	if !ok {
		return
	}

	s.defaultDelegate.Render(w, m, index, it)
	w.Write([]byte("\n"))
	w.Write([]byte(it.ViewInput()))
}

// Spacing implements list.ItemDelegate.
func (s *ListDelegate) Spacing() int {
	return s.defaultDelegate.Spacing()
}

// Update implements list.ItemDelegate.
func (s *ListDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	defer s.defaultDelegate.Update(msg, m)

	si, ok := m.SelectedItem().(types.Handlable)
	if !ok {
		return nil
	}

	var cmd tea.Cmd

	if s.activeItem != nil && si.Key() != s.activeItem.Key() {
		if s.activeItem != nil {
			cmd = tea.Batch(cmd, s.activeItem.BecameInactive())
		}
		s.activeItem = nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, keys.Select) {
			cmd = tea.Batch(cmd,
				s.Deselect(m),
				s.Select(si, m),
			)
		}
		if key.Matches(msg, keys.Deselect) && s.activeItem != nil {
			cmd = tea.Batch(cmd, s.Deselect(m))
		}
	default:
		return nil
	}

	if s.activeItem != nil {
		cmd = tea.Batch(cmd, s.activeItem.Update(msg))
	}

	return cmd
}

func (s *ListDelegate) Deselect(m *list.Model) tea.Cmd {
	if s.activeItem == nil {
		return nil
	}

	var cmd tea.Cmd

	cmd = tea.Batch(cmd, s.activeItem.BecameInactive())
	if s.activeItem.Validate() != nil {
		cmd = tea.Batch(cmd, s.activeItem.BecameActive())
		return cmd
	}

	s.activeItem = nil
	enableKeymaps(m)

	return cmd
}

func (s *ListDelegate) Select(newItem types.Handlable, m *list.Model) tea.Cmd {
	s.activeItem = newItem
	cmd := newItem.BecameActive()

	disableKeymaps(m)

	return cmd
}

func disableKeymaps(m *list.Model) {
	m.KeyMap.CursorUp.SetEnabled(false)
	m.KeyMap.CursorDown.SetEnabled(false)
	m.KeyMap.NextPage.SetEnabled(false)
	m.KeyMap.PrevPage.SetEnabled(false)
	m.KeyMap.GoToStart.SetEnabled(false)
	m.KeyMap.GoToEnd.SetEnabled(false)
	m.KeyMap.Filter.SetEnabled(false)
	m.KeyMap.ClearFilter.SetEnabled(false)
	m.KeyMap.CancelWhileFiltering.SetEnabled(false)
	m.KeyMap.AcceptWhileFiltering.SetEnabled(false)
	m.KeyMap.Quit.SetEnabled(false)
	m.KeyMap.ShowFullHelp.SetEnabled(false)
	m.KeyMap.CloseFullHelp.SetEnabled(false)
}

func enableKeymaps(m *list.Model) {
	m.KeyMap.CursorUp.SetEnabled(true)
	m.KeyMap.CursorDown.SetEnabled(true)
	m.KeyMap.NextPage.SetEnabled(true)
	m.KeyMap.PrevPage.SetEnabled(true)
	m.KeyMap.GoToStart.SetEnabled(true)
	m.KeyMap.GoToEnd.SetEnabled(true)
	m.KeyMap.Filter.SetEnabled(true)
	m.KeyMap.ClearFilter.SetEnabled(true)
	m.KeyMap.CancelWhileFiltering.SetEnabled(true)
	m.KeyMap.AcceptWhileFiltering.SetEnabled(false)
	m.KeyMap.Quit.SetEnabled(true)
	m.KeyMap.ShowFullHelp.SetEnabled(true)
	m.KeyMap.CloseFullHelp.SetEnabled(true)
}
