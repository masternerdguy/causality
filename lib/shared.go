package lib

const FLUX_BOUNDARY = 1
const FLUX_LOW = 0
const FLUX_HIGH = 1

var ARENA_LENGTH = -1
var ARENA_AREA = -1

var MAX_CYCLES = -1

var ELAPSED_SEED = -1
var FLUX_SEED = -1

func InitGlobals(al int, es int, fs int) {
	ARENA_LENGTH = al
	MAX_CYCLES = al
	ARENA_AREA = ARENA_LENGTH * ARENA_LENGTH
	ELAPSED_SEED = es
	FLUX_SEED = fs
}
