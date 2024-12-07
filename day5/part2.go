package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func init_map(scanner *bufio.Scanner) map[int][]int {
	var precedence map[int][]int = make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		nums := strings.Split(line, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])

		precedence[second] = append(precedence[second], first)
	}
	return precedence
}

func get_print_jobs(scanner *bufio.Scanner) [][]int {
	var print_jobs [][]int
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		var job []int
		for _, x := range nums {
			val, _ := strconv.Atoi(x)
			job = append(job, val)
		}
		print_jobs = append(print_jobs, job)
	}
	return print_jobs
}

func validate_jobs(precedence map[int][]int, jobs [][]int) int {
	total := 0
	for _, job := range jobs {
		var seen []int
		failed := false
		for _, val := range job {

			// check prior has been seen
			var required []int = precedence[val]
			for _, pred := range required {
				if !slices.Contains(seen, pred) && slices.Contains(job, pred) {
					failed = true

					// go through list and move any require items before this one
					// keep going till its in the rgith order
					break
				}
			}

			if failed {
				break
			}

			// add myself
			seen = append(seen, val)
		}

		if failed {
			sort.Slice(job, func(i, j int) bool {
				i, j = job[i], job[j]
				return slices.Contains(precedence[j], i)
			})
			total += job[len(job)/2]
		}
	}
	return total
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	precedence := init_map(scanner)
	jobs := get_print_jobs(scanner)
	val := validate_jobs(precedence, jobs)

	fmt.Println(val)
}
