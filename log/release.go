//go:build !debug

package log

import (
	"os"
)

func LogToFile() (*os.File, error) {
	return nil, nil
}
