package main

import (
	"bufio"
	"fmt"
	"os"
)

type index struct {
	row int
	col int
}

type position struct {
	pos1 index
	pos2 index
	pos3 index
	pos4 index
}

func get_indices() []position {
	return []position{
		position{ // up
			pos1: index{row: 0, col: 0},
			pos2: index{row: 1, col: 0},
			pos3: index{row: 2, col: 0},
			pos4: index{row: 3, col: 0},
		},
		position{ // up right
			pos1: index{row: 0, col: 0},
			pos2: index{row: 1, col: 1},
			pos3: index{row: 2, col: 2},
			pos4: index{row: 3, col: 3},
		},
		position{ // right
			pos1: index{row: 0, col: 0},
			pos2: index{row: 0, col: 1},
			pos3: index{row: 0, col: 2},
			pos4: index{row: 0, col: 3},
		},
		position{ // down right
			pos1: index{row: 0, col: 0},
			pos2: index{row: -1, col: 1},
			pos3: index{row: -2, col: 2},
			pos4: index{row: -3, col: 3},
		},
		position{ // down
			pos1: index{row: 0, col: 0},
			pos2: index{row: -1, col: 0},
			pos3: index{row: -2, col: 0},
			pos4: index{row: -3, col: 0},
		},
		position{ // down left
			pos1: index{row: 0, col: 0},
			pos2: index{row: -1, col: -1},
			pos3: index{row: -2, col: -2},
			pos4: index{row: -3, col: -3},
		},
		position{ // left
			pos1: index{row: 0, col: 0},
			pos2: index{row: 0, col: -1},
			pos3: index{row: 0, col: -2},
			pos4: index{row: 0, col: -3},
		},
		position{ // up left
			pos1: index{row: 0, col: 0},
			pos2: index{row: 1, col: -1},
			pos3: index{row: 2, col: -2},
			pos4: index{row: 3, col: -3},
		},
	}
}

func get_puzzle() []string {
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
	return grid
}

func get_from_grid(grid []string, pos index) (string, bool) {
	if pos.row >= len(grid) || pos.row < 0 || pos.col >= len(grid[0]) || pos.col < 0 {
		return "", true
	}
	return string(grid[pos.row][pos.col]), false
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func scan(row int, col int, grid []string, positions []position) int {
	// list of all eight strings from from x
	var words []string

	for _, pos := range positions {
		ch1, check1 := get_from_grid(grid, index{row: row + pos.pos1.row, col: col + pos.pos1.col})
		ch2, check2 := get_from_grid(grid, index{row: row + pos.pos2.row, col: col + pos.pos2.col})
		ch3, check3 := get_from_grid(grid, index{row: row + pos.pos3.row, col: col + pos.pos3.col})
		ch4, check4 := get_from_grid(grid, index{row: row + pos.pos4.row, col: col + pos.pos4.col})

		if !check1 && !check2 && !check3 && !check4 {
			word := ch1 + ch2 + ch3 + ch4
			words = append(words, word)
		}
	}

	count := 0
	for _, x := range words {
		if x == "XMAS" || reverse(x) == "XMAS" {
			count++
		}
	}
	return count
}

func main() {
	grid := get_puzzle()
	indices := get_indices()
	count := 0
	for row, str := range grid {
		for col, chr := range str {
			if chr == 'X' {
				count += scan(row, col, grid, indices)
			}
		}
	}

	fmt.Println(count)

}
