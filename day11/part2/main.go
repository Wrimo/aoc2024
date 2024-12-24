package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	value  int
	blinks int
}

func get_rocks() []int {
	var vals []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	for _, x := range line {
		num, _ := strconv.Atoi(x)
		vals = append(vals, num)
	}
	return vals
}

func split(val int) (int, int) {
	str := strconv.Itoa(val)
	x, _ := strconv.Atoi(str[0 : len(str)/2])
	y, _ := strconv.Atoi(str[len(str)/2:])
	return x, y
}

func blink(num int, blink_count int, max_blink int, cache map[Position]int) int {
	var new []int
	position := Position{value: num, blinks: blink_count}

	if val, ok := cache[position]; ok {
		return val
	}
	if num == 0 {
		new = append(new, 1)
	} else if len(strconv.Itoa(num))%2 == 0 {
		first, second := split(num)
		new = append(new, first, second)

	} else {
		new = append(new, num*2024)
	}

	if blink_count == max_blink {
		sum := len(new)
		cache[position] = sum
		return sum
	}

	sum := 0
	next_blink := blink_count + 1
	for _, x := range new {
		sum += blink(x, next_blink, max_blink, cache)
	}
	cache[position] = sum
	return sum
}

func perform_blinks(rocks []int, blink_count int) int {
	sum := 0
	cache := make(map[Position]int)
	for _, x := range rocks {
		sum += blink(x, 1, blink_count, cache)
	}
	return sum
}

func main() {
	rocks := get_rocks()
	// fmt.Println(rocks)
	result := perform_blinks(rocks, 75)
	fmt.Println(result)
}
