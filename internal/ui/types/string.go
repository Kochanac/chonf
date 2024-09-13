package types

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var _ Handlable = &StringItem{}

type StringItem struct {
	Item
	Value string

	input *huh.Input
}

// Validate implements Handlable.
func (s *StringItem) Validate() error {
	return s.input.Error()
}

func (s *StringItem) Key() string {
	return s.Item.Key
}

// ViewInput implements Handlable.
func (s *StringItem) ViewInput() string {
	return s.input.View()
}

// BecameActive implements Handlable.
func (s *StringItem) BecameActive() tea.Cmd {
	slog.Debug("became active", "key", s.Key)
	s.input.Focus()

	return nil
}

// BecameInactive implements Handlable.
func (s *StringItem) BecameInactive() tea.Cmd {
	slog.Debug("became inactive", "key", s.Key)
	s.input.Blur()

	return nil
}

// Update implements Handlable.
func (s *StringItem) Update(msg tea.Msg) tea.Cmd {
	_, cmd := s.input.Update(msg)
	return cmd
}

func NewString(i Item, value string) *StringItem {
	si := StringItem{
		Item:  i,
		Value: value,
	}

	inp := huh.NewInput().
		Prompt(">").
		Value(&si.Value)

	si.input = inp

	return &si
}
