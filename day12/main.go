package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func addFac(x int) int {
	out := 0
	for i := x; i > 0; i++ {
		out += i
	}
	return out
}

type Spring struct {
	row    string
	groups []int
}
type Cache map[string]int

func groupString(group *[]int) string {
    str := "("
    for _, g := range *group {
        str += fmt.Sprint(g) + ", "
    }
    if len(*group) == 1 {
        str = str[:len(str)-1]
    } else {
        str = str[:len(str)-2]
    }
    str += ")"
    return str
}

func countPossibilities(s Spring, cache *Cache) int {
    
    if len(s.groups) == 0 {
        if !strings.Contains(s.row, "#") {
            return 1
        } else {
            return 0
        }
    }
    groupStr := groupString(&s.groups)
    key := s.row + groupStr
    val, ok := (*cache)[key]
    if ok {
        return val
    }
    if len(s.row) == 0 {
        return 0
    }
    nextCharacter := s.row[0]
    nextGroup := s.groups[0]
    pound := func () int {
        if nextGroup > len(s.row) {
            return 0
        }
        thisGroup := strings.ReplaceAll(s.row[:nextGroup], "?", "#")
        if thisGroup != strings.Repeat("#", nextGroup) {
            return 0
        }
        if len(s.row) == nextGroup {
            if len(s.groups) == 1 {
                return 1
            } else {
                return 0
            }
        }
        if s.row[nextGroup] == '.' || s.row[nextGroup] == '?' {
            return countPossibilities(Spring{s.row[nextGroup+1:], s.groups[1:]}, cache)
        } 
        return 0
    }
    dot := func () int {
        return countPossibilities(Spring{s.row[1:], s.groups}, cache)
    }
    var out int
    if nextCharacter == '#' {
        out = pound()
    } else if nextCharacter == '.' {
        out = dot()
    } else if nextCharacter == '?' {
        out = dot() + pound()
    }
    // print
    str := fmt.Sprintf("%s ", s.row)
    str += groupStr
    str += " " + fmt.Sprint(out)
    //fmt.Println(str)

    (*cache)[key] = out
    return out
}

func NewSpring(line string) Spring {
	row := strings.Split(line, " ")
	groupsStrings := strings.Split(row[1], ",")
    groups := []int{}
    for _, group := range groupsStrings {
        num, _ := strconv.Atoi(string(group))
        groups = append(groups, num)
    }
	return Spring{
		row[0],
		groups,
	}
}
func parse(scanner *bufio.Scanner) []Spring {
    springs := []Spring{}
	for scanner.Scan() {
        springs = append(springs, NewSpring(scanner.Text()))
	}
	return springs
}

func part1() int {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    springs := parse(bufio.NewScanner(file))
    sum := 0
    for _, s := range springs {
        cache := Cache{}
        sum += countPossibilities(s, &cache)
        fmt.Println("----------")
    }

	return sum
}

func NewSpring2(line string) Spring {
	row := strings.Split(line, " ")
    record := row[0]
	groupsStrings := strings.Split(row[1], ",")
    groups := []int{}
    for _, group := range groupsStrings {
        num, _ := strconv.Atoi(string(group))
        groups = append(groups, num)
    }
    outRec := ""
    outGroups := []int{}
    for i := 0; i < 5; i++ {
        outRec += record + "?"
        outGroups = append(outGroups, groups...)
    }
    outRec = outRec[:len(outRec)-1]
	return Spring{
        outRec,
		outGroups,
	}
}
func parse2(scanner *bufio.Scanner) []Spring {
    springs := []Spring{}
	for scanner.Scan() {
        springs = append(springs, NewSpring2(scanner.Text()))
	}
	return springs
}

func part2() uint {
	// get line that the pipe traces.
	// Then walk around it clockwise. Look left and right.
	// Anything to the right is inside, anything left is outside
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

    springs := parse2(bufio.NewScanner(file))
    sum := uint(0)
    for _, s := range springs {
        cache := Cache{}
        fmt.Println(s.groups, s.row)
        ans := uint(countPossibilities(s, &cache))
        fmt.Println(ans)
        sum += ans
        fmt.Println("----------")
    }

	return sum
}

func main() {
    fmt.Println("Part 1:", part1())
    fmt.Println("Part 2:", part2())
}
