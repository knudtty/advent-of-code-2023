package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

type Sequence []int

func (self *Sequence) nextOrder() Sequence {
	seq := Sequence{}
	for i := range (*self)[:len(*self)-1] {
		diff := (*self)[i+1] - (*self)[i]
		seq = append(seq, diff)
	}
	return seq
}
func (self *Sequence) allZero() bool {
	for _, v := range *self {
		if v != 0 {
			return false
		}
	}
	return true
}

func NewSequence(line *string) Sequence {
	seq := Sequence{}
	cur, next := 0, 0
	for cur < len(*line) {
		next = advanceTil(line, cur, isSpace)
		num, _ := strconv.Atoi((*line)[cur:next])
		seq = append(seq, num)
		cur = next + 1
	}
	return seq
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		orders := []Sequence{NewSequence(&line)}
		seq := &orders[0]
		for !seq.allZero() {
			orders = append(orders, seq.nextOrder())
			seq = &orders[len(orders)-1]
		}
		for i := len(orders) - 1; i > 0; i-- {
            newVal := orders[i-1][len(orders[i-1])-1] + orders[i][len(orders[i])-1]
			orders[i-1] = append(orders[i-1], newVal)
		}
		sum += orders[0][len(orders[0])-1]
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
	for scanner.Scan() {
		line := scanner.Text()
		orders := []Sequence{NewSequence(&line)}
		seq := &orders[0]
		for !seq.allZero() {
			orders = append(orders, seq.nextOrder())
			seq = &orders[len(orders)-1]
		}
		for i := len(orders) - 1; i > 0; i-- {
            newVal := orders[i-1][0] - orders[i][0] 
			orders[i-1] = append(Sequence{newVal}, orders[i-1]...)
		}
		sum += orders[0][0]
	}
	return sum
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
