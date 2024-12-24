package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type index struct {
	row   int
	col   int
	start *index
}

type FoundEnd struct {
	start index
	end   index
}

func get_grid() ([][]int, []index) {
	var grid [][]int
	var trailheads []index
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for j, x := range line {
			height := int(x - '0')
			if string(x) == "." {
				height = -9
			}

			if height == 0 {
				trailheads = append(trailheads, index{row: i, col: j})
			}
			row = append(row, height)
		}
		grid = append(grid, row)
		i += 1
	}

	for i := range trailheads {
		trailheads[i].start = &trailheads[i]
	}
	return grid, trailheads
}

func get_from_grid(grid [][]int, pos index) (int, bool) {
	if pos.row >= len(grid) || pos.row < 0 || pos.col >= len(grid[0]) || pos.col < 0 {
		return -1, true
	}
	return grid[pos.row][pos.col], false
}

func get_next_steps(grid [][]int, position index) ([]index, []index) {
	current_height, _ := get_from_grid(grid, position)
	var ends []index
	var positions []index

	indices := []index{
		index{row: position.row + 1, col: position.col}, // up
		index{row: position.row - 1, col: position.col}, // down
		index{row: position.row, col: position.col + 1}, // right
		index{row: position.row, col: position.col - 1}, // left
	}

	for _, ind := range indices {
		height, err := get_from_grid(grid, ind)
		if err {
			continue
		}
		ind.start = position.start
		if current_height+1 == height {
			if height == 9 {
				ends = append(ends, ind)
			} else {
				positions = append(positions, ind)
			}
		}
	}
	return positions, ends
}

func traverse(grid [][]int, trailheads []index) int {
	var highpoints []FoundEnd
	positions := trailheads
	for len(positions) != 0 {
		var new_positions []index
		for _, pos := range positions {
			new_pos, ends := get_next_steps(grid, pos)
			new_positions = append(new_positions, new_pos...)

			if len(ends) > 0 {
				for _, x := range ends {
					ending := FoundEnd{start: *x.start, end: x}
					if !slices.Contains(highpoints, ending) { // only add each start to end once
						highpoints = append(highpoints, ending)
					}
				}
			}
		}
		positions = new_positions
	}
	return len(highpoints)
}

func print_grid(grid [][]int) {
	for _, x := range grid {
		for _, y := range x {
			fmt.Print(y, " ")
		}
		fmt.Println()
	}
}

func main() {
	grid, heads := get_grid()
	result := traverse(grid, heads)
	fmt.Println(result)
}
