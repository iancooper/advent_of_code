package advent_of_code

import "strconv"

func FindGammaBits(report []string) string {

	cols, bitSetCounts, bitNotSetCounts := findBitCounts(report)

	gammaBits := setGammaBits(cols, bitSetCounts, bitNotSetCounts)

	return string(gammaBits)
}

func FindEpsilonBits(report []string) string {

	cols, bitSetCounts, bitNotSetCounts := findBitCounts(report)

	epsilonBits := setEpsilonBits(cols, bitSetCounts, bitNotSetCounts)

	return string(epsilonBits)
}

func OxygenGeneratorRating(report []string) (bool, string) {

	return lifeSupportRating(report, filterTtoMostFrequent)
}

func CO2ScrubberRating(report []string) (bool, string) {

	return lifeSupportRating(report, filterToLeastFrequent)
}

func lifeSupportRating(report []string, filter func(int, []int, []int, []string, int) []string) (bool, string) {
	var frequentBitMatches []string
	frequentBitMatches = report

	cols, bitSetCounts, bitNotSetCounts := findBitCounts(frequentBitMatches)

	for i := 0; i < cols; i++ {

		frequentBitMatches = filter(cols, bitSetCounts, bitNotSetCounts, frequentBitMatches, i)

		//if we have only one match left, we are done
		if len(frequentBitMatches) == 1 {
			break
		}

	}

	return len(frequentBitMatches) == 1, frequentBitMatches[0]
}

func filterTtoMostFrequent(cols int, bitSetCounts []int, bitNotSetCounts []int, frequentBitMatches []string, i int) []string {
	cols, bitSetCounts, bitNotSetCounts = findBitCounts(frequentBitMatches)

	if bitSetCounts[i] > bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '1')
	} else if bitSetCounts[i] < bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '0')
	} else if bitSetCounts[i] == bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '1')
	}
	return frequentBitMatches
}

func filterToLeastFrequent(cols int, bitSetCounts []int, bitNotSetCounts []int, frequentBitMatches []string, i int) []string {
	cols, bitSetCounts, bitNotSetCounts = findBitCounts(frequentBitMatches)

	if bitSetCounts[i] < bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '1')
	} else if bitSetCounts[i] > bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '0')
	} else if bitSetCounts[i] == bitNotSetCounts[i] {
		frequentBitMatches = findMatchingReports(frequentBitMatches, i, '0')
	}
	return frequentBitMatches
}

func ConvertBinaryStringToInt(binary string) int64 {

	i, _ := strconv.ParseInt(binary, 2, 64)
	return i
}

func findMatchingReports(candidates []string, pos int, match rune) []string {
	matches := make([]string, 0, len(candidates))

	for _, report := range candidates {
		chars := []rune(report)
		if chars[pos] == match {
			matches = append(matches, report)
		}
	}

	return matches
}

func setEpsilonBits(cols int, bitSetCounts []int, bitNotSetCounts []int) []byte {
	//set according to lowest frequency value
	//TODO: We could just reverse the gamma bits to avoid the second loop
	epsilonBits := make([]byte, 0, cols)
	for j := 0; j < cols; j++ {
		if bitSetCounts[j] >= bitNotSetCounts[j] {
			epsilonBits = append(epsilonBits, []byte("0")...)
		} else {
			epsilonBits = append(epsilonBits, []byte("1")...)
		}
	}
	return epsilonBits
}

func setGammaBits(cols int, bitSetCounts []int, bitNotSetCounts []int) []byte {
	//set according to most fequent value
	gammaBits := make([]byte, 0, cols)
	for j := 0; j < cols; j++ {
		if bitSetCounts[j] >= bitNotSetCounts[j] {
			gammaBits = append(gammaBits, []byte("1")...)
		} else {
			gammaBits = append(gammaBits, []byte("0")...)
		}
	}
	return gammaBits
}

func findBitCounts(report []string) (int, []int, []int) {
	cols := len(report[0])
	bitSetCounts := make([]int, cols)
	bitNotSetCounts := make([]int, cols)

	//figure out the count of bits
	for _, v := range report {
		chars := []rune(v)
		for i := 0; i < cols; i++ {
			if string(chars[i]) == "1" {
				bitSetCounts[i] += 1
			} else {
				bitNotSetCounts[i] += 1
			}
		}
	}
	return cols, bitSetCounts, bitNotSetCounts
}
