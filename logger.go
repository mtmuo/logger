package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	TriggerNum       = uint64(50)
	RotationInterval = time.Minute * 30
	NotifyInterval   = time.Minute * 30
	notice           = NewNotice()
	std              = logrus.New()
	loggerHooks      = newHooks()
)

type Logger struct {
	module string
	*logrus.Entry
}

func (l *Logger) Reset() {

}

func (l *Logger) Recover(fun func(err error)) {
	if err := recover(); err != nil {
		if fun != nil {
			fun(fmt.Errorf("%v", err))
		} else {
			l.Error(err)
		}
	}
}

func (l *Logger) Send(subject, context string) {
	notice.Send(fmt.Sprintf("%s:%s", l.module, subject), context)
}

func (l *Logger) SetTriggerNum(triggerNum uint64) {
	loggerHooks.setTriggerNum(l.module, triggerNum)
}

func (l *Logger) SetRotationInterval(rotationInterval time.Duration) {
	loggerHooks.setRotationInterval(l.module, rotationInterval)
}

func (l *Logger) SetNotifyInterval(notifyInterval time.Duration) {
	loggerHooks.setRotationInterval(l.module, notifyInterval)
}

func WithModule(moduleName string) *Logger {
	// 判断是否存在监控统计
	loggerHooks.addHook(moduleName)
	return &Logger{
		module: moduleName,
		Entry:  std.WithField("module", moduleName),
	}
}
