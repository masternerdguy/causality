package arena

import (
	"flux/auto"
	"fmt"
)

var arena = make([][]auto.Cell, ARENA_LENGTH)
var arenaRender = make([][]auto.RenderCell, ARENA_LENGTH)

const ARENA_LENGTH = 25

/**/

func InitArena() {
	// loop over rows
	for x := 0; x < ARENA_LENGTH; x++ {
		// init cells
		arena[x] = make([]auto.Cell, ARENA_LENGTH)
		arenaRender[x] = make([]auto.RenderCell, ARENA_LENGTH)

		// store global coordinates for convenient inspection
		for y := 0; y < ARENA_LENGTH; y++ {
			// arena
			arena[x][y].G_X = x
			arena[x][y].G_Y = y

			// framebuffer
			arenaRender[x][y].G_X = x
			arenaRender[x][y].G_Y = y
			arenaRender[x][y].G_S = "*"
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
			fmt.Print(arenaRender[x][y].G_S)

			// spacer
			fmt.Print(" ")
		}

		// line break
		fmt.Println()
	}

	// indicate end of frame
	fmt.Println("\n")
}
