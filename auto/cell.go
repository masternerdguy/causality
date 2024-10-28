package auto

import (
	"math/rand"
)

// boundary for transition from "alive" to "dead"
const FLUX_BOUNDARY = 4

type Cell struct {
	// global x position used for initial setup and rendering only
	G_X int

	// global y position used for initial setup and rendering only
	G_Y int

	// input channel for events from the causal past
	Past chan (int)

	// output channels for events going into the causal future
	Future []*chan (int)

	// diagnostic channel to update framebuffer
	Render *chan (string)

	// internal flux state
	I_f int
}

func (c *Cell) Randomize() {
	// randomize internal flux to be over or under the boundary
	c.I_f = int((rand.Float32()-0.5)*2.5) + FLUX_BOUNDARY
}

type RenderCell struct {
	// global x position used for initial setup and rendering only
	G_X int

	// global y position used for initial setup and rendering only
	G_Y int

	// state character
	G_S string

	// update channel
	Update chan (string)

	// render channel to trigger redraw of arena
	Render *chan (string)
}

func (c *RenderCell) Listen() {
	for {
		// receive display updates from channel
		f := <-c.Update

		// update stored value
		c.G_S = f

		// pass upwards for global redrawing
		*c.Render <- c.G_S
	}
}
