package faultline

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	SetLogger(log.New(os.Stderr, "faultline: ", log.LstdFlags))
}

func SetLogger(l *log.Logger) {
	logger = l
}
