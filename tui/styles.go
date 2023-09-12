package tui

import (
	"github.com/Bilou4/godo/configuration"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type TuiStyles struct {
	FocusedStyleColor lipgloss.Style
	StatusStyle       lipgloss.Style
	HelpStyle         lipgloss.Style
	StatusBarStyle    lipgloss.Style
	StatusText        lipgloss.Style
	ColumnStyle       lipgloss.Style

	// items styles
	list.DefaultDelegate
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
	if tuiCfg.HelpForegroundColor == "" {
		tuiCfg.HelpForegroundColor = "241"
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

	if tuiCfg.DimmedTitleForegroundLight == "" {
		tuiCfg.DimmedTitleForegroundLight = "#A49FA5"
	}
	if tuiCfg.DimmedTitleForegroundDark == "" {
		tuiCfg.DimmedTitleForegroundDark = "#777777"
	}
	if tuiCfg.DimmedDescForegroundLight == "" {
		tuiCfg.DimmedDescForegroundLight = "#C2B8C2"
	}
	if tuiCfg.DimmedDescForegroundDark == "" {
		tuiCfg.DimmedDescForegroundDark = "#4D4D4D"
	}
	defaultDelegate := list.NewDefaultDelegate()
	nt := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.NormalTitleLight, Dark: tuiCfg.NormalTitleDark}).
		Padding(0, 0, 0, 2)
	st := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedTitleBorderForegroundLight, Dark: tuiCfg.SelectedTitleBorderForegroundDark}).
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedTitleForegroundLight, Dark: tuiCfg.SelectedTitleForegroundDark}).
		Padding(0, 0, 0, 1)
	dt := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.DimmedTitleForegroundLight, Dark: tuiCfg.DimmedTitleForegroundDark}).
		Padding(0, 0, 0, 2)
	defaultDelegate.Styles = list.DefaultItemStyles{
		NormalTitle: nt,

		NormalDesc: nt.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.NormalDescLight, Dark: tuiCfg.NormalDescDark}),

		SelectedTitle: st,

		SelectedDesc: st.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.SelectedDescForegroundLight, Dark: tuiCfg.SelectedDescForegroundDark}),

		DimmedTitle: dt,

		DimmedDesc: dt.Copy().
			Foreground(lipgloss.AdaptiveColor{Light: tuiCfg.DimmedDescForegroundLight, Dark: tuiCfg.DimmedDescForegroundDark}),

		FilterMatch: lipgloss.NewStyle().Underline(true),
	}

	return &TuiStyles{
		FocusedStyleColor: lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(tuiCfg.FocusedStyleColor)),
		StatusStyle: lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color(tuiCfg.StatusForeground)).
			Background(lipgloss.Color(tuiCfg.StatusBackground)).
			Padding(0, 1).
			MarginRight(1),
		HelpStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(tuiCfg.HelpForegroundColor)),

		StatusBarStyle: statusBarStyle,
		StatusText:     lipgloss.NewStyle().Inherit(statusBarStyle),
		ColumnStyle: lipgloss.NewStyle().Padding(1, 2).
			Border(lipgloss.HiddenBorder()),
		DefaultDelegate: defaultDelegate,
	}
}

// Status Bar variables
const (
	statusKey = "STATUS"
)
