package logger_test

import (
	"github.com/wesovilabs/koazee/logger"
	"testing"
)

func TestDebugInfo(t *testing.T) {
	logger.DebugInfo("my-stream", "Say hello to %s", "Sally")
	logger.Enabled=false
	logger.DebugInfo("my-stream", "Say hello to %s", "Sally")
}
