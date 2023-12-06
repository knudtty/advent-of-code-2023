package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Map struct {
	dest int
	src  int
	rng  int
}
type AlmanacMap struct {
	inner []Map
}

func (self *AlmanacMap) push(m Map) {
	self.inner = append(self.inner, m)
}

func (self *AlmanacMap) mapValue(n int) int {
	for _, m := range self.inner {
		if n >= m.src && n < m.src+m.rng {
			return n - m.src + m.dest
		}
	}
	return n
}

type Node struct {
	link  string
	value AlmanacMap
}

func NewNode(link *string) *Node {
	node := Node{}
	node.link = *link
	return &node
}

type AlmanacGraph struct {
	inner map[string]Node
}

func parseSeeds(line *string) (seeds []int) {
	var peekIdx, cur int
	for cur = 0; (*line)[cur] != ':'; cur++ {
	}
	for cur < len(*line) {
		for ; !unicode.IsDigit(rune((*line)[cur])); cur++ {
		} // skip until digit
		for peekIdx = cur; peekIdx < len(*line) && unicode.IsDigit(rune((*line)[peekIdx])); peekIdx++ {
		} // skip until digit
		num, _ := strconv.Atoi((*line)[cur:peekIdx])
		seeds = append(seeds, num)
		cur = peekIdx
	}
	return seeds
}

type Seed struct {
	start  int
	length int
}

func parseSeeds2(line *string) (seeds []Seed) {
	var peekIdx, cur int
	for cur = 0; (*line)[cur] != ':'; cur++ {
	}
	count := 0
	for cur < len(*line) {
		for ; !unicode.IsDigit(rune((*line)[cur])); cur++ {
		} // skip until digit
		for peekIdx = cur; peekIdx < len(*line) && unicode.IsDigit(rune((*line)[peekIdx])); peekIdx++ {
		} // skip until digit
		num, _ := strconv.Atoi((*line)[cur:peekIdx])
		if count%2 == 0 {
			seeds = append(seeds, Seed{})
			seeds[len(seeds)-1].start = num
		} else {
			seeds[len(seeds)-1].length = num
		}
		cur = peekIdx
		count++
	}
	return seeds
}

type NodeMap = map[string]*Node

func parseMaps(scanner *bufio.Scanner) NodeMap {
	var key, link string
	theMap := make(NodeMap)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
		} else if unicode.IsLetter(rune(line[0])) {
			cur, next := 0, 0
			for next = 0; line[next] != '-'; next++ {
			}
			key = line[cur:next]
			for cur = next + 1; line[cur] != '-'; cur++ {
			}
			cur++ // past the -
			for next = cur; line[next] != ' '; next++ {
			}
			link = line[cur:next]
			theMap[key] = NewNode(&link)
		} else if unicode.IsDigit(rune(line[0])) {
			node := theMap[key]
			var dest, src, rng int
			next, cur := 0, 0
			for next = cur; unicode.IsDigit(rune(line[next])); next++ {
			} // skip until not a digit
			dest, _ = strconv.Atoi(line[cur:next])
			cur = next
			cur++
			for next = cur; unicode.IsDigit(rune(line[next])); next++ {
			} // skip until not a digit
			src, _ = strconv.Atoi(line[cur:next])
			cur = next
			cur++
			for next = cur; next < len(line) && unicode.IsDigit(rune(line[next])); next++ {
			} // skip until not a digit
			rng, _ = strconv.Atoi(line[cur:next])
			node.value.push(Map{dest, src, rng})
		}
	}
	return theMap
}

func parse(scanner *bufio.Scanner) ([]int, NodeMap) {
	scanner.Scan()
	line := scanner.Text()
	seeds := parseSeeds(&line)
	maps := parseMaps(scanner)
	return seeds, maps
}

func parse2(scanner *bufio.Scanner) ([]Seed, NodeMap) {
	scanner.Scan()
	line := scanner.Text()
	seeds := parseSeeds2(&line)
	maps := parseMaps(scanner)
	return seeds, maps
}

func calcLocation(seed int, nodeMap *NodeMap) int {
	key := "seed"
	node := (*nodeMap)[key]
	ans := seed
	for ; node != nil; node = (*nodeMap)[node.link] {
		ans = node.value.mapValue(ans)
	}
	return ans
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	seeds, maps := parse(scanner)
	minLoc := int(^uint(0) >> 1)
	for _, seed := range seeds {
		loc := calcLocation(seed, &maps)
		minLoc = min(loc, minLoc)
	}

	return minLoc
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	seeds, maps := parse2(scanner)
	minLoc := int(^uint(0) >> 1)
    total := len(seeds)
	for j, seed := range seeds {
		rng := seed.length / 10
		for i := seed.start; i < seed.start+seed.length; i++ {
			//fmt.Println(i)
			loc := calcLocation(i, &maps) // TODO
			minLoc = min(loc, minLoc)
			if (i-seed.length)%rng == 0 {
				fmt.Println((i-seed.start)/rng * 10, "% done ", j+1, " / ", total)
			}
		}
	}

	return minLoc
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
