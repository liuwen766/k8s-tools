package main

import (
	"fmt"
	"sync"
)

// Logger 日志对象是个单例实例
type Logger struct {
	name string
	log  []string
	lock sync.Mutex
}

var instance *Logger
var once sync.Once

// GetLogger 获取日志对象
func GetLogger(name string) *Logger {
	once.Do(func() {
		instance = &Logger{name: name}
	})
	return instance
}

func (l *Logger) Log(message string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.log = append(l.log, fmt.Sprintf("%s: %s", l.name, message))
}

func (l *Logger) PrintLog() {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, log := range l.log {
		fmt.Println(log)
	}
}
