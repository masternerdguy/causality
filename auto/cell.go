package auto

import (
	"flux/lib"
)

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

	// internal cycle counter
	i_c int
}

func (c *Cell) SetFlux(f int) {
	// store flux internally
	c.i_f = f
}

func (c *Cell) SetAge(a int) {
	// store "age" (elapsed cycles) internally
	c.i_c = a
}

func (c *Cell) Listen() {
	// loop
	for {
		// receive updates from causal past
		v := <-c.Past

		// increment cycle count
		c.i_c++

		// accumulate past in internal state
		c.i_f += v

		// toroidal bound check
		if c.i_f < lib.FLUX_LOW {
			c.i_f = lib.FLUX_HIGH
		}

		if c.i_f > lib.FLUX_HIGH {
			c.i_f = lib.FLUX_LOW
		}

		// determine value to propagate to the causal future
		fv := 0
		var sv string

		if c.i_f >= lib.FLUX_BOUNDARY {
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

		// exit if cycle count exceeded
		if c.i_c > lib.MAX_CYCLES {
			break
		}
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
