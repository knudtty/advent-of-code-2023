package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getNum(line *string, idx int) (num int, nSize int) {
	val := 0
	count := 0
	length := len(*line)
	var c byte
	for idx < length {
		c = (*line)[idx]
		if c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
			count++
			idx++
		} else {
			break
		}
	}
	return val, count
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	prevLine := ""
	line := ""
	nextLine := ""

	counter := 0
	notLastLine := true
	for notLastLine {
		readIdx := 0
		if counter == 0 {
			notLastLine = scanner.Scan()
			line = scanner.Text()
			notLastLine = scanner.Scan()
			nextLine = scanner.Text()
		} else {
			line = nextLine
			notLastLine = scanner.Scan()
			nextLine = scanner.Text()
		}
		for readIdx < len(line) {
			c := line[readIdx]
			if c == '.' || c == '\n' {
				readIdx++
				continue
			} else if c >= '0' && c <= '9' {
				num, nSize := getNum(&line, readIdx)
				// check for symbols around
				var cur, end int
				if readIdx != 0 {
					cur = readIdx - 1
				} else {
					cur = 0
				}
				end = min(readIdx+nSize, len(line)-1)
				isPartNumber := false
				// check line above
				if counter > 0 {
					for !isPartNumber && cur <= end {
						c := prevLine[cur]
						if (c < '0' || c > '9') && c != '.' {
							isPartNumber = true
						}
						cur++
					}
				}
				// reset cur
				if readIdx != 0 {
					cur = readIdx - 1
				} else {
					cur = readIdx
				}
				// check this line
				for !isPartNumber && cur <= end {
					c := line[cur]
					if (c < '0' || c > '9') && c != '.' {
						isPartNumber = true
					}
					cur++
				}
				// reset cur
				if readIdx != 0 {
					cur = readIdx - 1
				} else {
					cur = readIdx
				}
				// check line below
				for !isPartNumber && cur <= end {
					if nextLine != "" {
						c := nextLine[cur]
						if (c < '0' || c > '9') && c != '.' {
							isPartNumber = true
						}
					} else {
						cur = end
					}
					cur++
				}
				if isPartNumber {
					sum += num
				}
				readIdx += nSize
			} else {
				readIdx++
			}
		}
		prevLine = line

		counter++
	}
	return sum
}

type pos struct {
	x int
	y int
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	prevLine := ""
	line := ""
	nextLine := ""

	counter := 0
	notLastLine := true
	gears := make(map[pos][]int)
	for notLastLine {
		readIdx := 0
		if counter == 0 {
			notLastLine = scanner.Scan()
			line = scanner.Text()
			notLastLine = scanner.Scan()
			nextLine = scanner.Text()
		} else {
			line = nextLine
			notLastLine = scanner.Scan()
			nextLine = scanner.Text()
		}
		for readIdx < len(line) {
			c := line[readIdx]
			if c >= '0' && c <= '9' {
				num, nSize := getNum(&line, readIdx)
				// check for symbols around
                cur := max(0, readIdx - 1)
                end := min(readIdx+nSize, len(line)-1)
				// check line above
				if counter > 0 {
					for cur <= end {
						c := prevLine[cur]
						if c == '*' {
							gears[pos{cur, counter - 1}] = append(gears[pos{cur, counter - 1}], num)
						}
						cur++
					}
				}
				// reset cur
				if readIdx != 0 {
					cur = readIdx - 1
				} else {
					cur = readIdx
				}
				// check this line
				for cur <= end {
					c := line[cur]
					if c == '*' {
						gears[pos{cur, counter}] = append(gears[pos{cur, counter}], num)
					}
					cur++
				}
				// reset cur
				if readIdx != 0 {
					cur = readIdx - 1
				} else {
					cur = readIdx
				}
				// check line below
				for cur <= end {
					if nextLine != "" {
						c := nextLine[cur]
						if c == '*' {
							gears[pos{cur, counter + 1}] = append(gears[pos{cur, counter + 1}], num)
						}
					} else {
						cur = end
					}
					cur++
				}
				readIdx += nSize
			} else {
				readIdx++
			}
		}
		prevLine = line

		counter++
	}
	sum := 0
    for _, gearSet := range gears {
        if len(gearSet) == 2 {
            sum += gearSet[0] * gearSet[1]
        }
    }
	return sum
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
