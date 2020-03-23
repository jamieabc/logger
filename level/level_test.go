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
	assert.Equal(t, level.TraceValue, level.ValidLevels[level.Trace], "wrong trace")
	assert.Equal(t, level.DebugValue, level.ValidLevels[level.Debug], "wrong trace")
	assert.Equal(t, level.InfoValue, level.ValidLevels[level.Info], "wrong trace")
	assert.Equal(t, level.WarnValue, level.ValidLevels[level.Warn], "wrong trace")
	assert.Equal(t, level.ErrorValue, level.ValidLevels[level.Error], "wrong trace")
	assert.Equal(t, level.CriticalValue, level.ValidLevels[level.Critical], "wrong trace")
	assert.Equal(t, level.OffValue, level.ValidLevels[level.Off], "wrong trace")
}
