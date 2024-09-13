package types

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var _ Handlable = &BoolItem{}

type BoolItem struct {
	Item
	Value bool

	input *huh.Confirm
}

// Validate implements Handlable.
func (s *BoolItem) Validate() error {
	return s.input.Error()
}

func (s *BoolItem) Key() string {
	return s.Item.Key
}

// ViewInput implements Handlable.
func (s *BoolItem) ViewInput() string {
	return s.input.View()
}

// BecameActive implements Handlable.
func (s *BoolItem) BecameActive() tea.Cmd {
	s.input.Focus()

	return nil
}

// BecameInactive implements Handlable.
func (s *BoolItem) BecameInactive() tea.Cmd {
	s.input.Blur()

	return nil
}

// Update implements Handlable.
func (s *BoolItem) Update(msg tea.Msg) tea.Cmd {
	_, cmd := s.input.Update(msg)
	return cmd
}

func NewBool(i Item, value bool) *BoolItem {
	si := BoolItem{
		Item:  i,
		Value: value,
	}

	inp := huh.NewConfirm().
		Inline(true).
		Affirmative("True").
		Negative("False").
		Value(&si.Value)

	inp.WithKeyMap(huh.NewDefaultKeyMap()) // this is bullshit. todo report to huh

	si.input = inp

	return &si
}
