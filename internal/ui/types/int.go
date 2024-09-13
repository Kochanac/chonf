package types

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var _ Handlable = &IntItem{}

type IntItem struct {
	Item
	Value string

	input *huh.Input
}

// Validate implements Handlable.
func (s *IntItem) Validate() error {
	return s.input.Error()
}

func (s *IntItem) Key() string {
	return s.Item.Key
}

// ViewInput implements Handlable.
func (s *IntItem) ViewInput() string {
	return s.input.View()
}

// BecameActive implements Handlable.
func (s *IntItem) BecameActive() tea.Cmd {
	s.input.Focus()
	return nil
}

// BecameInactive implements Handlable.
func (s *IntItem) BecameInactive() tea.Cmd {
	s.input.Blur()
	return nil
}

// Update implements Handlable.
func (s *IntItem) Update(msg tea.Msg) tea.Cmd {
	_, cmd := s.input.Update(msg)
	return cmd
}

func NewInt(i Item, value string) *IntItem {
	si := IntItem{
		Item:  i,
		Value: value,
	}

	inp := huh.NewInput().
		Prompt("#").
		Validate(func(s string) error {
			_, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			return nil
		}).
		Value(&si.Value)

	si.input = inp

	return &si
}
