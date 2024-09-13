package events

import "github.com/charmbracelet/bubbles/help"

// todo make custom help

type HelpUpdate struct {
	NewHelp help.Model
}
