package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

const (
	TEN   = 10
	JACK  = 11
	QUEEN = 12
	KING  = 13
	ACE   = 14
)

const (
	HIGH_CARD       = 0
	ONE_PAIR        = 1
	TWO_PAIR        = 2
	THREE_OF_A_KIND = 3
	FULL_HOUSE      = 4
	FOUR_OF_A_KIND  = 5
	FIVE_OF_A_KIND  = 6
)

type Hand struct {
	hand [5]int
	bid  int
	kind int
}

func NewHand(i string) Hand {
	hand := Hand{}
	cur := 0
	for ; cur < len(hand.hand); cur++ {
		var val int
		c := i[cur]
		switch c {
		case 'A':
			val = ACE
		case 'K':
			val = KING
		case 'Q':
			val = QUEEN
		case 'J':
			val = JACK
		case 'T':
			val = TEN
		default:
			val, _ = strconv.Atoi(i[cur : cur+1])
		}
		hand.hand[cur] = val
	}
	cur++ // skip a space

	saved := cur
	for ; cur < len(i) && unicode.IsDigit(rune(i[cur])); cur++ {
	}
	num, _ := strconv.Atoi(i[saved:cur])
	hand.bid = num
	return hand
}

func parse(scanner *bufio.Scanner) []Hand {
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, NewHand(line))
	}

	return hands
}

func calculateWinnings(hands *[]Hand) int {
	winnings := 0
	for i, hand := range *hands {
		winnings += hand.bid * (i + 1)
	}

	return winnings
}

func (self *Hand) classifyHand() {
	cardMap := make(map[int]int)
	for _, c := range self.hand {
		cardMap[c]++
	}
	switch len(cardMap) {
	case 1:
		self.kind = FIVE_OF_A_KIND
	case 2:
		for _, v := range cardMap {
			if v == 4 || v == 1 {
				self.kind = FOUR_OF_A_KIND
				break
			} else {
				self.kind = FULL_HOUSE
				break
			}
		}
	case 3:
		for _, v := range cardMap {
			if v == 3 {
				self.kind = THREE_OF_A_KIND
				break
			} else if v == 2 {
				self.kind = TWO_PAIR
			}
		}
	case 4:
		self.kind = ONE_PAIR
	case 5:
		self.kind = HIGH_CARD
	}
}

func (self *Hand) classifyWildHand() {
	cardMap := make(map[int]int)
	for _, c := range self.hand {
		cardMap[c]++
	}
	jackNum := cardMap[JACK]
	switch len(cardMap) {
	case 1:
		self.kind = FIVE_OF_A_KIND
	case 2:
		if jackNum > 0 {
			self.kind = FIVE_OF_A_KIND
		} else {
			for _, v := range cardMap {
				if v == 4 || v == 1 {
					self.kind = FOUR_OF_A_KIND
					break
				} else {
					self.kind = FULL_HOUSE
					break
				}
			}
		}
	case 3:
		if jackNum == 1 {
			for _, v := range cardMap {
				if v == 3 {
					self.kind = FOUR_OF_A_KIND
					break
				} else if v == 2 {
					self.kind = FULL_HOUSE
                    break
				}
			}
		} else if jackNum >= 2 {
			self.kind = FOUR_OF_A_KIND
        } else {
			for _, v := range cardMap {
				if v == 3 {
					self.kind = THREE_OF_A_KIND
					break
				} else if v == 2 {
					self.kind = TWO_PAIR
                    break
				}
			}
		}
	case 4:
		if jackNum > 0 {
			self.kind = THREE_OF_A_KIND
		} else {
			self.kind = ONE_PAIR
		}
	case 5:
		if jackNum > 0 {
			self.kind = ONE_PAIR
		} else {
			self.kind = HIGH_CARD
		}
	}
}

type ByHand []Hand

func (b ByHand) Len() int      { return len(b) }
func (b ByHand) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByHand) Less(i, j int) bool {
	if b[i].kind < b[j].kind {
		return true
	} else if b[i].kind > b[j].kind {
		return false
	} else {
		lesser := false
		for k := 0; k < len(b[i].hand); k++ {
			if b[i].hand[k] < b[j].hand[k] {
				lesser = true
				break
			} else if b[i].hand[k] > b[j].hand[k] {
				lesser = false
				break
			}
		}
		return lesser
	}
}

type ByWildHand []Hand

func (b ByWildHand) Len() int      { return len(b) }
func (b ByWildHand) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByWildHand) Less(i, j int) bool {
	if b[i].kind < b[j].kind {
		return true
	} else if b[i].kind > b[j].kind {
		return false
	} else {
		lesser := false
		for k := 0; k < len(b[i].hand); k++ {
			if b[i].hand[k] == JACK && b[j].hand[k] != JACK {
				lesser = true
                break
			} else if b[i].hand[k] != JACK && b[j].hand[k] == JACK {
				lesser = false
                break
			} else if b[i].hand[k] < b[j].hand[k] {
				lesser = true
				break
			} else if b[i].hand[k] > b[j].hand[k] {
				lesser = false
				break
			}
		}
		return lesser
	}
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// parse into hand and bid
	hands := parse(scanner)
	for i := 0; i < len(hands); i++ {
		hands[i].classifyHand()
	}
	// sort data based on some ranking function
	sort.Sort(ByHand(hands))

	// iterate through hands and multiply position (rank) by bid
	return calculateWinnings(&hands)
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// parse into hand and bid
	hands := parse(scanner)
	for i := 0; i < len(hands); i++ {
		hands[i].classifyWildHand()
	}
	// sort data based on some ranking function
	sort.Sort(ByWildHand(hands))

	// iterate through hands and multiply position (rank) by bid
	return calculateWinnings(&hands)
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
