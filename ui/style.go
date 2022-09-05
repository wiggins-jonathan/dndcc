package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	magenta = lipgloss.Color("#ff3399")
	blue    = lipgloss.Color("#0099ff")
	green   = lipgloss.Color("#50ff4a")

	lg                = lipgloss.NewStyle()
	titleStyle        = lg.Foreground(green)
	itemStyle         = lg.PaddingLeft(4)
	selectedItemStyle = lg.MarginLeft(4).Background(magenta).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	detailsStyle      = lg.PaddingLeft(4).Width(85)
	detailName        = lg.Foreground(blue)
	listStyle         = lg.PaddingLeft(4)

	footerUnselected = lg.PaddingLeft(4).Foreground(lipgloss.Color("#5c5c5c"))
	footerSelected   = lg.PaddingLeft(4).Foreground(green)
)
