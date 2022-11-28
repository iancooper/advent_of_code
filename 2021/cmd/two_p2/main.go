package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code_2021/internal"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../../test/testdata/two.input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	changes := make([]string, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		changes = append(changes, s)
	}

	positionChanges := aoc.ReadPositionChangesFrom(changes)
	position := aoc.ApplyAimedPositionChanges(positionChanges)

	product := position.Horizontal * position.Vertical
	log.Printf("depth: %d distance %d navigation product %d", position.Vertical, position.Horizontal, product)
}
