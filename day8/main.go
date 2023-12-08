package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

// TODO: Make into a parsing package
type advance func(rune) bool

func advanceTil(line *string, cur int, fn advance) int {
	next := cur
	for ; next < len(*line) && !fn(rune((*line)[next])); next++ {
	}
	return next
}
func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}
func isNotLetter(r rune) bool {
	return !unicode.IsLetter(r)
}
func isSpace(r rune) bool {
	return r == ' '
}
func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}
func isNotDigit(r rune) bool {
	return !unicode.IsDigit(r)
}
func isLetterOrDigit(r rune) bool {
	return isLetter(r) || isDigit(r)
}
func isNotLetterOrDigit(r rune) bool {
	return !isLetterOrDigit(r)
}
func isRParentheses(r rune) bool {
	return r != ')'
}
func isComma(r rune) bool {
	return r == ','
}

// ^^^TODO: Make into a parsing package

type Node struct {
	left  string
	right string
}
type Directions struct {
	inner []bool
	cur   int
}

func (d *Directions) push(b bool) {
	d.inner = append(d.inner, b)
}
func (d *Directions) len() int {
	return len(d.inner)
}
func (d *Directions) next() bool {
	out := d.inner[d.cur]
	d.cur++
	if d.cur == d.len() {
		d.cur = 0
	}
	return out
}

func parseDirections(line *string) Directions {
	d := Directions{}
	for i := 0; i < len(*line); i++ {
		var turnRight bool
		switch (*line)[i] {
		case 'R':
			turnRight = true
		case 'L':
			turnRight = false
		}
		d.push(turnRight)
	}
	return d
}

func parseNode(line string) (string, Node) {
	var (
		cur  int = 0
		next int = 0
		key  string
		node Node
	)
	next = advanceTil(&line, cur, isSpace)
	key = line[cur:next]
	cur = next
	cur = advanceTil(&line, cur, isLetterOrDigit)
	next = advanceTil(&line, cur, isComma)
	node.left = line[cur:next]
	cur = next
	cur = advanceTil(&line, cur, isLetterOrDigit)
	next = advanceTil(&line, cur, isNotLetterOrDigit)
	node.right = line[cur:next]
	cur = next
	return key, node
}

func parse(scanner *bufio.Scanner) (Directions, map[string]Node) {
	scanner.Scan()
	line := scanner.Text()
	// exapmle usage
	directions := parseDirections(&line)
	scanner.Scan()
	line = scanner.Text()
	m := map[string]Node{}
	for scanner.Scan() {
		key, node := parseNode(scanner.Text())
		m[key] = node
	}
	cur := 0
	cur = advanceTil(&line, cur, isLetter)
	return directions, m
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// parse into hand and bid
	d, m := parse(scanner)
	key := "AAA"
	count := 0
	for ; key != "ZZZ"; count++ {
		turnRight := d.next()
		cur := m[key]
		if turnRight {
			key = cur.right
		} else {
			key = cur.left
		}
	}

	// iterate through hands and multiply position (rank) by bid
	return count
}

func part2() int {
    file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// parse into hand and bid
	d, m := parse(scanner)
	keys := []string{}
	for key, _ := range m {
		if key[2] == 'A' {
			keys = append(keys, key)
		}
	}
	count := 0
    allEndInZ := false
	for ; !allEndInZ; count++ {
		turnRight := d.next()
        theseEndInZ := true
		for i, key := range keys {
			cur := m[key]
			if turnRight {
				keys[i] = cur.right
			} else {
				keys[i] = cur.left
			}
            if keys[i][2] != 'Z' {
                theseEndInZ = false
            }
		}
        allEndInZ = theseEndInZ
	}

	// iterate through hands and multiply position (rank) by bid
	return count
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 1: ", part2())
}
