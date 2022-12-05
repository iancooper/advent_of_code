package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code/2022/internal"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/Users/ian.cooper/go/src/github.com/iancooper/advent_of_code/2022/test/testdata/one.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var packs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		packs = append(packs, s)
	}

	greatest := aoc.GreatestCaloriesCounter(packs)
	log.Printf("%d greatest calories in pack", greatest)
}
