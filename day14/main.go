package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(scanner *bufio.Scanner) [][]byte {
	f := [][]byte{}
	for scanner.Scan() {
		f = append(f, []byte(scanner.Text()))
	}
	return f
}

func tiltNorth(f *[][]byte) {
	for y, row := range *f {
		for x, c := range row {
			if c == 'O' {
				idx := y - 1
				for ; idx >= 0 && (*f)[idx][x] == '.'; idx-- {
				}
				(*f)[y][x] = '.'
				(*f)[idx+1][x] = 'O'
			}
		}
	}
}

func tiltWest(f *[][]byte) {
	for y, row := range *f {
		for x, c := range row {
			if c == 'O' {
				idx := x - 1
				for ; idx >= 0 && (*f)[y][idx] == '.'; idx-- {
				}
				(*f)[y][x] = '.'
				(*f)[y][idx+1] = 'O'
			}
		}
	}
}

func tiltSouth(f *[][]byte) {
	for y := len(*f) - 1; y >= 0; y-- {
		for x := len((*f)[y]) - 1; x >= 0; x-- {
			c := (*f)[y][x]
			if c == 'O' {
				idx := y + 1
				for ; idx < len(*f) && (*f)[idx][x] == '.'; idx++ {
				}
				(*f)[y][x] = '.'
				(*f)[idx-1][x] = 'O'
			}
		}
	}
}

func tiltEast(f *[][]byte) {
	for y := len(*f) - 1; y >= 0; y-- {
		for x := len((*f)[y]) - 1; x >= 0; x-- {
			c := (*f)[y][x]
			if c == 'O' {
				idx := x + 1
				for ; idx < len((*f)[y]) && (*f)[y][idx] == '.'; idx++ {
				}
				(*f)[y][x] = '.'
				(*f)[y][idx-1] = 'O'
			}
		}
	}
}
func calcLoad(f [][]byte) int {
	sum := 0
	for y, row := range f {
		for _, c := range row {
			if c == 'O' {
				sum += (len(f) - y)
			}
		}
	}
	return sum
}

func part1() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	field := parse(bufio.NewScanner(file))
	for _, row := range field {
		fmt.Println(string(row))
	}
	tiltNorth(&field)
	fmt.Println()
	for _, row := range field {
		fmt.Println(string(row))
	}

	return calcLoad(field)
}
func fieldsEqual(f1 *[][]byte, f2 *[][]byte) bool {
	for y, row := range *f1 {
		for x := range row {
			if (*f1)[y][x] != (*f2)[y][x] {
				return false
			}
		}
	}
	return true
}

func deepCopy(f [][]byte) [][]byte {
	out := [][]byte{}
	for y, row := range f {
		out = append(out, []byte{})
		for _, c := range row {
			out[y] = append(out[y], c)
		}
	}
	return out
}

func part2() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	field := parse(bufio.NewScanner(file))
	// This actually repeats in a cycle every so often, and the value at 1000 iterations will be the same at 1000000000
	for i := 0; i < 1000; i++ {
		tiltNorth(&field)
		tiltWest(&field)
		tiltSouth(&field)
		tiltEast(&field)
	}

	return calcLoad(field)
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
