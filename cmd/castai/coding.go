package castai

import (
	"log"
	"math/rand/v2"
)

/*
A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random (usually 4 possibilities, except for fleas on the edge of the grid or at the corners).
What is the expected number of unoccupied squares after 50 rings of the bell?
Give your answer rounded to six decimal places.

n*n grid
r - bell rings

4 4

x x x x
x x x x
x x x x
x x x x
*/
func getDirection() int {
	return rand.IntN(3)
}

func jump(grid [][]int, gridLength int) {

	//  random 0-3
	/*
		0 left
		1 top
		2 right
		3 bottom

		rand.IntN(100)
	*/

	if grid[0][0] > 0 {
		topLeft := getDirection()
		if topLeft < 2 {
			grid[0][1] += 1
		} else {
			grid[1][0] += 1
		}
		grid[0][0] -= 1
	}

	if grid[0][gridLength-1] > 0 {
		topRight := getDirection()
		if topRight < 2 {
			grid[0][gridLength-2] += 1
		} else {
			grid[1][gridLength-1] += 1
		}
		grid[0][gridLength-1] -= 1
	}

	if grid[gridLength-1][0] > 0 {
		bottomleft := getDirection()
		if bottomleft < 2 {
			grid[gridLength-2][0] += 1
		} else {
			grid[gridLength-1][1] += 1
		}
		grid[gridLength-1][0] -= 1
	}

	if grid[gridLength-1][gridLength-1] > 0 {
		bottomRight := getDirection()
		if bottomRight < 2 {
			grid[gridLength-2][gridLength-1] += 1
		} else {
			grid[gridLength-1][gridLength-2] += 1
		}
		grid[gridLength-1][gridLength-1] -= 1
	}

	for i := 1; i < gridLength-2; i++ {
		for j := 1; j < gridLength-2; j++ {
			if grid[i][j] > 0 {
				direction := getDirection()
				if direction == 0 { // top
					grid[i-1][j] += 1
				} else if direction == 1 { // right
					grid[i][j+1] += 1
				} else if direction == 2 { // bottom
					grid[i+1][j] += 1
				} else { // left
					grid[i][j-1] += 1
				}
			}
		}
	}
}

func EulerFlee(gridLength int, bellRings int) float32 {
	var result float32
	var grid = make([][]int, gridLength)
	for i := range grid {
		grid[i] = make([]int, gridLength)
	}

	for i := 0; i < gridLength; i++ {
		for j := 0; j < gridLength; j++ {
			grid[i][j] = 1
		}
	}

	for i := 0; i < bellRings; i++ {
		jump(grid, gridLength)
	}
	log.Println(grid)
	return result
}
