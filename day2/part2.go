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

func diffCalculate(current int, next int) (bool, int) {
	diff, sign := abs(current - next)
	valid := diff > 0 && diff <= 3

	return valid, sign 
}

func confirmSign(vals []int) bool {
	sum := 0
	for _, x := range vals { 
		sum += x
	}
	sum, _ = abs(sum)
	
	return sum == len(vals)
}

func tryExcept(row []int, index int) bool {
	split1 = append(row[:index], row[index+1:]...)
	
	split2 := append([]int(nil), row...)
	split2 = append(split2[:index+1], split2[index+2:]...)

	fmt.Println(row, index, split1, split2)
	return verify(split1, false) || verify(split2, false)
}

func verify(row []int, allowReplacement bool) bool {
	var polarity []int
	worked := true
	for index, num := range row {
		if index == len(row) - 1 {
			continue
		}
		valid, sign := diffCalculate(num, row[index + 1])
		polarity = append(polarity, sign)
		valid = (valid && confirmSign(polarity))

		if allowReplacement && tryExcept(row, index) {
			worked = true 
			break
		}
		
		if !valid { 
			worked = false
			break
		}
	}
	fmt.Println(row, worked)
	return worked

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
		fmt.Println("row", row)
		valid := verify(row, true)
		if valid {
			result += 1
		}
	}

	fmt.Println(result)
}