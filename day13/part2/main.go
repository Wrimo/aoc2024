package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"slices"
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
		// prize.x += 10000000000000
		// prize.y += 10000000000000

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
		fmt.Println(val)
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

func get_from_map(values map[Position]int, pos Position) int {
	val, ok := values[pos]
	if !ok {
		return MAX_INT
	}
	return val
}

func pop(items []Position) (Position, []Position) {
	return items[0], items[1:]
}

func add(items []Position, pos Position, prob Problem, costs map[Position]int) []Position {
	_, ok := costs[pos]
	if pos.x > prob.prize.x || pos.y > prob.prize.y || pos.x < 0 || pos.y < 0 || ok || slices.Contains(items, pos) {
		return items
	}
	items = append(items, pos)
	return items
}

func magnitude(vec Position) float64 {
	return math.Sqrt(math.Pow(float64(vec.x), 2) + math.Pow(float64(vec.y), 2))
}

func solve(prob Problem) int {
	costs := make(map[Position]int)

	amag, bmag := magnitude(prob.abutton), magnitude(prob.bbutton)
	var times int
	var new_start Position
	if amag > bmag {
		times = min(prob.prize.x/prob.abutton.x, prob.prize.y/prob.abutton.y)
		new_start = Position{x: prob.abutton.x * times, y: prob.abutton.y * times}
		costs[new_start] = times * 3
	} else {
		times = min(prob.prize.x/prob.bbutton.x, prob.prize.y/prob.bbutton.y)
		new_start = Position{x: prob.bbutton.x * times, y: prob.bbutton.y * times}
		costs[new_start] = times * 1
	}

	costs[prob.prize] = 0
	// fmt.Println(costs)

	var work []Position
	var pos Position

	work = add(work, Position{x: prob.prize.x - prob.abutton.x, y: prob.prize.y - prob.abutton.y}, prob, costs)
	work = add(work, Position{x: prob.prize.x - prob.bbutton.x, y: prob.prize.y - prob.bbutton.y}, prob, costs)

	for len(work) > 0 {
		pos, work = pop(work)
		x, y := pos.x, pos.y
		aPosition := Position{x: x + prob.abutton.x, y: y + prob.abutton.y}
		bPosition := Position{x: x + prob.bbutton.x, y: y + prob.bbutton.y}

		aval := get_from_map(costs, aPosition)
		bval := get_from_map(costs, bPosition)

		acost := aval + 3
		bcost := bval + 1
		if (acost < bcost || bval == MAX_INT) && aval != MAX_INT {
			costs[pos] = acost
		} else if bval != MAX_INT {
			costs[pos] = bcost
		}

		next_position1 := Position{x: x - prob.abutton.x, y: y - prob.abutton.y}
		next_position2 := Position{x: x - prob.bbutton.x, y: y - prob.bbutton.y}

		work = add(work, next_position1, prob, costs)
		work = add(work, next_position2, prob, costs)
	}

	return costs[Position{x: 0, y: 0}]
}

func main() {
	problems := get_problems()
	result := solve_all(problems)
	fmt.Println(result)
}
