package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type nums struct {
	n1 int
	n2 int
}

func (nums *nums) assignChar(c rune) {
	if nums.n1 == -1 {
		nums.n1 = int(c - '0')
	} else {
		nums.n2 = int(c - '0')
	}
}

func (nums *nums) assignNum(n int) {
	if nums.n1 == -1 {
		nums.n1 = n
	} else {
		nums.n2 = n
	}
}

func newNums() nums {
	return nums{n1: -1, n2: -1}
}

func sliceMatches(line *string, x int, y int, num string) bool {
	if len(*line) >= x+y {
		if (*line)[x:x+y] == num {
			return true
		}
	}
	return false
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	counter := 1
	sum := 0
	for scanner.Scan() {
		nums := newNums()
		line := scanner.Text()
		for _, c := range line {
			if c >= '0' && c <= '9' {
				nums.assignChar(c)
			} else {
				continue
			}
		}
		if nums.n2 == -1 {
			nums.n2 = nums.n1
		}
		val := nums.n1*10 + nums.n2
		sum += val
		counter++
	}
	return sum
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	counter := 1
	sum := 0
	for scanner.Scan() {
		nums := newNums()
		line := scanner.Text()
		for i, c := range line {
			if c >= '0' && c <= '9' {
				nums.assignChar(c)
			} else {
				switch c {
				case 'o':
					if sliceMatches(&line, i, 3, "one") {
						nums.assignNum(1)
					}
					break
				case 't':
					if sliceMatches(&line, i, 3, "two") {
						nums.assignNum(2)
					} else if sliceMatches(&line, i, 5, "three") {
						nums.assignNum(3)
					}
					break
				case 'f':
					if sliceMatches(&line, i, 4, "four") {
						nums.assignNum(4)
					} else if sliceMatches(&line, i, 4, "five") {
						nums.assignNum(5)
					}
					break
				case 's':
					if sliceMatches(&line, i, 3, "six") {
						nums.assignNum(6)
					} else if sliceMatches(&line, i, 5, "seven") {
						nums.assignNum(7)
					}
					break
				case 'e':
					if sliceMatches(&line, i, 5, "eight") {
						nums.assignNum(8)
					}
					break
				case 'n':
					if sliceMatches(&line, i, 4, "nine") {
						nums.assignNum(9)
					}
					break
				}
			}
		}
		if nums.n2 == -1 {
			nums.n2 = nums.n1
		}
		val := nums.n1*10 + nums.n2
		sum += val
		counter++
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
