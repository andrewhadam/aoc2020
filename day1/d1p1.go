package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []string
	//	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s
}

func main() {
	part1()
	part2()
}

func part1() {

	var input []string = readInput()
	sort.Strings(input)

	var found = false

	for i := 0; i < len(input); i++ {
		for j := len(input) - 1; j >= 0; j-- {
			ii, _ := strconv.Atoi(input[i])
			jj, _ := strconv.Atoi(input[j])
			if ii+jj < 2020 {
				continue
			} else if ii+jj == 2020 {
				fmt.Println("Found it")
				fmt.Printf("I: %d, J: %d\n", ii, jj)
				fmt.Printf("Solution: %d\n", ii*jj)
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

}

func part2() {

	var input []string = readInput()
	sort.Strings(input)

	var found = false

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			for k := 0; k < len(input); k++ {
				ii, _ := strconv.Atoi(input[i])
				jj, _ := strconv.Atoi(input[j])
				kk, _ := strconv.Atoi(input[k])
				if i == j || i == k || j == k {
					fmt.Println("Funny")
					continue
					//else if ii+jj+kk > 2020 {
					//	break
				} else if ii+jj+kk == 2020 {
					fmt.Println("Found it")
					fmt.Printf("I: %d, J: %d, K: %d\n", ii, jj, kk)
					fmt.Printf("Solution: %d\n", ii*jj*kk)
					found = true

				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
}
