package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Race struct {
	time int
	dist int
}

func (self *Race) checkPossibleFinishes() int {
	sum := 0
	for i := 0; i < self.time; i++ {
		distance := i * (self.time - i)
		if distance > self.dist {
			sum++
		}
	}
	return sum
}

func parse(scanner *bufio.Scanner) []Race {
	var races []Race
	scanner.Scan()
	line := scanner.Text()
	cur := 0
	next := 0
	for cur < len(line) {
		for ; cur < len(line) && !unicode.IsDigit(rune(line[cur])); cur++ {
		} // skip until digit is found
		next = cur
		for ; next < len(line) && unicode.IsDigit(rune(line[next])); next++ {
		} // skip until digit is not found
		num, _ := strconv.Atoi(line[cur:next])
		races = append(races, Race{})
		races[len(races)-1].time = num
		cur = next
	}
	scanner.Scan()
	line = scanner.Text()
	cur = 0
	next = 0
	count := 0
	for cur < len(line) {
		for ; cur < len(line) && !unicode.IsDigit(rune(line[cur])); cur++ {
		} // skip until digit is found
		next = cur
		for ; next < len(line) && unicode.IsDigit(rune(line[next])); next++ {
		} // skip until digit is not found
		num, _ := strconv.Atoi(line[cur:next])
		races[count].dist = num
		cur = next
		count++
	}
	return races
}

func parse2(scanner *bufio.Scanner) Race {
	var race Race
	scanner.Scan()
	line := scanner.Text()
	cur := 0
	next := 0
	time := ""
	for cur < len(line) {
		for ; cur < len(line) && !unicode.IsDigit(rune(line[cur])); cur++ {
		} // skip until digit is found
		next = cur
		for ; next < len(line) && unicode.IsDigit(rune(line[next])); next++ {
		} // skip until digit is not found
		time += line[cur:next]
		cur = next
	}
	num, _ := strconv.Atoi(time)
	race.time = num
	scanner.Scan()
	line = scanner.Text()
	cur = 0
	next = 0
	dist := ""
	for cur < len(line) {
		for ; cur < len(line) && !unicode.IsDigit(rune(line[cur])); cur++ {
		} // skip until digit is found
		next = cur
		for ; next < len(line) && unicode.IsDigit(rune(line[next])); next++ {
		} // skip until digit is not found
		dist += line[cur:next]
		cur = next
	}
	num, _ = strconv.Atoi(dist)
	race.dist = num

	return race
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	races := parse(scanner)
	total := 1
	for _, race := range races {
		sum := race.checkPossibleFinishes()
		total *= sum
	}

	return total
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	race := parse2(scanner)
	return race.checkPossibleFinishes()
}

func main() {
	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}
