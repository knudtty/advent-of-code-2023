package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type OrdPair struct {
	a int
	b int
}

func NewOrdPair(a int, b int) OrdPair {
	return OrdPair{
		min(a, b),
		max(a, b),
	}
}

type Pos struct {
	x int
	y int
}

var (
	EMPTY = '.'
	FULL  = '#'
)

type Distances map[OrdPair]int
type Galaxy struct {
	space     []string
	rowGap    map[int]bool
	colGap    map[int]bool
	dist      Distances
	locations []Pos
	gapSize   int
}

func NewGalaxy(space []string, gapSize int) Galaxy {
	g := Galaxy{
		space,
		map[int]bool{},
		map[int]bool{},
		Distances{},
		[]Pos{},
		gapSize,
	}
	return g
}

func (g *Galaxy) push(s string) {
	g.space = append(g.space, s)
}
func (g *Galaxy) expand() {
	occupiedRows := map[int]bool{}
	occupiedCols := map[int]bool{}
	for y, row := range g.space {
		for x, c := range row {
			if c == FULL {
				occupiedRows[y] = true
				occupiedCols[x] = true
			}
		}
	}
	for y, row := range g.space {
		for x := range row {
			if !occupiedCols[x] {
				g.colGap[x] = true
			}
		}
		if !occupiedRows[y] {
			g.rowGap[y] = true
		}
	}
}
func (g *Galaxy) getRowGap(y1 int, y2 int) int {
	a := min(y1, y2)
	b := max(y1, y2)
	gap := b - a
	for i := a; i < b; i++ {
		if g.rowGap[i] {
			gap += g.gapSize
		}
	}
	return gap
}
func (g *Galaxy) getColGap(x1 int, x2 int) int {
	a := min(x1, x2)
	b := max(x1, x2)
	gap := b - a
	for i := a; i < b; i++ {
		if g.colGap[i] {
			gap += g.gapSize
		}
	}
	return gap
}
func (g *Galaxy) fillLocations() {
	for y, row := range g.space {
		for x, c := range row {
			if c == FULL {
				g.locations = append(g.locations, Pos{x, y})
			}
		}
	}
}
func (g *Galaxy) fillDistances() {
	if len(g.locations) == 0 {
		g.fillLocations()
	}
	for i, pos1 := range g.locations {
		for j, pos2 := range g.locations {
			if pos1 == pos2 {
				continue
			}
			ordPair := NewOrdPair(i, j)
			g.dist[ordPair] = g.getRowGap(pos1.y, pos2.y) + g.getColGap(pos1.x, pos2.x)
		}
	}
}
func (g *Galaxy) print() {
	for y, row := range g.space {
		for x, c := range row {
            if g.colGap[x] {
                for i := 0; i < g.gapSize; i++ {
			        fmt.Printf("%c", EMPTY)
                }
            }
			fmt.Printf("%c", c)
		}
        if g.rowGap[y] {
            fmt.Println()
            for i := 0; i < (len(row) + len(g.rowGap) * g.gapSize + 1); i++ {
                fmt.Printf("%c", EMPTY)
            }
        }
		fmt.Println()
	}
}
func (g *Galaxy) distanceSums() int {
	sum := 0
	for _, v := range g.dist {
		sum += v
	}
	return sum
}

func parse(scanner *bufio.Scanner, gapSize int) Galaxy {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return NewGalaxy(lines, gapSize)
}

func part1() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gal := parse(bufio.NewScanner(file), 1) // field
	gal.expand()
	gal.fillDistances()

	return gal.distanceSums()
}

func part2() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	gal := parse(bufio.NewScanner(file), 999_999) // field
	gal.expand()
	gal.fillDistances()

	return gal.distanceSums()
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
