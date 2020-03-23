// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2020 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package writer

type Type int

const (
	Stdout Type = iota
	Seelog
)

// Writer - provides both logger & info interfaces to operate
type Writer interface {
	Logger
	Info
}

// Logger - log messages
type Logger interface {
	// Log a simple string
	// e.g.
	//   log.Trace("a log message")
	Trace(string)

	// Log a formatted string with arguments like fmt.Sprintf()
	// e.g.
	//   log.Tracef("the value = %d", xValue)
	Tracef(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Tracec(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Tracec(func() string)

	// Log a simple string
	// e.g.
	//   log.Debug("a log message")
	Debug(string)

	// Log a formatted string with arguments lige fmt.Sprintf()
	// e.g.
	//   log.Debugf("the value = %d", xValue)
	Debugf(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Debugc(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Debugc(func() string)

	// Log a simple string
	// e.g.
	//   log.Info("a log message")
	Info(string)

	// Log a formatted string with arguments lige fmt.Sprintf()
	// e.g.
	//   log.Infof("the value = %d", xValue)
	Infof(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Infoc(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Infoc(func() string)

	// Log a simple string
	// e.g.
	//   log.Warn("a log message")
	Warn(string)

	// Log a formatted string with arguments lige fmt.Sprintf()
	// e.g.
	//   log.Warnf("the value = %d", xValue)
	Warnf(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Warnc(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Warnc(func() string)

	// Log a simple string
	// e.g.
	//   log.Error("a log message")
	Error(string)

	// Log a formatted string with arguments lige fmt.Sprintf()
	// e.g.
	//   log.Errorf("the value = %d", xValue)
	Errorf(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Errorc(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Errorc(func() string)

	// Log a simple string
	// e.g.
	//   log.Critical("a log message")
	Critical(string)

	// Log a formatted string with arguments lige fmt.Sprintf()
	// e.g.
	//   log.Criticalf("the value = %d", xValue)
	Criticalf(string, ...interface{})

	// Log from a closure, any function returning a string is suitable
	// and the closure will only be executed if the log level is low enough.
	// This allows complex
	// e.g.
	//   log.Criticalc(func() string {
	//       return fmt.Sprintf("the sin(%f) = %f", x, math.sin(x))
	//   })
	Criticalc(func() string)
}

// Info - provides self information
type Info interface {
	Tag() string
	Level() string
	UpdateLevel(string) error
}

// New - create writer
func New(t Type, args ...interface{}) (Writer, error) {
	switch t {
	case Stdout:
		return newStdout()
	case Seelog:
		return newSeelog(args)
	default:
		return newStdout()
	}
}
