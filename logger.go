package simpletracer

import "log"

type Logger interface {
	Printf(format string, a ...interface{})
}

type defaultLogger struct {}

func (l defaultLogger) Printf(format string, a ...interface{}) {
	log.Printf(format, a...)
}