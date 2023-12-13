package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

var (
	NORTH = Pos{0, -1}
	EAST  = Pos{1, 0}
	SOUTH = Pos{0, 1}
	WEST  = Pos{-1, 0}
)

type Pos struct {
	x int
	y int
}

func (p *Pos) add(dir Pos) Pos {
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
	return Slot{0, -1} // null byte if not inside
}

func (p *Player) nextDir() {
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
		if p.dir == NORTH {
			p.dir = Pos{1, 0}
		} else {
			p.dir = Pos{0, 1}
		}
		break
	case '7':
		if p.dir == NORTH {
			p.dir = Pos{-1, 0}
		} else {
			p.dir = Pos{0, 1}
		}
		break
	case 'J':
		if p.dir == SOUTH {
			p.dir = Pos{-1, 0}
		} else {
			p.dir = Pos{0, -1}
		}
		break
	case 'L':
		if p.dir == SOUTH {
			p.dir = Pos{1, 0}
		} else {
			p.dir = Pos{0, -1}
		}
		break
	case 'S':
		oppositeDir := Pos{-p.dir.x, -p.dir.y}
		lastPos := p.pos.add(oppositeDir)
		possiblePos := p.determineValid()
		for _, pos := range possiblePos {
			if pos != lastPos {
                p.dir = Pos{pos.x - p.pos.x, pos.y - p.pos.y}
            }
		}
	}
}
func (p *Player) moveToNext() {
    p.pos = p.pos.add(p.dir) // add in next direction
    p.nextDir()
}
func (p *Player) updateDist() {
	p.dist++
	(*p.f)[p.pos.y][p.pos.x].dist = p.dist
}

func (p *Player) determineValid() [2]Pos {
	var pos [2]Pos
	posInArray := func(c byte, arr []byte) bool {
		for _, b := range arr {
			if b == c {
				return true
			}
		}
		return false
	}
    
	northPos := p.pos.add(NORTH)
	eastPos := p.pos.add(EAST)
	southPos := p.pos.add(SOUTH)
	westPos := p.pos.add(WEST)
    idx := 0
	if posInArray(p.getSlot(northPos).val, []byte{'7', '|', 'F'}) {
		pos[idx] = northPos
        idx++
	}
	if posInArray(p.getSlot(eastPos).val, []byte{'J', '-', '7'}) {
		pos[idx] = eastPos
        idx++
	}
	if posInArray(p.getSlot(southPos).val, []byte{'|', 'L', 'J'}) {
		pos[idx] = southPos
        idx++
	}
	if posInArray(p.getSlot(westPos).val, []byte{'L', '-', 'F'}) {
		pos[idx] = westPos
        idx++
	}
	return pos
}
func (p *Player) hasUnchartedTerritory() bool {
	return p.getSlot(p.pos.add(p.dir)).dist == 0
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
		p1.updateDist()
		p1.moveToNext()
		p2.updateDist()
		p2.moveToNext()
	}

	return maxDist(p1.f)
}

func (p *Player) getRightDir() Pos {
	var right Pos
	switch p.dir {
	case NORTH:
		right = EAST
		break
	case EAST:
		right = SOUTH
		break
	case SOUTH:
		right = WEST
		break
	case WEST:
		right = NORTH
		break
	}
	return right
}

func (p *Player) inBounds(pos Pos) bool {
    return pos.x >= 0 && pos.y >= 0 && pos.y < len(*p.f) && pos.x < len((*p.f)[pos.y])
}
func (p *Player) findAllDots(pos Pos, m *map[Pos]bool) {
	// base
	if p.getSlot(pos).dist > 0 {
        //fmt.Println("Sending back because dist > 0")
        return
    } else if (*m)[pos] {
        //fmt.Println("Sending back because already true")
        return
    } else if !p.inBounds(pos) {
        //fmt.Println("Sending back because not inbounds")
        return
	} else {
        //fmt.Println("Valid dist 0")
		(*m)[pos] = true
    }
	// during
	for _, dir := range []Pos{NORTH, EAST, SOUTH, WEST} {
		newPos := pos.add(dir)
		p.findAllDots(newPos, m)
	}
}

func (p *Player) countInside() int {
	pold := p
	startFound := false
	start := Pos{0, 0}
	for y, row := range *p.f {
		for x, s := range row {
			if s.dist > 0 {
				startFound = true
				start = Pos{x, y}
				break
			}
		}
		if startFound {
			break
		}
	}
	p.dir = Pos{1, 0} // always moving right from the start
    p.pos = start
	m := map[Pos]bool{}
	for {
		p.moveToNext()

        posRight := p.pos.add(p.getRightDir())
        p.findAllDots(posRight, &m)
		if p.pos == start {
			break
		}
	}
	p = pold
    p.printInside(&m)
	return len(m)
}

func (p *Player) printInside(m *map[Pos]bool) {
    for y, row := range *p.f {
        for x := range row {
            if (*m)[Pos{x, y}] {
                fmt.Printf("x")
            } else {
                fmt.Printf(".")
            }

        }
        fmt.Println()
    }
}
func (p *Player) printPath() {
    for _, row := range *p.f {
        for _, c := range row {
            if c.dist == 0 {
                fmt.Printf(".")
            } else {
                fmt.Printf("x")
            }

        }
        fmt.Println()
    }
    fmt.Println()
}

func part2() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	f := parse(scanner) // field
	p1 := NewPlayer(&f)
	p1.dist = 1
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
		p1.updateDist()
		p2.moveToNext()
		p2.updateDist()
	}
    p1.printPath()
	return p1.countInside()
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
