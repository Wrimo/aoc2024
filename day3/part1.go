package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Operation struct { 
	op int
	lhand int
	rhand int
}

func get_num(txt string, start int) (int, bool, int) {
	fmt.Println("get_num")
	digits := ""
	var i int
	for i = start; i < len(txt); i++ {
		if !unicode.IsDigit(rune(txt[i])) {
			break
		}
		digits = digits + string(txt[i])
	}

	fmt.Println("started at", string(txt[start]), "digits", digits)

	if len(digits) > 3 || len(digits) == 0 {
		return 0, false, i
	}
	x, _ := strconv.Atoi(digits)
	return x, true, i
}

func get_mul(txt string, start int) int {
	for i := start; i < len(txt) - 2; i++ {
		fmt.Println(i)
		if txt[i:i+3] == "mul" {
			return i+3
		}
	}
	return len(txt)
}

func get_arguments(txt string, start int) (bool, int, int, int) {
	index := start
	if index > len(txt) - 5 { 
		return false, len(txt), 0, 0
	}

	if txt[index] != '(' {
		return false, index, 0, 0
	}
	index++ 
	
	l, worked, i := get_num(txt, index)
	fmt.Println("l", l, worked)
	if !worked {
		return false, index, 0, 0
	}
	index = i

	if txt[index] != ',' {
		return false, index, 0, 0
	}
	index++

	r, worked, i  := get_num(txt, index)
	fmt.Println("r", r, worked)
	if !worked {
		return false, index, 0, 0
	}
	index = i
	
	if txt[index] != ')' {
		return false, index, 0, 0
	}
		
	return true, index, l, r
}

func parse(txt string) []Operation {
	var ops []Operation
	for i := 0; i < len(txt); i++ {
		i = get_mul(txt, i)
		
		worked, index, l, r := get_arguments(txt, i)
		fmt.Println(worked, index, l, r)
		if worked {
			var op Operation 
			op.lhand = l
			op.rhand = r
			ops = append(ops, op)
			fmt.Println(ops)
		}
		i = index
	}
	return ops
}

func perform_operations(ops []Operation) int {
	total := 0
	for _, i := range ops { 
		total += i.rhand * i.lhand
	}
	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operations := parse(scanner.Text())
	fmt.Println(perform_operations(operations))
}
