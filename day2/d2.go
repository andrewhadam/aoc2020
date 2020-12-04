package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passwords struct {
	startint int
	endint   int
	value    string
	passwd   string
}

func readInput() []passwords {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []passwords
	scanner := bufio.NewScanner(file)

	matchstart := regexp.MustCompile("(\\d*)-.*")
	matchend := regexp.MustCompile(".*-(\\d*).*")
	matchvalue := regexp.MustCompile("\\d*-\\d*\\s(\\w).*")
	matchpswd := regexp.MustCompile(".*: (.*)")

	for scanner.Scan() {
		rs, _ := strconv.Atoi(matchstart.FindStringSubmatch(scanner.Text())[1])
		rt, _ := strconv.Atoi(matchend.FindStringSubmatch(scanner.Text())[1])
		ru := matchvalue.FindStringSubmatch(scanner.Text())
		rv := matchpswd.FindStringSubmatch(scanner.Text())

		s = append(s, passwords{startint: rs, endint: rt, value: ru[1], passwd: rv[1]})

	}
	//	fmt.Println(s)
	return s
}

func main() {
	part1()
	part2()
}

func part1() {
	var input []passwords = readInput()
	var validpasswd int = 0

	for _, s := range input {
		c := strings.Count(s.passwd, s.value)

		if c <= s.endint && c >= s.startint {
			validpasswd += 1
		}
	}
	fmt.Println(validpasswd)
}

func part2() {
	var input []passwords = readInput()
	var validpasswd int = 0

	//	3-4 c: czlx

	for _, s := range input {

		fmt.Printf("Start: %d, End: %d\n", s.startint, s.endint)
		fmt.Printf("Value, %s, Passwd: %s\n", s.value, s.passwd)
		if string(s.passwd[s.startint-1]) == s.value || string(s.passwd[s.endint-1]) == s.value {
			//		if string(s.startint) == s.value || string(s.endint) == s.value {
			//			fmt.Println((s.startint - 1))
			//			fmt.Println((s.endint - 1))
			//			fmt.Println(string(s.passwd[s.endint] - 1))

			if string(s.passwd[s.startint-1]) == s.value && string(s.passwd[s.endint-1]) == s.value {
				fmt.Println("Internal not valid")
				continue
			}
			validpasswd += 1
			fmt.Println("Incremented")
		} else {
			fmt.Println("Not valid")
		}
	}
	//		if string(s.startint) != s.value && string(s.endint) != s.value {
	//			fmt.Println("Not valid")
	//		} else
	//		 else if string(s.startint) == s.value || string(s.endint) == s.value {
	//			validpasswd += 1
	//		}

	fmt.Println(validpasswd)
}
