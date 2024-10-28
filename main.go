package main

import (
	"flux/arena"
	"math/rand"
	"runtime"
	"time"
)

const RNG_SEED = 0

// entry point
func main() {
	// single core
	runtime.GOMAXPROCS(1)

	// seed rng for consistent results
	rand.Seed(RNG_SEED)

	// initialize arena and framebuffer
	arena.InitArena()

	// draw empty frame
	arena.DrawFrame()

	// don't exit
	for {
		time.Sleep(0)
	}
}
