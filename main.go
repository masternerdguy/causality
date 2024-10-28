package main

import (
	"flux/arena"
	"flux/lib"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

// entry point
func main() {
	// get args
	cliArgs := os.Args[1:]

	// ensure 3 args
	if len(cliArgs) != 3 {
		panic("No args! Please include <arena_length> <elapsed seed> <flux seed>")
	}

	// parse args
	al, _ := strconv.Atoi(cliArgs[0])
	es, _ := strconv.Atoi(cliArgs[1])
	fs, _ := strconv.Atoi(cliArgs[2])

	// setup
	lib.InitGlobals(al, es, fs)

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
