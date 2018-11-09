package logger

import (
	"fmt"
	l "log"
	"os"
	"time"
)

const prefix = "[koazee] "
const timeFormat = "15:04:05.000000"

var log *l.Logger

func init() {
	log = l.New(os.Stdout, prefix, 0)
}

func DebugInfo(traceID string, format string, args ...interface{}) {
	log.Printf(fmt.Sprintf("%v %s %v", time.Now().Format(timeFormat), traceID, fmt.Sprintf(format, args...)))
}
