package logger

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

func (logLvl LogLevel) IsValid() bool {
	switch logLvl {
	case LogLevelError,
		LogLevelWarning,
		LogLevelInfo:
		return true
	default:
		return false

	}
}

type ExtendedLog struct {
	*log.Logger
	logLevel LogLevel
}

func (log *ExtendedLog) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}

	log.logLevel = logLvl
}

func (log ExtendedLog) Infoln(msg string) {
	log.println(LogLevelInfo, "[INFO]", msg)
}

func (log ExtendedLog) Warnln(msg string) {
	log.println(LogLevelWarning, "[WARN]", msg)
}

func (log ExtendedLog) Errorln(msg string) {
	log.println(LogLevelError, "[ERR]", msg)
}

func (log *ExtendedLog) println(logLvl LogLevel, prefix, msg string) {
	if log.logLevel < logLvl {
		return
	}

	log.Logger.Println(prefix + msg)
}

func NewExtendedLog() *ExtendedLog {
	return &ExtendedLog{
		Logger:   log.New(os.Stderr, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}
