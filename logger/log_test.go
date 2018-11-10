package logger_test

import (
	"testing"

	"github.com/wesovilabs/koazee/logger"
)

func TestDebugInfo(t *testing.T) {
	logger.DebugInfo("Say hello to %s", "Sally")
	logger.Enabled = false
	logger.DebugInfo("Say hello to %s", "Sally")
}
