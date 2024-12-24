package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type MemoryValue struct {
	id     int
	length int
}

func get_memory() []MemoryValue {
	var memory []MemoryValue
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	used := true
	id := 0

	for _, x := range line {
		len := int(x - '0')
		var mem MemoryValue
		if used {
			mem = MemoryValue{id: id, length: len}
			id += 1
		} else {
			mem = MemoryValue{id: -1, length: len}
		}
		memory = append(memory, mem)
		used = !used
	}
	return memory
}

func get_first_open(memory []MemoryValue, length int) (int, int) {
	for index := range memory {
		val := memory[index]
		if val.id == -1 && val.length >= length {
			return index, val.length - length
		}
	}
	return -1, 0
}
func collapse(memory []MemoryValue) []MemoryValue {
	var val MemoryValue
	var new_list []MemoryValue
	for _, x := range memory {
		if x.id == val.id {
			val.length += val.length
		} else {
			new_list = append(new_list, val)
			val = x
		}
	}
	return new_list
}

func minimize(memory []MemoryValue) []MemoryValue {
	for i := range memory {
		k := len(memory) - 1 - i

		val := memory[k]
		if val.id == -1 {
			continue
		}

		index, diff := get_first_open(memory, val.length)

		if index == -1 || k <= index {
			continue
		}
		memory[index].length = val.length
		memory[k], memory[index] = memory[index], val
		// memory[index].moved = true
		if diff != 0 {
			memory = slices.Insert(memory, index+1, MemoryValue{id: -1, length: diff})
		}
	}
	return memory
}

func add_memory(memory []int, val int, amount int) []int {
	for range amount {
		memory = append(memory, val)
	}
	return memory
}

func to_int(memory []MemoryValue) []int {
	var new []int
	for _, x := range memory {
		new = add_memory(new, x.id, x.length)
	}
	return new
}

func checksum(memory []int) int {
	sum := 0
	for i := range memory {
		if memory[i] == -1 {
			continue
		}
		sum += i * memory[i]
	}
	return sum
}

func main() {
	memory := get_memory()
	memory = minimize(memory)
	checksum := checksum(to_int(memory))
	fmt.Print(checksum)
}
