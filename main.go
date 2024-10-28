package main

import (
	"flux/arena"
	"flux/lib"
	"fmt"
	"runtime"
	"time"
)

// entry point
func main() {
	// use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// initialize arena and framebuffer
	arena.InitArena()

	// draw empty frame
	arena.DrawFrame()

	// wait
	wait()
}

func wait() {
	// don't exit
	for {
		// don't peg cpu and minimize interference with scheduling
		time.Sleep(0)

		// decrement sentinel value
		lib.Sentinel--

		// exit if underflowed - we are halted
		if lib.Sentinel <= lib.SENTINEL_UNDERFLOW {
			// notify user
			fmt.Println("Sentinel value underflowed - program halted with above output!")

			// exit
			break
		}
	}
}
