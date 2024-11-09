# Introduction
`causality` is an esoteric programming language ("esolang", as the kids say) inspired by causal set theory and Conway's Game of Life. This strange bird began life out of a desire to do some crackpot simulations regarding the nature of the universe.

# Execution Model
## Program Format
In `causality`, a program is a plain text file containing a square matrix of numbers. For example, a simple 3x3 program would be

    0 0 0
    0 0 0
    0 0 0

The numbers provided are interpreted as the initial value of each cell's internal signal counter. By varying these, control flow may be achieved.

Running this program will return the following final state:

	* * *
	*   *
	* * *

## Execution Process
Program execution began at coordinate (1,1). An `*` represents a "live" cell, which may be interpreted as a `1`, and ` ` represents a "dead" cell, which may be interpreted as a `0`. Cells that never participated in program execution are marked with `_` for convenient identification.

Unlike Conway's Game of Life, in which the future state of a cell is determined by its immediate neighbours, `causality` is based on the idea that data flows from the causal past to the causal future. Each cell passes a signal (either a `1` or a `-1`) to its "lefthand" and "righthand" descendants. **This can be visualized as a light cone with only its edges.** A cell that is "alive" will pass `1`, while a cell that is dead will pass `-1`.

Upon receiving a signal from the causal past, a cell will accumulate the value into its internal state (this is just addition) - which will change the cell's "live" / "dead" state. Each cell also has an additional internal counter that tracks the number of signals received, which is also incremented. Finally, the cell will signal its descendants. On the screen, propagation is right to left.

**When a cell's internal signal counter reaches the length of the matrix, it will cease to propagate signals. This means that cells that were initialized with higher signal counters will participate in fewer program cycles.** Note that cells which are initialized with a signal counter at or above the length of the matrix will never participate. In the case of a 3x3 program, this value is `3`.

Execution is started by a special `0` signal injected into the center. **The cell set is toroidal and signals wrap at program "edges".** When no cell changes have occurred for a while, the program is assumed to have "halted".

## Included Program
`causality` ships with a larger program included for reference:

    1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1
	0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 2
	0 0 1 0 0 0 0 0 4 0 0 0 0 1 0 0 0 0 0 4 0 0 3
	0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 4
	0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 5
	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 6
	0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 7
	0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 8
	0 0 4 0 0 0 0 0 1 0 0 0 0 4 0 0 0 0 0 1 0 0 9
	0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 1
	0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 2
	1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 3
	0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 4
	0 0 1 0 0 0 0 0 4 0 0 0 0 1 0 0 0 0 0 4 0 0 5
	0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 6
	0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 7
	0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 8
	0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 9
	0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 0 0 1
	0 0 4 0 0 0 0 0 1 0 0 0 0 4 0 0 0 0 0 1 0 0 2
	0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 0 3
	0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 1 4
	1 2 3 4 5 6 7 8 9 1 2 3 4 5 6 7 8 9 1 2 3 4 5

This program may be run via `go run main.go program.txt`.

If you actually precompiled this, you can use `causality program.txt`.

The program will return the following final state:

      * * * * * * * * * *   * * * * * * * * * *   
	*   * * * * * * * * * *   * * * * * * * * * * 
	* *   * * * * * * * * * *   * * * * * * * *   
	* * *   * * * * * * * * * *   * * * * * * * * 
	* * * *   * * * * * * * * * *   * * * * * *   
	* * * * * * * * * * * * * * * * * * * * * * * 
	* * * * * *   * * * * * * * * * *   * * * *   
	* * * * * * *   * * * * * * * * * *   * * * * 
	* * * * * * * *   * * * * * * * * * *   * *   
	* * * * * * * * *   * * * * * * * * * *   *   
	* * * * * * * * * *   * * * * * * * * * *   * 
	  * * * * * * * * * * * * * * * * * * * * *   
	*   * * * * * * * * * *   * * * * * * * * * * 
	* *   * * * * * * * * * *   * * * * * * * *   
	* * *   * * * * * * * * * *   * * * * * * * * 
	* * * *   * * * * * * * * * *   * * * * * *   
	* * * * * * * * * * * * * * * * * * * * * * * 
	* * * * * *   * * * * * * * * * *   * * * *   
	* * * * * * *   * * * * * * * * * *   * * *   
	* * * * * * * *   * * * * * * * * * *   * * * 
	* * * * * * * * *   * * * * * * * * * *   *   
	* * * * * * * * * *   * * * * * * * * * *   * 
	  *   *   *   *     *   *   *   *     *   *  

Execution isn't exactly fast - on a high-end CPU from ~2012 `causality` takes ~5 hours to execute the worst-case 100x100 program. However, programs below ~30x30 execute significantly faster. Also, you won't get much speedup on newer hardware from what I've seen.

## Best Practices
To maximize your productivity with `causality`, I recommend the following:

* For maximum self-interactivity, use odd-length matrices (1x1, 3x3, 5x5, etc).
* For minimal self-interactivity, use even-length matrices (2x2, 4x4, 6x6, etc).
* Negative numbers may be used to give cells "extra life".
* "Barriers" may be created by initializing a cell with an initial signal counter value at or above the length of the matrix. This can be used to create causally disconnected regions within your program that are unreachable.
* Similarly, an initial signal value just below the threshold may be used to execute certain regions "only once".
