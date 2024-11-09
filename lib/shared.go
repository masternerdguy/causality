package lib

const FLUX_BOUNDARY = 1
const FLUX_LOW = 0
const FLUX_HIGH = 1

var ARENA_LENGTH = -1
var ARENA_AREA = -1

var MAX_CYCLES = -1

func InitGlobals(al int) {
	ARENA_LENGTH = al
	ARENA_AREA = ARENA_LENGTH * ARENA_LENGTH
	MAX_CYCLES = ARENA_LENGTH
}
