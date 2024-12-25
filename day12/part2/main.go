package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type index struct {
	row int
	col int
}

type direction struct {
	old int
	new int
}

type node struct {
	value   rune
	visited bool
}

func get_from_grid(grid [][]node, pos index) (*node, bool) {
	if pos.row >= len(grid) || pos.row < 0 || pos.col >= len(grid[0]) || pos.col < 0 {
		return nil, true
	}
	return &grid[pos.row][pos.col], false
}

func calculate_offsets(position index) []index {
	return []index{
		index{row: position.row + 1, col: position.col}, // up
		index{row: position.row - 1, col: position.col}, // down
		index{row: position.row, col: position.col + 1}, // right
		index{row: position.row, col: position.col - 1}, // left
	}
}

func pop(items []index) (index, []index) {
	return items[len(items)-1], items[:len(items)-1]
}

func get_grid() [][]node {
	var grid [][]node
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var row []node
		for _, x := range line {
			n := node{value: x, visited: false}
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}

func get_regions(grid [][]node) [][]index {
	var regions [][]index
	for i, y := range grid {
		for j, node := range y {
			if node.visited {
				continue
			}
			regions = append(regions, floodfill(grid, index{row: i, col: j}))
		}
	}
	return regions
}

func floodfill(grid [][]node, position index) []index {
	var nodes []index
	var next []index

	start_node, _ := get_from_grid(grid, position)
	start_node.visited = true
	symbol := start_node.value
	nodes = append(nodes, position)
	next = append(next, position)
	var current index
	for len(next) > 0 {
		current, next = pop(next)
		for _, index := range calculate_offsets(current) {
			new_node, err := get_from_grid(grid, index)
			if err || new_node.value != symbol || new_node.visited {
				continue
			}
			new_node.visited = true
			nodes = append(nodes, index)
			next = append(next, index)
		}

	}

	return nodes
}

func print_region(region []index, grid [][]node) {
	node, _ := get_from_grid(grid, region[0])
	fmt.Println(len(region), string(node.value), "sides:", calculate_sides(grid, region), region)
}

func last_item_equal(list []direction, val direction) bool {
	if len(list) == 0 {
		return false
	}
	return list[len(list)-1] == val
}

func count_nonnegative(list []direction) int {
	num := 0
	for _, x := range list {
		if x.new == -1 && x.old == -1 {
			continue
		}
		num += 1
	}
	return num
}

func calculate_sides(grid [][]node, region []index) int {
	var horizontal_sides []direction
	var vertical_sides []direction
	// vertical down
	for i := 0; i < len(grid)+1; i++ {
		for j := 0; j < len(grid[0])+1; j++ {
			position := index{row: i, col: j}
			past_position := index{row: i - 1, col: j}
			contains_old := slices.Contains(region, past_position)
			contains_new := slices.Contains(region, position)

			var direct direction
			if contains_new {
				direct = direction{old: i - 1, new: i}
			} else {
				direct = direction{old: i, new: i - 1}
			}
			if contains_old != contains_new {
				if !last_item_equal(horizontal_sides, direct) {
					horizontal_sides = append(horizontal_sides, direct)
				}
			} else {
				horizontal_sides = append(horizontal_sides, direction{old: -1, new: -1})
			}
			// fmt.Println(position, horizontal_sides)
		}

	}

	for j := 0; j < len(grid[0])+1; j++ {
		for i := 0; i < len(grid)+1; i++ {
			position := index{row: i, col: j}
			past_horizontal := index{row: i, col: j - 1}

			contains_old := slices.Contains(region, past_horizontal)
			contains_new := slices.Contains(region, position)

			var direct direction
			if contains_new {
				direct = direction{old: j - 1, new: j}
			} else {
				direct = direction{old: j, new: j - 1}
			}

			if contains_old != contains_new {
				if !last_item_equal(vertical_sides, direct) {
					vertical_sides = append(vertical_sides, direct)
				}
			} else {
				vertical_sides = append(vertical_sides, direction{old: -1, new: -1})
			}
		}

	}

	// fmt.Println(horizontal_sides)
	horiz := count_nonnegative(horizontal_sides)
	vert := count_nonnegative(vertical_sides)
	return horiz + vert
}

func calculate_result(grid [][]node, regions [][]index) int {
	sum := 0
	for _, region := range regions {
		// print_region(region, grid)
		sum += len(region) * calculate_sides(grid, region)
	}
	return sum
}
func main() {
	grid := get_grid()
	regions := get_regions(grid)
	result := calculate_result(grid, regions)
	fmt.Println(result)
	// area is just length of regions list
	// perimeter is the sum of the number of nodes that are of a different type at each node

}
