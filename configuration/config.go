package configuration

type Config struct {
	DbPath string    `json:"path_to_the_database"`
	Tui    TuiConfig `json:"tui,omitempty"`
}

type TuiConfig struct {
	FocusedStyleColor        string `json:"focused_style_color,omitempty"`
	StatusForeground         string `json:"status_foreground,omitempty"`
	StatusBackground         string `json:"status_background,omitempty"`
	HelpForegroundColor      string `json:"help_foreground_color,omitempty"`
	StatusBarForegroundLight string `json:"status_bar_foreground_light,omitempty"`
	StatusBarForegroundDark  string `json:"status_bar_foreground_dark,omitempty"`

	StatusBarBackgroundLight string `json:"status_bar_background_light,omitempty"`
	StatusBarBackgroundDark  string `json:"status_bar_background_dark,omitempty"`

	NormalTitleLight string `json:"normal_title_light,omitempty"`
	NormalTitleDark  string `json:"normal_title_dark,omitempty"`

	NormalDescLight string `json:"normal_desc_light,omitempty"`
	NormalDescDark  string `json:"normal_desc_dark,omitempty"`

	SelectedTitleBorderForegroundLight string `json:"selected_title_border_foreground_light,omitempty"`
	SelectedTitleBorderForegroundDark  string `json:"selected_title_border_foreground_dark,omitempty"`
	SelectedTitleForegroundLight       string `json:"selected_title_foreground_light,omitempty"`
	SelectedTitleForegroundDark        string `json:"selected_title_foreground_dark,omitempty"`

	SelectedDescForegroundLight string `json:"selected_desc_foreground_light,omitempty"`
	SelectedDescForegroundDark  string `json:"selected_desc_foreground_dark,omitempty"`

	DimmedTitleForegroundLight string `json:"dimmed_title_foreground_light,omitempty"`
	DimmedTitleForegroundDark  string `json:"dimmed_title_foreground_dark,omitempty"`

	DimmedDescForegroundLight string `json:"dimmed_desc_foreground_light,omitempty"`
	DimmedDescForegroundDark  string `json:"dimmed_desc_foreground_dark,omitempty"`
}
