package logger

import (
	"fmt"
	l "log"
	"os"
	"time"
)

const prefix = "[koazee] "
const timeFormat = "15:04:05.000000"

// Enabled true means log is enabled
var Enabled = false

var log = l.New(os.Stdout, prefix, 0)

func DebugInfo(format string, args ...interface{}) {
	if Enabled {
		log.Printf(fmt.Sprintf("%v %v", time.Now().Format(timeFormat), fmt.Sprintf(format, args...)))
	}
}
