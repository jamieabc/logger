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
	TraceKey    = "trace"
	DebugKey    = "debug"
	InfoKey     = "info"
	WarnKey     = "warn"
	ErrorKey    = "error"
	CriticalKey = "critical"
	OffKey      = "off"
)

// This needs to correspond to seelog levels
var ValidLevels = map[string]int{
	TraceKey:    TraceValue,
	DebugKey:    DebugValue,
	InfoKey:     InfoValue,
	WarnKey:     WarnValue,
	ErrorKey:    ErrorValue,
	CriticalKey: CriticalValue,
	OffKey:      OffValue,
}
