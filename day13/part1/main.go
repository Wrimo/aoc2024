package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Problem struct {
	abutton Position
	bbutton Position
	prize   Position
}

const MAX_INT = (1<<bits.UintSize)/2 - 1

func get_vector(line string, startIndex int) Position {
	segments := strings.Split(line, " ")

	x_txt := segments[startIndex]
	y_txt := segments[startIndex+1]

	x, _ := strconv.Atoi(x_txt[2 : len(x_txt)-1])
	y, _ := strconv.Atoi(y_txt[2:])
	return Position{x: x, y: y}
}

func get_problems() []Problem {
	var problems []Problem
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		abutton := get_vector(line, 2)

		scanner.Scan()
		line = scanner.Text()
		bbutton := get_vector(line, 2)

		scanner.Scan()
		line = scanner.Text()
		prize := get_vector(line, 1)

		scanner.Scan()
		scanner.Text()

		prob := Problem{
			abutton: abutton,
			bbutton: bbutton,
			prize:   prize,
		}
		problems = append(problems, prob)
	}

	return problems
}

func solve_all(problems []Problem) int {
	result := 0
	for _, prob := range problems {
		val := solve(prob)
		// fmt.Println(val)
		if val != MAX_INT {
			result += val
		}
	}
	return result
}

func get_from_grid(grid [][]int, pos Position) (int, bool) {
	if pos.y >= len(grid) || pos.y < 0 || pos.x >= len(grid[0]) || pos.x < 0 {
		return MAX_INT, true
	}
	return grid[pos.y][pos.x], false
}

func solve(prob Problem) int {
	grid := make([][]int, prob.prize.y+1)
	for i := range grid {
		grid[i] = make([]int, prob.prize.x+1)
	}

	for i := range grid {
		for j := range grid[0] {
			grid[i][j] = MAX_INT
		}
	}

	grid[prob.prize.y][prob.prize.x] = 0

	// fmt.Println(len(grid), len(grid[0]))
	for i := range grid {
		for j := range grid[0] {
			y := len(grid) - 1 - i
			x := len(grid[0]) - 1 - j

			aPosition := Position{x: x + prob.abutton.x, y: y + prob.abutton.y}
			bPosition := Position{x: x + prob.bbutton.x, y: y + prob.bbutton.y}

			aval, _ := get_from_grid(grid, aPosition)
			bval, _ := get_from_grid(grid, bPosition)

			acost := aval + 3
			bcost := bval + 1

			if (acost < bcost || bval == MAX_INT) && aval != MAX_INT {
				grid[y][x] = acost
			} else if bval != MAX_INT {
				grid[y][x] = bcost
			}

		}
	}

	return grid[0][0]
}

func main() {
	problems := get_problems()
	result := solve_all(problems)
	fmt.Println(result)
}
