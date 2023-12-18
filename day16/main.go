package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func parse(scanner *bufio.Scanner) BeamField {
	bf := BeamField{}
	for scanner.Scan() {
		bf.f = append(bf.f, []byte(scanner.Text()))
	}
	return bf
}
type BeamField struct {
	f [][]byte
	m BeamMap
}

func (b BeamField) inBounds(p Vec2) bool {
	return p.y >= 0 && p.y < len(b.f) && p.x >= 0 && p.x < len(b.f[p.y])
}
func (b BeamField) getByte(p Vec2) (byte, error) {
	if b.inBounds(p) {
		return b.f[p.y][p.x], nil
	}
	return 0, errors.New("Outside of field")
}

type Vec2 struct {
	x int
	y int
}

var (
	NORTH Vec2 = Vec2{0, -1}
	SOUTH Vec2 = Vec2{0, 1}
	EAST  Vec2 = Vec2{1, 0}
	WEST  Vec2 = Vec2{-1, 0}
)

type Beam struct {
	pos Vec2
	dir Vec2
}
type BeamMap map[Beam]bool

var beamField *BeamField

func (b *Beam) next() ([]Beam, error) {
	newBeams := []Beam{}
	c, err := beamField.getByte(b.pos)
	if err != nil {
		return newBeams, err
	}
	switch c {
	case '.':
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, Vec2{b.dir.x, b.dir.y}})
		break
	case '-':
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, Vec2{1, 0}})
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, Vec2{-1, 0}})
		break
	case '|':
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, Vec2{0, 1}})
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, Vec2{0, -1}})
		break
	case '\\':
		var dir Vec2
		switch b.dir {
		case NORTH:
			dir = WEST
			break
		case SOUTH:
			dir = EAST
			break
		case WEST:
			dir = NORTH
			break
		case EAST:
			dir = SOUTH
			break
		}
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, dir})
		break
	case '/':
		var dir Vec2
		switch b.dir {
		case NORTH:
			dir = EAST
			break
		case SOUTH:
			dir = WEST
			break
		case WEST:
			dir = SOUTH
			break
		case EAST:
			dir = NORTH
			break
		}
		newBeams = append(newBeams, Beam{Vec2{b.pos.x, b.pos.y}, dir})
		break
	}
	for i := range newBeams {
		newBeams[i].pos.x += newBeams[i].dir.x
		newBeams[i].pos.y += newBeams[i].dir.y
	}
	return newBeams, nil
}
func lase(b Beam) {
	// base case
	if !beamField.inBounds(b.pos) || beamField.m[b] {
		return
	}
	beamField.m[b] = true
	// update dir and pos
	newBeams, err := b.next()
	if err != nil {
		return
	}
	for _, beam := range newBeams {
		lase(beam)
	}
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
	beamField = &field
	beamField.m = BeamMap{}
	start := Beam{Vec2{0, 0}, Vec2{1, 0}}
	lase(start) // recursive call that gets the business done
	posMap := map[Vec2]bool{}
	for k := range beamField.m {
		posMap[k.pos] = true
	}
	return len(posMap)
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
	beamField = &field
	beamStarts := []Beam{}
	height := len(field.f)
	width := len(field.f[0])
	for i := 0; i < height; i++ {
		beamStarts = append(beamStarts, Beam{Vec2{0, i}, Vec2{1, 0}})          // from left
		beamStarts = append(beamStarts, Beam{Vec2{width - 1, i}, Vec2{-1, 0}}) // from right
	}
	for i := 0; i < width; i++ {
		beamStarts = append(beamStarts, Beam{Vec2{i, 0}, Vec2{0, 1}})           // from top
		beamStarts = append(beamStarts, Beam{Vec2{i, height - 1}, Vec2{0, -1}}) // from bottom
	}
	maxExcited := 0
	for _, start := range beamStarts {
		beamField.m = BeamMap{}
		lase(start) // recursive call that gets the business done
		posMap := map[Vec2]bool{}
		for k := range beamField.m {
			posMap[k.pos] = true
		}
		if len(posMap) > maxExcited {
			maxExcited = len(posMap)
		}
	}
	return maxExcited
}

func main() {
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}
