package tui

import (
	"github.com/Bilou4/godo/configuration"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type TuiStyles struct {
	FocusedStyleBorder lipgloss.Style
	StatusStyle        lipgloss.Style
	HelpStyle          help.Styles
	StatusBarStyle     lipgloss.Style
	StatusText         lipgloss.Style
	ColumnStyle        lipgloss.Style

	ItemsStyle list.DefaultDelegate
	ListStyles list.Styles
}

func newTuiStyles(tuiCfg configuration.TuiConfig) *TuiStyles {
	if tuiCfg.FocusedStyleColor == "" {
		tuiCfg.FocusedStyleColor = "#a60598"
	}
	if tuiCfg.StatusForeground == "" {
		tuiCfg.StatusForeground = "#FFFDF5"
	}
	if tuiCfg.StatusBackground == "" {
		tuiCfg.StatusBackground = "#FF5F87"
	}
	if tuiCfg.StatusBarForegroundLight == "" {
		tuiCfg.StatusBarForegroundLight = "#343433"
	}
	if tuiCfg.StatusBarForegroundDark == "" {
		tuiCfg.StatusBarForegroundDark = "#C1C6B2"
	}
	if tuiCfg.StatusBarBackgroundLight == "" {
		tuiCfg.StatusBarBackgroundLight = "#D9DCCF"
	}
	if tuiCfg.StatusBarBackgroundDark == "" {
		tuiCfg.StatusBarBackgroundDark = "#353533"
	}
	statusBarStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.StatusBarForegroundLight, Dark: tuiCfg.StatusBarForegroundDark}).
		Background(lipgloss.AdaptiveColor{Light: tuiCfg.StatusBarBackgroundLight, Dark: tuiCfg.StatusBarBackgroundDark})

	if tuiCfg.NormalDescDark == "" {
		tuiCfg.NormalDescDark = "#777777"
	}
	if tuiCfg.NormalDescLight == "" {
		tuiCfg.NormalDescLight = "#A49FA5"
	}
	if tuiCfg.NormalTitleDark == "" {
		tuiCfg.NormalTitleDark = "#dddddd"
	}
	if tuiCfg.NormalTitleLight == "" {
		tuiCfg.NormalTitleLight = "#1a1a1a"
	}

	if tuiCfg.SelectedTitleBorderForegroundDark == "" {
		tuiCfg.SelectedTitleBorderForegroundDark = "#AD58B4"
	}

	if tuiCfg.SelectedTitleBorderForegroundLight == "" {
		tuiCfg.SelectedTitleBorderForegroundLight = "#F793FF"
	}
	if tuiCfg.SelectedTitleForegroundDark == "" {
		tuiCfg.SelectedTitleForegroundDark = "#EE6FF8"
	}
	if tuiCfg.SelectedTitleForegroundLight == "" {
		tuiCfg.SelectedTitleForegroundLight = "#EE6FF8"
	}

	if tuiCfg.SelectedDescForegroundLight == "" {
		tuiCfg.SelectedDescForegroundLight = "#F793FF"
	}
	if tuiCfg.SelectedDescForegroundDark == "" {
		tuiCfg.SelectedDescForegroundDark = "#AD58B4"
	}

	if tuiCfg.TitleBackgroundColor == "" {
		tuiCfg.TitleBackgroundColor = "62"
	}

	if tuiCfg.TitleForegroundColor == "" {
		tuiCfg.TitleForegroundColor = "230"
	}

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.NormalTitleLight, Dark: tuiCfg.NormalTitleDark}).
		Padding(0, 0, 0, 2)

	selectedStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedTitleBorderForegroundLight, Dark: tuiCfg.SelectedTitleBorderForegroundDark}).
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedTitleForegroundLight, Dark: tuiCfg.SelectedTitleForegroundDark}).
		Padding(0, 0, 0, 1)

	defaultDelegate := list.NewDefaultDelegate()

	defaultDelegate.Styles = list.DefaultItemStyles{
		NormalTitle: normalStyle,

		NormalDesc: normalStyle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.NormalDescLight, Dark: tuiCfg.NormalDescDark}),

		SelectedTitle: selectedStyle,

		SelectedDesc: selectedStyle.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedDescForegroundLight, Dark: tuiCfg.SelectedDescForegroundDark}),
	}

	listStyles := list.DefaultStyles()
	listStyles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color(tuiCfg.TitleBackgroundColor)).
		Foreground(lipgloss.Color(tuiCfg.TitleForegroundColor)).
		Padding(0, 1)

	return &TuiStyles{
		FocusedStyleBorder: lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(tuiCfg.FocusedStyleColor)),

		StatusStyle: lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color(tuiCfg.StatusForeground)).
			Background(lipgloss.Color(tuiCfg.StatusBackground)).
			Padding(0, 1).
			MarginRight(1),

		StatusBarStyle: statusBarStyle,
		StatusText:     lipgloss.NewStyle().Inherit(statusBarStyle),
		ColumnStyle: lipgloss.NewStyle().Padding(1, 2).
			Border(lipgloss.HiddenBorder()),
		ItemsStyle: defaultDelegate,
		ListStyles: listStyles,
	}
}

// Status Bar variables
const (
	statusKey = "STATUS"
)
