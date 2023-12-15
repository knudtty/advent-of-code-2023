package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func hash(seed string) int {
	cur := 0
	for _, c := range []byte(seed) {
		cur += int(c)
		cur *= 17
		cur %= 256
	}
	return cur
}
func part1() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	seeds := parse(bufio.NewScanner(file))
	sum := 0
	for _, seed := range seeds {
		sum += hash(seed)
	}
	return sum
}

type Contents struct {
	label string
	val   int
}

func part2() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	boxes := [256][]Contents{}
	seeds := parse(bufio.NewScanner(file))
	for _, seed := range seeds {
		if seed[len(seed)-1] == '-' {
			label := strings.ReplaceAll(seed, "-", "")
			hashed := hash(label)
			for i, contents := range boxes[hashed] {
				if contents.label == label {
					boxes[hashed] = append(boxes[hashed][:i], boxes[hashed][i+1:]...)
				}
			}
		} else {
			spl := strings.Split(seed, "=")
			num, _ := strconv.Atoi(spl[1])
			c := Contents{spl[0], num}
			hashed := hash(c.label)
            prevExists := false
			for i, elem := range boxes[hashed] {
				if elem.label == c.label {
					boxes[hashed][i] = c
                    prevExists = true
                    break
				}
			}
            if !prevExists {
                boxes[hashed] = append(boxes[hashed], c)
            }
		}
	}
	focusingPower := 0
	for i, box := range boxes {
		for j, c := range box {
			power := (i + 1) * (j + 1) * c.val
			focusingPower += power
		}
	}
	return focusingPower
}
func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
