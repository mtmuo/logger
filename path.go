package logger

import (
	"os"
	"path/filepath"
	"strings"
)

type Exec struct {
	Path    string
	AppName string
	Ext     string
}

func Executable() Exec {
	exePath, _ := os.Executable()

	return Exec{
		Path:    filepath.Dir(exePath),
		AppName: strings.TrimRight(filepath.Base(exePath), filepath.Ext(exePath)),
		Ext:     filepath.Ext(exePath),
	}
}
