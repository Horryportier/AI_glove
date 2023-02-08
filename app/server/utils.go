package server

import (
	"github.com/charmbracelet/lipgloss"
)

var (
        red = lipgloss.Color("#FF0034")
        green = lipgloss.Color("#51FF00")
        ErrorStyle = lipgloss.NewStyle().Foreground(red)
        GoodStyle = lipgloss.NewStyle().Foreground(green)
)
