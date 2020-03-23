// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2020 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package writer

import (
	"fmt"
	"github.com/bitmark-inc/logger/level"
	"os"
	"strings"
	"time"
)

type stdout struct {
	tag          string
	formatPrefix string
	textPrefix   string
	level        string
	levelNumber  int
}

func newStdout() (Writer, error) {
	return &stdout{}, nil
}

func std(message, level, tag string) {
	now := time.Now()
	fmt.Fprintf(os.Stdout, "%s [%s] %s %s\n",
		now, strings.ToUpper(level), tag, message)
}

func (s *stdout) Trace(message string) {
	if s.levelNumber <= level.TraceValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Tracef(format string, arguments ...interface{}) {
	if s.levelNumber <= level.TraceValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Tracec(closure func() string) {
	if s.levelNumber <= level.TraceValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Debug(message string) {
	if s.levelNumber <= level.DebugValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Debugf(format string, arguments ...interface{}) {
	if s.levelNumber <= level.DebugValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Debugc(closure func() string) {
	if s.levelNumber <= level.DebugValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Info(message string) {
	if s.levelNumber <= level.InfoValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Infof(format string, arguments ...interface{}) {
	if s.levelNumber <= level.InfoValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Infoc(closure func() string) {
	if s.levelNumber <= level.InfoValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Warn(message string) {
	if s.levelNumber <= level.WarnValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Warnf(format string, arguments ...interface{}) {
	if s.levelNumber <= level.WarnValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Warnc(closure func() string) {
	if s.levelNumber <= level.WarnValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Error(message string) {
	if s.levelNumber <= level.ErrorValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Errorf(format string, arguments ...interface{}) {
	if s.levelNumber <= level.ErrorValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Errorc(closure func() string) {
	if s.levelNumber <= level.ErrorValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Critical(message string) {
	if s.levelNumber <= level.CriticalValue {
		std(s.formatPrefix+message, s.level, s.tag)
	}
}

func (s *stdout) Criticalf(format string, arguments ...interface{}) {
	if s.levelNumber <= level.CriticalValue {
		msg := fmt.Sprintf(s.formatPrefix+format, arguments...)
		std(msg, s.level, s.tag)
	}
}

func (s *stdout) Criticalc(closure func() string) {
	if s.levelNumber <= level.CriticalValue {
		std(s.formatPrefix+closure(), s.level, s.tag)
	}
}

func (s *stdout) Tag() string {
	return s.tag
}

func (s *stdout) Level() string {
	return s.level
}

func (s *stdout) UpdateLevel(newLevel string) error {
	if num, ok := level.ValidLevels[newLevel]; ok {
		s.levelNumber = num
		s.level = newLevel
		return nil
	}
	return fmt.Errorf("level %s invalid", newLevel)
}
