package arena

import (
	"flux/auto"
	"flux/lib"
	"fmt"
)

// where the simulation happens
var arena [][]*auto.Cell

// intermediary to deliver updates to display
var arenaBuffer [][]*auto.RenderCell

// what gets displayed
var arenaDraw [][]string

// channel to update what gets displayed
var arenaUpdate chan lib.ArenaChange[int, int, string]

// number of updates since start
var frameCounter int64

func InitArena() {
	/* Recompute vars */
	arena = make([][]*auto.Cell, lib.ARENA_LENGTH)
	arenaBuffer = make([][]*auto.RenderCell, lib.ARENA_LENGTH)
	arenaDraw = make([][]string, lib.ARENA_LENGTH)

	/* Cell and "framebuffer" setup */

	// initialize global redraw channel for diagnostic output
	arenaUpdate = make(chan lib.ArenaChange[int, int, string])

	// loop over rows
	for x := 0; x < lib.ARENA_LENGTH; x++ {
		// init cells
		arena[x] = make([]*auto.Cell, lib.ARENA_LENGTH)
		arenaBuffer[x] = make([]*auto.RenderCell, lib.ARENA_LENGTH)
		arenaDraw[x] = make([]string, lib.ARENA_LENGTH)

		// loop over cells
		for y := 0; y < lib.ARENA_LENGTH; y++ {
			// arena cell
			arena[x][y] = &auto.Cell{
				// global coordinates used only for setup
				G_X: x,
				G_Y: y,

				// initialize channel for causal past
				Past: make(chan int, lib.ARENA_AREA+1),

				// initialize empty future
				Future: make([]chan int, 0),
			}

			// "framebuffer" display cell
			arenaBuffer[x][y] = &auto.RenderCell{
				// global coordinates used only for diagnostic output
				G_X: x,
				G_Y: y,
				// make channel to update display value
				Update: make(chan string),
				// hook channel to perform a redraw
				Render: arenaUpdate,
			}

			// hook update channel into cell so we can see the simulation
			arena[x][y].Render = arenaBuffer[x][y].Update

			// start listening for cell changes
			go func() {
				arenaBuffer[x][y].Listen()
			}()
		}
	}

	// loop over rows
	for x := 0; x < lib.ARENA_LENGTH; x++ {
		// loop over cells
		for y := 0; y < lib.ARENA_LENGTH; y++ {
			// get causal future coordinates for this cell
			ft := descendantCoordinates(x, y)

			fl := ft[0]
			fr := ft[1]

			// store future channels to create causal connections
			cl := arena[fl[0]][fl[1]].Past
			cr := arena[fr[0]][fr[1]].Past

			arena[x][y].Future = append(arena[x][y].Future, cl)
			arena[x][y].Future = append(arena[x][y].Future, cr)

			// initial value
			arena[x][y].SetAge(1)

			// start listening for input from the causal past
			go func() {
				arena[x][y].Listen()
			}()
		}
	}

	/* Default blank frame */

	// loop over rows
	for x := 0; x < lib.ARENA_LENGTH; x++ {
		// loop over cells
		for y := 0; y < lib.ARENA_LENGTH; y++ {
			// initial char so we can be sure things are actually changing properly
			arenaDraw[x][y] = "_"
		}
	}

	/* Global watcher for arena changes */

	go func() {
		for {
			// receive update from "framebuffer" cell
			m := <-arenaUpdate
			arenaDraw[m.X][m.Y] = m.S

			// draw frame
			DrawFrame()
		}
	}()

	/* Erase global coordinates in simulation */

	// loop over rows
	for x := 0; x < lib.ARENA_LENGTH; x++ {
		// loop over cells
		for y := 0; y < lib.ARENA_LENGTH; y++ {
			// erase global coordinates so they can't be used
			arena[x][y].G_X = -1
			arena[x][y].G_Y = -1
		}
	}

	/* Initial "spark" to start propagation */

	// get coordinates of entry point
	ex := lib.ARENA_LENGTH / 2
	ey := lib.ARENA_LENGTH / 2

	// send initial impulse to network
	arena[ex][ey].Past <- 0
}

func DrawFrame() {
	fmt.Printf("Update %d\n", frameCounter)

	// loop over rows
	for x := 0; x < lib.ARENA_LENGTH; x++ {
		// loop over columns
		for y := 0; y < lib.ARENA_LENGTH; y++ {
			// draw cell
			fmt.Print(arenaDraw[y][x]) // swapping coordinates to transform from worldspace to screenspace

			// spacer
			fmt.Print(" ")
		}

		// line break
		fmt.Println()
	}

	// increment counter
	frameCounter++

	// update sentinel value so the program doesn't die
	lib.Sentinel = lib.SENTINEL_RESET

	// end of frame
	fmt.Println()
}

// Returns the immediate descendant coordinates of the provided coordinates
func descendantCoordinates(x int, y int) [][]int {
	// empty slice to store output
	o := make([][]int, 0)

	// future time
	gy := y - 1

	// toroidal bound checks (time)
	if gy < 0 {
		gy = lib.ARENA_LENGTH - 1
	}

	if gy == lib.ARENA_LENGTH {
		gy = 0
	}

	// "lefthand" descendant
	lx := x - 1

	// toroidal bound checks (space)
	if lx < 0 {
		lx = lib.ARENA_LENGTH - 1
	}

	if lx == lib.ARENA_LENGTH {
		lx = 0
	}

	// "righthand" descendant
	rx := x + 1

	// toroidal bound checks (space)
	if rx < 0 {
		rx = lib.ARENA_LENGTH - 1
	}

	if rx == lib.ARENA_LENGTH {
		rx = 0
	}

	// build "lefthand" coordinate
	lp := make([]int, 0)
	lp = append(lp, lx)
	lp = append(lp, gy)

	// build "righthand" coordinate
	rp := make([]int, 0)
	rp = append(rp, rx)
	rp = append(rp, gy)

	// store coordinates
	o = append(o, lp)
	o = append(o, rp)

	// return result
	return o
}
