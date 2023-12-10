package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
)

type advance func(rune) bool

func advanceTil(line *string, cur int, fn advance) int {
	next := cur
	for ; next < len(*line) && !fn(rune((*line)[next])); next++ {
	}
	return next
}
func isSpace(r rune) bool {
	return r == ' '
}

type Slot struct {
	val  byte
	dist int
}

func NewSlots(line *string) []Slot {
	var slots []Slot
	for _, v := range *line {
		slot := Slot{byte(v), 0}
		slots = append(slots, slot)
	}
	return slots
}

func parse(scanner *bufio.Scanner) [][]Slot {
	field := [][]Slot{}
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, NewSlots(&line))
	}
	return field
}

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

type Pos struct {
	x int
	y int
}

func (p *Pos) move(dir Pos) Pos {
	return Pos{p.x + dir.x, p.y + dir.y}
}

type Player struct {
	pos  Pos
	dir  Pos
	dist int
	f    *[][]Slot
}

func NewPlayer(field *[][]Slot) Player {
	for y, row := range *field {
		for x, c := range row {
			if c.val == 'S' {
				return Player{
					Pos{
						x,
						y,
					},
					Pos{
						0,
						0,
					},
					0,
					field,
				}
			}
		}
	}
	return Player{}
}

func (p *Player) getSlot(pos Pos) Slot {
	if pos.y < len(*p.f) && pos.y >= 0 && pos.x < len((*p.f)[pos.y]) && pos.x >= 0 {
		return (*p.f)[pos.y][pos.x]
	}
	return Slot{0, 0} // null byte if not inside
}

func (p *Player) moveToNext() {
	p.pos = p.pos.move(p.dir) // move in next direction
	// determine next direction
	s := p.getSlot(p.pos)
	// TODO
	switch s.val {
	case '-':
		// do not modify direction
		break
	case '|':
		// do not modify direction
		break
	case 'F':
		north := Pos{0, -1}
		if p.dir == north {
			p.dir = Pos{1, 0}
		} else {
			p.dir = Pos{0, 1}
		}
		break
	case '7':
		north := Pos{0, -1}
		if p.dir == north {
			p.dir = Pos{-1, 0}
		} else {
			p.dir = Pos{0, 1}
		}
		break
	case 'J':
		south := Pos{0, 1}
		if p.dir == south {
			p.dir = Pos{-1, 0}
		} else {
			p.dir = Pos{0, -1}
		}
		break
	case 'L':
		south := Pos{0, 1}
		if p.dir == south {
			p.dir = Pos{1, 0}
		} else {
			p.dir = Pos{0, -1}
		}
		break
	}
	p.dist++
	(*p.f)[p.pos.y][p.pos.x].dist = p.dist
}

func (p *Player) determineValid() []Pos {
	var pos []Pos
	posInArray := func(c byte, arr []byte) bool {
		for _, b := range arr {
			if b == c {
				return true
			}
		}
		return false
	}
	north := Pos{p.pos.x, p.pos.y - 1}
	east := Pos{p.pos.x + 1, p.pos.y}
	south := Pos{p.pos.x, p.pos.y + 1}
	west := Pos{p.pos.x - 1, p.pos.y}
	if posInArray(p.getSlot(north).val, []byte{'7', '|', 'F'}) {
		pos = append(pos, north)
	}
	if posInArray(p.getSlot(east).val, []byte{'J', '-', '7'}) {
		pos = append(pos, east)
	}
	if posInArray(p.getSlot(south).val, []byte{'|', 'L', 'J'}) {
		pos = append(pos, south)
	}
	if posInArray(p.getSlot(west).val, []byte{'L', '-', 'F'}) {
		pos = append(pos, west)
	}
	return pos
}
func (p *Player) hasUnchartedTerritory() bool {
	return p.getSlot(p.pos.move(p.dir)).dist == 0
}

func maxDist(f *[][]Slot) int {
	maxn := 0
	for _, row := range *f {
		for _, v := range row {
			maxn = max(v.dist, maxn)
		}
	}
	return maxn
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	f := parse(scanner) // field
	p1 := NewPlayer(&f)
	p2 := p1
	positions := p1.determineValid()
	positions[0].x -= p1.pos.x
	positions[0].y -= p1.pos.y
	positions[1].x -= p1.pos.x
	positions[1].y -= p1.pos.y
	p1.dir = positions[0]
	p2.dir = positions[1]
	for p1.hasUnchartedTerritory() && p2.hasUnchartedTerritory() {
		p1.moveToNext()
		p2.moveToNext()
	}

	return maxDist(p1.f)
}

func main() {
	fmt.Println("Part 1: ", part1())
}
