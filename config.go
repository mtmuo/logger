package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Config struct {
	Stdout           bool      `json:"stdout" yaml:"stdout"`
	FileName         string    `json:"fileName" yaml:"fileName"`
	TimeFormat       string    `json:"timeFormat" yaml:"timeFormat"`
	Path             string    `json:"path" yaml:"path"`
	Level            string    `json:"level" yaml:"level"`
	MaxFile          uint      `json:"maxFile" yaml:"maxFile"`
	TriggerNum       uint64    `json:"triggerNum" yaml:"triggerNum"`
	RotationInterval int64     `json:"rotationInterval" yaml:"rotationInterval"`
	NotifyInterval   uint      `json:"notifyInterval" yaml:"notifyInterval"`
	Email            *Email    `json:"email" yaml:"email"`
	WxPusher         *WxPusher `json:"wxPusher" yaml:"wxPusher"`
}

func (c Config) fileName() string {
	if c.FileName == "" {
		c.FileName = "%Y%m%d.log"
	}
	//
	_ = os.MkdirAll(c.path(), os.ModePerm)
	return filepath.Join(c.path(), "%Y%m%d.log")
}

func (c Config) path() string {
	if c.Path == "" {
		return "./logs/"
	}
	if filepath.IsAbs(c.Path) {
		return c.Path
	}
	executable := Executable()
	return filepath.Join(executable.Path, c.Path)
}

func (c Config) Formatter() *logrus.TextFormatter {
	if c.TimeFormat == "" {
		c.TimeFormat = "2006-01-02 15:04:05"
	}
	return &logrus.TextFormatter{
		TimestampFormat: c.TimeFormat,
	}
}

func (c Config) ParseLevel() logrus.Level {
	level, err := logrus.ParseLevel(c.Level)
	if err != nil {
		return logrus.InfoLevel
	}
	return level
}

func (c Config) maxFile() uint {
	if c.MaxFile == 0 {
		return 7
	}
	return c.MaxFile
}
