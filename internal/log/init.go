package log

import (
	"log"
	"os"
	"strings"
	"time"
)

type LogType byte

const (
	File LogType = iota
	Stdout
)

type LogDest struct {
	Type LogType
	Path string `yaml:"path"`
	f    *os.File
}

type Logger struct {
	destArr []LogDest
}

func NewLogger(dests []LogDest) *Logger {
	return &Logger{destArr: dests}
}

func (l *Logger) Println(msg string) {
	for _, dest := range l.destArr {
		switch dest.Type {
		case File:
			_, err := dest.f.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + msg + "\n")
			if err != nil {
				continue
			}
		case Stdout:
			log.Println(msg)
		}
	}
}

func ParseLogDest(path string) []LogDest {
	ldst := make([]LogDest, 0)
	for _, el := range strings.Split(path, " ") {
		switch el {
		case "stdout":
			ldst = append(ldst, LogDest{Type: Stdout})
		default:
			ldst = append(ldst, LogDest{Type: File, Path: el})
			f, err := os.OpenFile(el, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
				continue
			} else {
				ldst[len(ldst)-1].f = f
			}
		}
	}
	return ldst
}

func (l *Logger) Close() {
	for _, dest := range l.destArr {
		if dest.f != nil {
			dest.f.Close()
		}
	}
}
