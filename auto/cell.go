package auto

import (
	"flux/lib"
)

// boundary for transition from "alive" to "dead"
const FLUX_BOUNDARY = 1
const FLUX_LOW = 0
const FLUX_HIGH = 1

type Cell struct {
	// global x position used for initial setup and rendering only
	G_X int

	// global y position used for initial setup and rendering only
	G_Y int

	// input channel for events from the causal past
	Past chan (int)

	// output channels for events going into the causal future
	Future []chan (int)

	// diagnostic channel to update framebuffer
	Render chan (string)

	// internal flux state
	i_f int
}

func (c *Cell) Randomize() {
	// randomize internal flux to be over or under the boundary
	//c.i_f = int((rand.Float32()-0.5)*2.5) + FLUX_BOUNDARY

	c.i_f = 0
}

func (c *Cell) Listen() {
	for {
		// receive updates from causal past
		v := <-c.Past

		// accumulate in internal state
		c.i_f += v

		// toroidal bound check
		if c.i_f < FLUX_LOW {
			c.i_f = FLUX_HIGH
		}

		if c.i_f > FLUX_HIGH {
			c.i_f = FLUX_LOW
		}

		// determine value to propagate to the causal future
		fv := 0
		var sv string

		if c.i_f >= FLUX_BOUNDARY {
			// cell is "alive" - pass 1
			fv = 1

			// display as alive
			sv = "*"
		} else {
			// cell is "dead" - pass -1
			fv = -1

			// display as dead
			sv = " "
		}

		// propagate to the causal future
		for _, q := range c.Future {
			q <- fv
		}

		// and upwards to the "framebuffer"
		c.Render <- sv
	}
}

type RenderCell struct {
	// global x position used for initial setup and rendering only
	G_X int

	// global y position used for initial setup and rendering only
	G_Y int

	// display state character
	G_S string

	// update channel
	Update chan (string)

	// render channel to trigger redraw of arena
	Render chan (lib.ArenaChange[int, int, string])
}

func (c *RenderCell) Listen() {
	for {
		// receive display updates from channel
		f := <-c.Update

		// update stored display value
		c.G_S = f

		// build message for global rendering
		m := lib.ArenaChange[int, int, string]{
			X: c.G_X,
			Y: c.G_Y,
			S: f,
		}

		// pass upwards for global redrawing
		c.Render <- m
	}
}
