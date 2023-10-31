//go:build debug

package log

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func LogToFile() (*os.File, error) {
	return tea.LogToFile("godo-debug.log", "")
}
