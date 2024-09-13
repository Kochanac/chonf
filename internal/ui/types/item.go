package types

import "github.com/charmbracelet/bubbles/list"

var _ list.DefaultItem = Item{}

type Item struct {
	Key  string
	Desc string
}

// Description implements list.DefaultItem.
func (s Item) Description() string {
	return s.Desc
}

// FilterValue implements list.DefaultItem.
func (s Item) FilterValue() string {
	return s.Key
}

// Title implements list.DefaultItem.
func (s Item) Title() string {
	return s.Key
}
