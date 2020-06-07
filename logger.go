package logger

import (
	"io"
	"log"
)

const (
	TRACE = iota
	DEBUG = iota
	INFO = iota
	ERROR = iota
)

var (
	trace *log.Logger
	info *log.Logger
	debug *log.Logger
	err * log.Logger
)

type Logger struct {
	o io.Writer
	loglevel int
	flag int
}

func New(o io.Writer, level int, flag int) *Logger{
	switch level {
	case 0:
		trace = log.New(o, "TRACE", log.Lmsgprefix|flag)
		fallthrough
	case 1:
		debug = log.New(o,"DEBUG", log.Lmsgprefix|flag)
		fallthrough
	case 2:
		info = log.New(o,"INFO", log.Lmsgprefix|flag)
		fallthrough
	case 3:
		err = log.New(o, "ERROR", log.Lmsgprefix|flag)
	default:
		log.Fatalf("Please provide a log level")
	}
	return &Logger{o,level,flag}
}

func (l *Logger) Info(a interface{}) {
	if info != nil {
		info.Printf(":: %s\n", a)
	}
}

func (l *Logger) Debug(a interface{}) {
	if debug != nil {
		debug.Printf(":: %s\n",a)
	}
}

func (l *Logger) Error(a interface{}) {
	if err != nil {
		err.Fatalf(":: %s\n", a)
	}
}

func (l *Logger) Trace(a interface{}) {
	if trace != nil {
		trace.Printf(":: %s\n",a)
	}
}