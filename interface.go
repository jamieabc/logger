// Copyright (c) 2014-2018 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Generated file: Do _NOT_ Modify
// Generated on: 2017-09-19T11:23:21+08:00

package logger

import (
	"fmt"
	"github.com/bitmark-inc/logger/level"
)

func (l *L) Trace(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.TraceValue {
		l.log.Trace(l.formatPrefix + message)
	}
}

func (l *L) Tracef(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.TraceValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Trace(s)
	}
}

func (l *L) Tracec(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.TraceValue {
		l.log.Trace(l.formatPrefix + closure())
	}
}

func (l *L) Debug(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.DebugValue {
		l.log.Debug(l.formatPrefix + message)
	}
}

func (l *L) Debugf(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.DebugValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Debug(s)
	}
}

func (l *L) Debugc(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.DebugValue {
		l.log.Debug(l.formatPrefix + closure())
	}
}

func (l *L) Info(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.InfoValue {
		l.log.Info(l.formatPrefix + message)
	}
}

func (l *L) Infof(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.InfoValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Info(s)
	}
}

func (l *L) Infoc(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.InfoValue {
		l.log.Info(l.formatPrefix + closure())
	}
}

func (l *L) Warn(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.WarnValue {
		_ = l.log.Warn(l.formatPrefix + message)
	}
}

func (l *L) Warnf(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.WarnValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Warn(s)
	}
}

func (l *L) Warnc(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.WarnValue {
		_ = l.log.Warn(l.formatPrefix + closure())
	}
}

func (l *L) Error(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.ErrorValue {
		_ = l.log.Error(l.formatPrefix + message)
	}
}

func (l *L) Errorf(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.ErrorValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Error(s)
	}
}

func (l *L) Errorc(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.ErrorValue {
		_ = l.log.Error(l.formatPrefix + closure())
	}
}

func (l *L) Critical(message string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.CriticalValue {
		_ = l.log.Critical(l.formatPrefix + message)
	}
}

func (l *L) Criticalf(format string, arguments ...interface{}) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.CriticalValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Critical(s)
	}
}

func (l *L) Criticalc(closure func() string) {
	if !globalData.initialised || nil == l {
		panic("logger is not initialised")
	}
	if l.levelNumber <= level.CriticalValue {
		_ = l.log.Critical(l.formatPrefix + closure())
	}
}
