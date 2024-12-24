// package main

import (
	"bufio"
	"fmt"
	"os"
)

func add_memory(memory []int, val int, amount int) []int {
	for range amount {
		memory = append(memory, val)
	}
	return memory
}

func get_memory() []int {
	var memory []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	used := true
	id := 0

	for _, x := range line {
		amount := int(x - '0')
		if used {
			memory = add_memory(memory, id, amount)
			id += 1
		} else {
			memory = add_memory(memory, -1, amount)
		}
		used = !used
	}
	return memory
}

func get_next_position(memory []int, pos int) int {
	for memory[pos] == -1 && pos >= 0 {
		pos = pos - 1
	}
	return pos
}

func minimize(memory []int) {
	high := len(memory) - 1
	for i := range memory {
		if high <= i {
			break
		}
		if memory[i] == -1 {
			high = get_next_position(memory, high)
			if high == -1 {
				break
			}
			var new_slot MemoryValue
			if memory[i].length > memory[high].length {
				new_slot = MemoryValue{id: -1, length: memory[i].length - memory[high].length}
			}
			memory[i], memory[high] = memory[high], memory[i]
			if new_slot != nil {
				memory = slices.Insert(memory, i, new_slot)
			}
		}
	}
}

func checksum(memory []int) int {
	sum := 0
	end := get_next_position(memory, len(memory)-1) + 1
	for i := range end {
		sum += i * memory[i]
	}
	return sum
}

func main() {
	memory := get_memory()
	minimize(memory)
	checksum := checksum(memory)
	fmt.Print(checksum)
}
