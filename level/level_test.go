// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2020 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package level_test

import (
	"github.com/bitmark-inc/logger/level"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidLevels(t *testing.T) {
	assert.Equal(t, level.TraceValue, level.ValidLevels[level.TraceKey], "wrong trace")
	assert.Equal(t, level.DebugValue, level.ValidLevels[level.DebugKey], "wrong trace")
	assert.Equal(t, level.InfoValue, level.ValidLevels[level.InfoKey], "wrong trace")
	assert.Equal(t, level.WarnValue, level.ValidLevels[level.WarnKey], "wrong trace")
	assert.Equal(t, level.ErrorValue, level.ValidLevels[level.ErrorKey], "wrong trace")
	assert.Equal(t, level.CriticalValue, level.ValidLevels[level.CriticalKey], "wrong trace")
	assert.Equal(t, level.OffValue, level.ValidLevels[level.OffKey], "wrong trace")
}
