package arena

import (
	"flux/auto"
	"fmt"
)

var arena = make([][]*auto.Cell, ARENA_LENGTH)
var arenaBuffer = make([][]*auto.RenderCell, ARENA_LENGTH)
var arenaUpdate = make(chan string)

const ARENA_LENGTH = 25

func InitArena() {
	/* Boilerplate setup */

	// initialize global redraw channel for diagnostic output
	arenaUpdate = make(chan string)

	// loop over rows
	for x := 0; x < ARENA_LENGTH; x++ {
		// init cells
		arena[x] = make([]*auto.Cell, ARENA_LENGTH)
		arenaBuffer[x] = make([]*auto.RenderCell, ARENA_LENGTH)

		// loop over cells
		for y := 0; y < ARENA_LENGTH; y++ {
			// arena cell
			arena[x][y] = &auto.Cell{
				// global coordinates used only for setup
				G_X: x,
				G_Y: y,

				// initialize channel for causal past
				Past: make(chan int),

				// initialize empty future
				Future: make([]*chan int, 0),
			}

			// randomize arena cell flux
			arena[x][y].Randomize()

			// "framebuffer" display cell
			arenaBuffer[x][y] = &auto.RenderCell{
				// global coordinates used only for diagnostic output
				G_X: x,
				G_Y: y,
				// stash initial display value
				G_S: fmt.Sprintf("%d", arena[x][y].I_f),
				// make channel to update display value
				Update: make(chan string),
				// hook channel to perform a redraw
				Render: &arenaUpdate,
			}

			// hook update channel into cell so we can see the simulation
			arena[x][y].Render = &arenaBuffer[x][y].Update

			// start listening for cell changes
			go func() {
				arenaBuffer[x][y].Listen()
			}()
		}
	}

	/* Causal connectivity setup */

	// loop over rows
	for x := 0; x < ARENA_LENGTH; x++ {
		// loop over cells
		for y := 0; y < ARENA_LENGTH; y++ {
			// get future coordinates for this cell
			ft := descendantCoordinates(x, y)

			fl := ft[0]
			fr := ft[1]

			// store future channels to create causal connections
			cl := &arena[fl[0]][fl[1]].Past
			cr := &arena[fr[0]][fr[1]].Past

			arena[x][y].Future = append(arena[x][y].Future, cl)
			arena[x][y].Future = append(arena[x][y].Future, cr)
		}
	}
}

func DrawFrame() {
	// indicate start of frame
	fmt.Println("Frame\n")

	// loop over rows
	for x := 0; x < ARENA_LENGTH; x++ {
		// loop over columns
		for y := 0; y < ARENA_LENGTH; y++ {
			// draw cell
			fmt.Print(arenaBuffer[x][y].G_S)

			// spacer
			fmt.Print(" ")
		}

		// line break
		fmt.Println()
	}

	// indicate end of frame
	fmt.Println("\n")
}

func descendantCoordinates(x int, y int) [][]int {
	// empty slice to store output
	o := make([][]int, 0)

	// "lefthand" descendant
	lx := x - 1
	ly := y + 1

	// toroidal bound checks
	if lx < 0 {
		lx = ARENA_LENGTH - 1
	}

	if ly < 0 {
		ly = ARENA_LENGTH - 1
	}

	if lx == ARENA_LENGTH {
		lx = 0
	}

	if ly == ARENA_LENGTH {
		ly = 0
	}

	// "righthand" descendant
	rx := x + 1
	ry := y + 1

	// toroidal bound checks
	if rx < 0 {
		rx = ARENA_LENGTH - 1
	}

	if ry < 0 {
		ry = ARENA_LENGTH - 1
	}

	if rx == ARENA_LENGTH {
		rx = 0
	}

	if ry == ARENA_LENGTH {
		ry = 0
	}

	// build "lefthand" coordinate
	lp := make([]int, 2)
	lp = append(lp, lx)
	lp = append(lp, ly)

	// build "righthand" coordinate
	rp := make([]int, 2)
	rp = append(rp, rx)
	rp = append(rp, ry)

	// store coordinates
	o = append(o, lp)
	o = append(o, rp)

	// return result
	return o
}
