package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code_2021/internal"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("../../test/testdata/one.input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var depths []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s, _ := strconv.Atoi(scanner.Text())
		depths = append(depths, s)
	}

	increases := aoc.DepthIncreases(depths)
	log.Printf("%d increases in depth", increases)
}
