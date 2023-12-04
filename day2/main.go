package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	RED   = 0
	GREEN = 1
	BLUE  = 2
)

var possible = [3]int{12, 13, 14}

func colorInRange(color *string) bool {
	num := 0
	col := -1
	for _, c := range *color {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else {
			switch c {
			case 'b':
				col = BLUE
				break
			case 'g':
				col = GREEN
				break
			case 'r':
				col = RED
				break
			}
			if col != -1 {
				break
			}
		}
	}
	return possible[col] >= num
}
func turnIsGood(colors *[]string) bool {
	turnResult := true
	for _, color := range *colors {
		turnResult = turnResult && colorInRange(&color)
	}
	return turnResult
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	counter := 1
	for scanner.Scan() {
		game := scanner.Text()
		gameResult := true
		// search until finding ':'
		idx := 0
		for i, c := range game {
			if c == ':' {
				idx = i + 1 // "Game xx: "
				break
			}
		}
		turns := strings.Split(game[idx:], ";")
		for _, turn := range turns {
			colors := strings.Split(turn, ",")
			if gameResult && !turnIsGood(&colors) {
				gameResult = false
				break
			}
		}
		if gameResult {
			sum += counter
		} else {
		}
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

	sum := 0
	counter := 1
	for scanner.Scan() {
        redMax := 0
        blueMax := 0
        greenMax := 0
		game := scanner.Text()
		// search until finding ':'
		idx := 0
		for i, c := range game {
			if c == ':' {
				idx = i + 1 // "Game xx: "
				break
			}
		}
		turns := strings.Split(game[idx:], ";")
		for _, turn := range turns {
			colors := strings.Split(turn, ",")
			for _, color := range colors {
				num := 0
				col := -1
				for _, c := range color {
					if c >= '0' && c <= '9' {
						num = num*10 + int(c-'0')
					} else {
						switch c {
						case 'b':
							col = BLUE
                            blueMax = max(num, blueMax)
							break
						case 'g':
							col = GREEN
                            greenMax = max(num, greenMax)
							break
						case 'r':
							col = RED
                            redMax = max(num, redMax)
							break
						}
						if col != -1 {
							break
						}
					}
				}
			}
		}
        power := blueMax * greenMax * redMax
        sum += power
		counter++
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
