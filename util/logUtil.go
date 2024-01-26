package util

import (
	"log"
)

func BlueLog(format string, a ...interface{}) {
	log.Printf("\x1b[94m"+format+"\x1b[0m", a...)
}

func GreenLog(format string, a ...interface{}) {
	log.Printf("\x1b[92m"+format+"\x1b[0m", a...)
}

func YellowLog(format string, a ...interface{}) {
	log.Printf("\x1b[93m"+format+"\x1b[0m", a...)
}

func RedLog(format string, a ...interface{}) {
	log.Printf("\x1b[91m"+format+"\x1b[0m", a...)
}
