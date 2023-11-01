# godo


`godo` is a todo-list app written in Go usable through command line and a highly customizable TUI.

## Installation

Using the `go` command line.
```bash
go install github.com/Bilou4/godo@latest
```

Or simply get the last release and use it.

## How to use it?

To start using it, you need to run the `init` command.

```bash
godo init
```

And then start by creating a new list.

```bash
godo addList --list-name groceries
```

After that, you can create/update/remove tasks and lists. Or simply list everything.
```bash
# list tasks and list
godo ls

# create a task in the 'groceries' list
godo add --name tomatoes --list-id 1
```

For an everyday use, run the Terminal User Interface.

```bash
godo tui
```

For more information, use the `godo help` command.

## Customization
Default colors can be overwritten in the configuration file (`~/.config/godo/config.json`).

```json
{
    "path_to_the_database": "<path-to-your-database>",
    "tui": {
        "status_bar": {
            "status_foreground": "#F2E4DC",
            "status_background": "#D99311",
            "status_bar_foreground_light": "#F2E4DC",
            "status_bar_foreground_dark": "#F2E4DC",
            "status_bar_background_light": "#254559",
            "status_bar_background_dark": "#254559"
        },
        "task": {
            "normal_title_light": "#F2E4DC",
            "normal_title_dark": "#F2E4DC",
            "normal_desc_light": "#F2E4DC",
            "normal_desc_dark": "#F2E4DC",
            "selected_title_border_foreground_light": "#8C2F1B",
            "selected_title_border_foreground_dark": "#8C2F1B",
            "selected_title_foreground_light": "#D99311",
            "selected_title_foreground_dark": "#D99311",
            "selected_desc_foreground_light": "#D99311",
            "selected_desc_foreground_dark": "#D99311"
        },
        "list": {
            "focused_style_color": "#254559",
            "title_foreground_color": "#F2E4DC",
            "title_background_color": "#8C2F1B"
        }
    }
}
```

## Made with

- [The bubbletea TUI framework](https://github.com/charmbracelet/bubbletea)
- [The cobra framework to create CLI](https://github.com/spf13/cobra)
- [GORM, the ORM library for Go](https://gorm.io)
- [Go-pretty, to pretty print in console](https://github.com/jedib0t/go-pretty)