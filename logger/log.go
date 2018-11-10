package logger

import (
	"fmt"
	l "log"
	"os"
	"time"
)

const prefix = "[koazee] "
const timeFormat = "15:04:05.000000"

// Enabled is used to enabled or disable (when its value is false the koazee log
var Enabled = false

var log = l.New(os.Stdout, prefix, 0)

// DebugInfo prints the received message
func DebugInfo(format string, args ...interface{}) {
	if Enabled {
		log.Printf(fmt.Sprintf("%v %v", time.Now().Format(timeFormat), fmt.Sprintf(format, args...)))
	}
}
