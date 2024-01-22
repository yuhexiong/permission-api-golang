package util

import (
	"log"
)

func BlueLog(a ...interface{}) {
	log.Printf("\x1b[94m%v\x1b[0m", a...)
}

func GreenLog(a ...interface{}) {
	log.Printf("\x1b[92m%v\x1b[0m", a...)
}

func YellowLog(a ...interface{}) {
	log.Printf("\x1b[93m%v\x1b[0m", a...)
}

func RedLog(a ...interface{}) {
	log.Printf("\x1b[91m%v\x1b[0m", a...)
}
