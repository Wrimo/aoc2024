package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func blink(rocks []int) []int {
	var new []int
	for _, x := range rocks {
		if x == 0 {
			new = append(new, 1)
		} else if len(strconv.Itoa(x))%2 == 0 {
			first, second := split(x)
			new = append(new, first, second)

		} else {
			new = append(new, x*2024)
		}
	}
	return new
}

func perform_blinks(rocks []int, blink_count int) int {
	for range blink_count {
		rocks = blink(rocks)
	}
	return len(rocks)
}

func main() {
	rocks := get_rocks()
	result := perform_blinks(rocks, 35)
	fmt.Println(result)
}
