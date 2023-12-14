package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(scanner *bufio.Scanner) [][]string {
	f := [][]string{}
	f = append(f, []string{})
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			f[len(f)-1] = append(f[len(f)-1], line)
		} else {
			f = append(f, []string{})
		}
	}
	return f
}

func reflectIdx(f *[]string) int {
	for split := 1; split < len(*f); split++ {
		size := min(split, len(*f)-split)
		reflection := true
		for i := 0; i < size && reflection; i++ {
			if (*f)[split-i-1] != (*f)[split+i] {
				reflection = false
				break
			}
		}
		if reflection {
			return split
		}
	}
	return 0
}
func getCols(f *[]string) []string {
	cols := []string{}
	for i := 0; i < len((*f)[0]); i++ {
		col := ""
		for _, row := range *f {
			col += row[i : i+1]
		}
		cols = append(cols, col)
	}
	return cols
}

func part1() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fields := parse(bufio.NewScanner(file))
	sum := 0
	for _, field := range fields {
		row := reflectIdx(&field)
		cols := getCols(&field)
		col := reflectIdx(&cols)
		sum += col + 100*row
	}

	return sum
}

func offByOne(a string, b string) bool {
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count == 1
}
func reflectSmudgeIdx(f *[]string) int {
	for split := 1; split < len(*f); split++ {
		size := min(split, len(*f)-split)
		reflection := true
		smudgeApplied := false
		for i := 0; i < size && reflection; i++ {
			a := (*f)[split-i-1]
			b := (*f)[split+i]
			if a != b {
				if !smudgeApplied && offByOne(a, b) {
					smudgeApplied = true
				} else {
					reflection = false
                    break
				}
			}
		}
		if reflection && smudgeApplied {
			return split
		}
	}
	return 0
}

func part2() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fields := parse(bufio.NewScanner(file))
	sum := 0
	for _, field := range fields {
		row := reflectSmudgeIdx(&field)
		cols := getCols(&field)
		col := reflectSmudgeIdx(&cols)
		sum += col + 100*row
	}

	return sum
}
func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
