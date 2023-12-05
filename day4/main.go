package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Queue struct {
	inner []Card
}

func (self *Queue) push(v Card) {
	self.inner = append(self.inner, v)
}
func (self *Queue) pop() Card {
	tmp := self.inner[0]
	self.inner = self.inner[1:]
	return tmp
}
func (self *Queue) isEmpty() bool {
	return len(self.inner) == 0
}

type Card struct {
	winners []int
	myNums  []int
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toDigit(c byte) int {
	return int(c - '0')
}

func parseCard(line *string) Card {
	cur := 0
	for (*line)[cur] != ':' {
		cur++
	}
	cur++ // advance past the space
	card := Card{}
	parseNum := 0
	for ; (*line)[cur] != '|'; cur++ {
		c := (*line)[cur]
		if isDigit(c) {
			parseNum = parseNum*10 + toDigit(c)
		} else {
			if parseNum > 0 {
				card.winners = append(card.winners, parseNum)
				parseNum = 0
			}
		}
	}
	parseNum = 0
	for ; cur < len(*line); cur++ {
		c := (*line)[cur]
		if isDigit(c) {
			parseNum = parseNum*10 + toDigit(c)
		} else {
			if parseNum > 0 {
				card.myNums = append(card.myNums, parseNum)
				parseNum = 0
			}
		}
	}
	if parseNum > 0 {
		card.myNums = append(card.myNums, parseNum)
	}
	return card
}

func (self *Card) countWinners() int {
	nWinners := 0
	for _, winner := range self.winners {
		for _, myNum := range self.myNums {
			if myNum == winner {
				nWinners++
			}
		}
	}
	return nWinners
}

func (self *Card) scoreCard() int {
	var score int
	nWinners := self.countWinners()
	if nWinners > 0 {
		score = 1
		for i := 1; i < nWinners; i++ {
			score *= 2
		}
	}
	return score
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
		card := parseCard(&line)
		sum += card.scoreCard()
	}
	return sum
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	cards := []Card{}
	nCards := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(&line))
		nCards = append(nCards, 1)
	}
	sum := 0
	// have a running queue based on the results of winning cards
	for cardNum, card := range cards {
        sum += nCards[cardNum]
		for i := 0; i < nCards[cardNum]; i++ {
			nWinners := card.countWinners()
			for j := 0; j < nWinners; j++ {
				nCards[j+cardNum+1]++
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
