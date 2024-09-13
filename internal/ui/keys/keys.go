package keys

import "github.com/charmbracelet/bubbles/key"

var Select = key.NewBinding(
	key.WithKeys("enter"),
)

var Deselect = key.NewBinding(
	key.WithKeys("esc"),
)
