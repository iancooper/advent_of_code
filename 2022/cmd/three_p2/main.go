package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code/2022/internal"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/Users/ian.cooper/go/src/github.com/iancooper/advent_of_code/2022/test/testdata/three.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var backpacks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		backpacks = append(backpacks, s)
	}

	score := aoc.PriorityByBadge(backpacks, 3)
	log.Printf("%d priority items", score)
}
