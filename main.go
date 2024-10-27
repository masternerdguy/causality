package main

import "flux/arena"

// entry point
func main() {
	// initialize arena and framebuffer
	arena.InitArena()

	// draw empty frame
	arena.DrawFrame()
}
