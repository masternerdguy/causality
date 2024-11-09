package main

import (
	"causality/arena"
	"causality/lib"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// entry point
func main() {
	// get args
	cliArgs := os.Args[1:]

	// ensure 3 args
	if len(cliArgs) != 1 {
		panic("No args! Please include the path to your code file.")
	}

	// parse args
	fp := cliArgs[0]

	// load program
	p := lib.ParseFile(fp)
	log.Print(p)

	// init globals
	lib.InitGlobals(len(p))

	// use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// initialize arena and framebuffer
	arena.InitArena(p)

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
			fmt.Println("Sentinel value underflowed - program halted with above output! If this causes problems, just remove it from the code in main.go :)")

			// exit
			break
		}
	}
}
