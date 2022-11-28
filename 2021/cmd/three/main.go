package main

import (
	"bufio"
	"fmt"
	aoc "github.com/iancooper/advent_of_code_2021/internal"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../../test/testdata/three.input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	readings := make([]string, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readings = append(readings, scanner.Text())
	}

	gamma := aoc.FindGammaBits(readings)
	epsilon := aoc.FindEpsilonBits(readings)

	product := aoc.ConvertBinaryStringToInt(gamma) * aoc.ConvertBinaryStringToInt(epsilon)

	log.Printf("%d readings", product)
}
