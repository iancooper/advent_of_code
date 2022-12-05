package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code/2022/internal"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/Users/ian.cooper/go/src/github.com/iancooper/advent_of_code/2022/test/testdata/two.input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var games []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		games = append(games, s)
	}

	score := aoc.RoshamboContest(games)
	log.Printf("%d points scored in RPS", score)
}
