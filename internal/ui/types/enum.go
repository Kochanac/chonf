package types

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var _ Handlable = &EnumItem{}

type EnumItem struct {
	Item
	PossibleValues []string
	Value          string

	input    *huh.Select[string]
	isActive bool
}

// Validate implements Handlable.
func (s *EnumItem) Validate() error {
	return s.input.Error()
}

func (s *EnumItem) Key() string {
	return s.Item.Key
}

// ViewInput implements Handlable.
func (s *EnumItem) ViewInput() string {
	return s.input.View()
}

// BecameActive implements Handlable.
func (s *EnumItem) BecameActive() tea.Cmd {
	s.input.Focus()

	return nil
}

// BecameInactive implements Handlable.
func (s *EnumItem) BecameInactive() tea.Cmd {
	s.input.Blur()

	return nil
}

// Update implements Handlable.
func (s *EnumItem) Update(msg tea.Msg) tea.Cmd {
	_, cmd := s.input.Update(msg)
	return cmd
}

func NewEnum(i Item, possibleValues []string, value string) *EnumItem {
	si := EnumItem{
		Item:           i,
		PossibleValues: possibleValues,
		Value:          value,
	}

	inp := huh.NewSelect[string]()

	si.input = inp

	return &si
}
