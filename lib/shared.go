package lib

var ARENA_LENGTH = -1
var ARENA_AREA = -1
var MAX_CYCLES = -1
var SEED = -1

func InitGlobals(al int, s int) {
	ARENA_LENGTH = al
	MAX_CYCLES = al
	ARENA_AREA = ARENA_LENGTH * ARENA_LENGTH
	SEED = s
}
