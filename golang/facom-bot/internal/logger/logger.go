package logger

import (
	"log"
	"os"
)

// Info outputs logs to stdout and prefixes them with "[INFO]"
var Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix)

// Error outputs logs to stderr and prefixes them with "[ERROR]"
var Error = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix)
