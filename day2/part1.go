package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(val int) (int, int) {
	if val < 0 {
		return val * -1, -1
	}
	return val, 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines [][]int
    for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")

		var row []int
		for _, element := range nums {
			val, _ := strconv.Atoi(element)
			row = append(row, val)
		}
		lines = append(lines, row)
	}

	result := 0
	for _, row := range lines {
		maxDiff := true
		polarity := 0
		for index := range len(row) {
			if index == 0 { 
				continue
			}

			diff, sign := abs(row[index] - row[index - 1])
			polarity += sign

			if diff == 0 || diff > 3 {
				maxDiff = false
				break
			}
		}
		polarity, _ = abs(polarity)
		if maxDiff && polarity == len(row) - 1 { 
			result += 1
		}
	}

	fmt.Println(result)
}