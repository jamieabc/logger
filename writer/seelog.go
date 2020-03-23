// Copyright (c) 2014-2018 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// Generated file: Do _NOT_ Modify
// Generated on: 2017-09-19T11:23:21+08:00

package writer

import (
	"fmt"
	"github.com/bitmark-inc/logger/level"
	"github.com/cihub/seelog"
	"sync"
)

const (
	argCount = 4
)

type thirdPartyLoggerSeelog struct {
	sync.Mutex
	tag          string
	formatPrefix string
	textPrefix   string
	level        string
	levelNumber  int
	log          seelog.LoggerInterface
}

func (l *thirdPartyLoggerSeelog) Trace(message string) {
	if l.levelNumber <= level.TraceValue {
		l.log.Trace(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Tracef(format string, arguments ...interface{}) {
	if l.levelNumber <= level.TraceValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Trace(s)
	}
}

func (l *thirdPartyLoggerSeelog) Tracec(closure func() string) {
	if l.levelNumber <= level.TraceValue {
		l.log.Trace(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Debug(message string) {
	if l.levelNumber <= level.DebugValue {
		l.log.Debug(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Debugf(format string, arguments ...interface{}) {
	if l.levelNumber <= level.DebugValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Debug(s)
	}
}

func (l *thirdPartyLoggerSeelog) Debugc(closure func() string) {
	if l.levelNumber <= level.DebugValue {
		l.log.Debug(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Info(message string) {
	if l.levelNumber <= level.InfoValue {
		l.log.Info(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Infof(format string, arguments ...interface{}) {
	if l.levelNumber <= level.InfoValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		l.log.Info(s)
	}
}

func (l *thirdPartyLoggerSeelog) Infoc(closure func() string) {
	if l.levelNumber <= level.InfoValue {
		l.log.Info(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Warn(message string) {
	if l.levelNumber <= level.WarnValue {
		_ = l.log.Warn(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Warnf(format string, arguments ...interface{}) {
	if l.levelNumber <= level.WarnValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Warn(s)
	}
}

func (l *thirdPartyLoggerSeelog) Warnc(closure func() string) {
	if l.levelNumber <= level.WarnValue {
		_ = l.log.Warn(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Error(message string) {
	if l.levelNumber <= level.ErrorValue {
		_ = l.log.Error(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Errorf(format string, arguments ...interface{}) {
	if l.levelNumber <= level.ErrorValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Error(s)
	}
}

func (l *thirdPartyLoggerSeelog) Errorc(closure func() string) {
	if l.levelNumber <= level.ErrorValue {
		_ = l.log.Error(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Critical(message string) {
	if l.levelNumber <= level.CriticalValue {
		_ = l.log.Critical(l.formatPrefix + message)
	}
}

func (l *thirdPartyLoggerSeelog) Criticalf(format string, arguments ...interface{}) {
	if l.levelNumber <= level.CriticalValue {
		s := fmt.Sprintf(l.formatPrefix+format, arguments...)
		_ = l.log.Critical(s)
	}
}

func (l *thirdPartyLoggerSeelog) Criticalc(closure func() string) {
	if l.levelNumber <= level.CriticalValue {
		_ = l.log.Critical(l.formatPrefix + closure())
	}
}

func (l *thirdPartyLoggerSeelog) Tag() string {
	return l.tag
}

func (l *thirdPartyLoggerSeelog) Level() string {
	return l.level
}

func (l *thirdPartyLoggerSeelog) UpdateLevel(newLevel string) error {
	if num, ok := level.ValidLevels[newLevel]; ok {
		l.level = newLevel
		l.levelNumber = num
		return nil
	}
	return fmt.Errorf("level %s invalid", newLevel)
}

// make sure enough arguments are provided
func newSeelog(args []interface{}) (Writer, error) {
	if len(args) != argCount {
		return nil, fmt.Errorf("not enought arguments for seelog, expect %d but get %d", argCount, len(args))
	}

	l := args[3].(string)
	var levelNumber int
	if num, ok := level.ValidLevels[l]; !ok {
		return nil, fmt.Errorf("level %s invalid", l)
	} else {
		levelNumber = num
	}

	return &thirdPartyLoggerSeelog{
		tag:          args[0].(string),
		formatPrefix: args[1].(string),
		textPrefix:   args[2].(string),
		level:        l,
		levelNumber:  levelNumber,
		log:          seelog.Current,
	}, nil
}
