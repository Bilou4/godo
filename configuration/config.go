package configuration

type Config struct {
	DbPath string    `json:"path_to_the_database"`
	Tui    TuiConfig `json:"tui,omitempty"`
}

type TuiConfig struct {
	StatusBarConfig `json:"status_bar,omitempty"`
	TaskConfig      `json:"task,omitempty"`
	ListConfig      `json:"list,omitempty"`
}

type StatusBarConfig struct {
	StatusForeground         string `json:"status_foreground,omitempty"`
	StatusBackground         string `json:"status_background,omitempty"`
	StatusBarForegroundLight string `json:"status_bar_foreground_light,omitempty"`
	StatusBarForegroundDark  string `json:"status_bar_foreground_dark,omitempty"`

	StatusBarBackgroundLight string `json:"status_bar_background_light,omitempty"`
	StatusBarBackgroundDark  string `json:"status_bar_background_dark,omitempty"`
}

type ListConfig struct {
	FocusedStyleColor string `json:"focused_style_color,omitempty"`

	TitleForegroundColor string `json:"title_foreground_color,omitempty"`
	TitleBackgroundColor string `json:"title_background_color,omitempty"`
}

type TaskConfig struct {
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
}
