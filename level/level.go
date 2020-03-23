// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2020 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package level

// simple ordering to allow <= to decide if log will be output
const (
	_ = iota
	TraceValue
	DebugValue
	InfoValue
	WarnValue
	ErrorValue
	CriticalValue
	OffValue
)

const (
	Trace    = "trace"
	Debug    = "debug"
	Info     = "info"
	Warn     = "warn"
	Error    = "error"
	Critical = "critical"
	Off      = "off"
)

// This needs to correspond to seelog levels
var ValidLevels = map[string]int{
	Trace:    TraceValue,
	Debug:    DebugValue,
	Info:     InfoValue,
	Warn:     WarnValue,
	Error:    ErrorValue,
	Critical: CriticalValue,
	Off:      OffValue,
}
