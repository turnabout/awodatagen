package utilities

import (
	"fmt"
	"os"
	"runtime/debug"
)

func LogFatalIfErr(err error) {
	if err != nil {
		LogFatalF("Error: %s", err.Error())
	}
}

// Helper to log a fatal error with a stack trace, then exit the program
func LogFatalF(format string, a ...interface{}) {
	fmt.Println("Fatal error:")
	fmt.Printf(format, a)
	fmt.Println("Stack trace:")
	debug.PrintStack()
	os.Exit(1)
}

// Counts amount of bits in a number (hardcoded for 32-bit numbers)
func CountBits(n uint) uint {
	n = ((0xaaaaaaaa & n) >> 1) + (0x55555555 & n)
	n = ((0xcccccccc & n) >> 2) + (0x33333333 & n)
	n = ((0xf0f0f0f0 & n) >> 4) + (0x0f0f0f0f & n)
	n = ((0xff00ff00 & n) >> 8) + (0x00ff00ff & n)
	n = ((0xffff0000 & n) >> 16) + (0x0000ffff & n)
	return n
}
