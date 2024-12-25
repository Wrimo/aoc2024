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
	fmt.Println(len(region), string(node.value), len(region), calculate_perimeter(grid, region), region)
}

func calculate_perimeter(grid [][]node, region []index) int {
	perimeter := 0
	for _, position := range region {
		current_node, _ := get_from_grid(grid, position)
		for _, index := range calculate_offsets(position) {
			new_node, err := get_from_grid(grid, index)
			if err || new_node.value != current_node.value {
				perimeter += 1
			}
		}
	}
	return perimeter
}

func calculate_result(grid [][]node, regions [][]index) int {
	sum := 0
	for _, region := range regions {
		// print_region(region, grid)
		sum += len(region) * calculate_perimeter(grid, region)
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
