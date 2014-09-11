package logger

import (
	"fmt"
	"strings"
	"time"
)

type Severity int32

type Logger struct {
	level Severity
}

const (
	TRACE Severity = 0
	DEBUG          = 1
	INFO           = 2
	WARN           = 3
	ERROR          = 4
	FATAL          = 5
)

var severities = []string{
	TRACE: "TRACE",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

func severityByName(s string) Severity {
	s = strings.ToUpper(s)
	for i, name := range severities {
		if name == s {
			return Severity(i)
		}
	}
	return -1
}

func (log *Logger) SetLevelWithDefault(lv, defaultLv string) {
	err := log.SetLevel(lv)
	if err != nil {
		log.SetLevel(defaultLv)
		log.Warn("log level not valid. use default level: %s", defaultLv)
	}
}

func (log *Logger) SetLevel(lv string) error {
	if lv == "" {
		return fmt.Errorf("log level is blank")
	}

	sev := severityByName(lv)
	if sev < 0 {
		return fmt.Errorf("log level setting error")
	}

	log.level = sev

	return nil
}

func (log *Logger) Trace(format string, v ...interface{}) {
	if log.level <= TRACE {
		log.println(TRACE, format, v...)
	}
}

func (log *Logger) Debug(format string, v ...interface{}) {
	if log.level <= DEBUG {
		log.println(DEBUG, format, v...)
	}
}

func (log *Logger) Info(format string, v ...interface{}) {
	if log.level <= INFO {
		log.println(INFO, format, v...)
	}
}

func (log *Logger) Warn(format string, v ...interface{}) {
	if log.level <= WARN {
		log.println(WARN, format, v...)
	}
}

func (log *Logger) Error(format string, v ...interface{}) {
	if log.level <= ERROR {
		log.println(ERROR, format, v...)
	}
}

func (log *Logger) Fatal(format string, v ...interface{}) {
	if log.level <= FATAL {
		log.println(FATAL, format, v...)
	}
}

func (log *Logger) println(s Severity, format string, v ...interface{}) { //severities
	fmt.Printf(time.Now().Format("2006/01/02 15:04:05")+" ["+severities[s]+"] "+format+"\n", v...)
}
