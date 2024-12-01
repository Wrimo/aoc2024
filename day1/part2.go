package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func count(list []int, val int) int {
    count := 0
    for i := range(len(list)) {
        if list[i] == val {
            count += 1
        }
    }

    return count
}

func main() {
    var first []int
    var second []int
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        results := strings.Split(line, "   ")
        num1, _ := strconv.Atoi(results[0])
        num2, _ := strconv.Atoi(results[1])
        first = append(first, num1)
        second = append(second, num2)
    }

    var total int 
    for i := range len(first) {
        total += first[i] * count(second, first[i]) 
    }

    fmt.Println(total)

}
