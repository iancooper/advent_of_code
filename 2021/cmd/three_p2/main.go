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

	_, oxygen := aoc.OxygenGeneratorRating(readings)
	_, co2 := aoc.CO2ScrubberRating(readings)

	product := aoc.ConvertBinaryStringToInt(oxygen) * aoc.ConvertBinaryStringToInt(co2)

	log.Printf("%d readings", product)
}
