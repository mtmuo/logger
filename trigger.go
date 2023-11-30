package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type hooks struct {
	hooks map[string]*hook
	mutex sync.RWMutex
}

func (hooks *hooks) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (hooks *hooks) hook(name string) *hook {
	hooks.mutex.RLock()
	defer hooks.mutex.RUnlock()
	if h := hooks.hooks[name]; h != nil {
		return h
	}
	return hooks.hooks["global"]
}

func (hooks *hooks) Fire(entry *logrus.Entry) error {
	if entry.Level < logrus.ErrorLevel {
		// 处理超级异常项目
	}
	module, ok := entry.Data["module"].(string)
	if !ok {
		return nil
	}

	if text, err := entry.String(); err == nil {
		hooks.hook(module).Trigger(text)
	}
	return nil
}

func (hooks *hooks) setTriggerNum(moduleName string, triggerNum uint64) {
	hooks.mutex.RLock()
	defer hooks.mutex.RUnlock()
	if h := hooks.hooks[moduleName]; h != nil {
		h.ErrNum.Store(0)
		h.errors = make([]string, 0)
		h.triggerNum = triggerNum
	}
}

func (hooks *hooks) setRotationInterval(moduleName string, rotationInterval time.Duration) {
	hooks.mutex.RLock()
	defer hooks.mutex.RUnlock()
	if h := hooks.hooks[moduleName]; h != nil {
		h.rotationInterval = rotationInterval
	}
}

func (hooks *hooks) setNotifyInterval(moduleName string, notifyInterval time.Duration) {
	hooks.mutex.RLock()
	defer hooks.mutex.RUnlock()
	if h := hooks.hooks[moduleName]; h != nil {
		h.notifyInterval = notifyInterval
	}
}

func (hooks *hooks) addHook(name string) {
	hooks.mutex.Lock()
	defer hooks.mutex.Unlock()
	if h := hooks.hooks[name]; h != nil {
		return
	}
	hooks.hooks[name] = newHook(name)
}

type hook struct {
	moduleName       string
	modules          map[string]*hook
	ErrNum           atomic.Uint64
	triggerNum       uint64
	rotationInterval time.Duration
	notifyInterval   time.Duration
	lastClear        time.Time
	lastNotify       time.Time
	errors           []string
}

func (t *hook) Trigger(text string) {
	if t.lastClear.Before(time.Now().Add(-t.rotationInterval)) {
		t.errors = make([]string, 0)
		t.ErrNum.Swap(0)
		t.lastClear = time.Now()
	}

	t.errors = append(t.errors, text)
	t.ErrNum.Add(1)

	if t.ErrNum.Load() < t.triggerNum {
		return
	}

	t.Send(
		fmt.Sprintf("%s:trigger %d error", t.moduleName, t.triggerNum),
		fmt.Sprintf(`
moduleName:%s

triggerNum: %d rotationInterval:%s notifyInterval:%s start:%s

%s

`,
			t.moduleName,
			t.triggerNum,
			t.rotationInterval.String(),
			t.notifyInterval.String(),
			t.lastClear.Format("2006-01-02 15:04:05"),
			strings.Join(t.errors, "\n"),
		),
	)
	t.errors = make([]string, 0)
	t.ErrNum.Swap(0)
	t.lastClear = time.Now()
}

func (t *hook) Reset() {
	t.ErrNum.Store(0)
}

func (t *hook) Send(subject, context string) {
	if t.lastNotify.Before(time.Now().Add(-t.notifyInterval)) {
		t.lastNotify = time.Now()
		notice.Send(subject, context)
	}
}

func newHook(moduleName string) *hook {
	return &hook{
		moduleName:       moduleName,
		triggerNum:       TriggerNum,
		rotationInterval: RotationInterval,
		notifyInterval:   NotifyInterval,
	}
}
