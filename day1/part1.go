package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
    sort.Ints(first)
    sort.Ints(second)
    var total int 
    for i := range len(first) {
        diff := first[i] - second[i]
        if diff < 0 {
            diff = diff * -1
        }
        total += diff 
    }

    fmt.Println(total)

}
