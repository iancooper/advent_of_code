package advent_of_code

import "testing"

func TestMostCommonGammaBit(t *testing.T) {
	report := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expectedCommonGammaBit := "10110"

	commonGammaBit := FindGammaBits(report)

	if commonGammaBit != expectedCommonGammaBit {
		t.Errorf("Expect to see a gamma of %s but was %s", expectedCommonGammaBit, commonGammaBit)
	}

}

func TestMostCommonEpsilonBit(t *testing.T) {
	report := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expectedCommonEpsilonBit := "01001"

	commonEpsilonBit := FindEpsilonBits(report)

	if commonEpsilonBit != expectedCommonEpsilonBit {
		t.Errorf("Expect to see a gamma of %s but was %s", expectedCommonEpsilonBit, commonEpsilonBit)
	}

}

func TestOxygenGeneratorRating(t *testing.T) {
	report := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expectedOxygenGeneratorRating := "10111"

	found, oxygenGeneratorRating := OxygenGeneratorRating(report)

	if found == false {
		t.Errorf("Expected a single match but could not find one")
	}

	if oxygenGeneratorRating != expectedOxygenGeneratorRating {
		t.Errorf("Expected to see an oxygen rating of #{expectedOxygenGeneratorRating } but was #{oxygenGeneratorRating}")
	}
}

func TestCO2ScrubberRating(t *testing.T) {

	report := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expectedCO2ScrubberRating := "01010"

	found, C02ScrubberRating := CO2ScrubberRating(report)

	if found == false {
		t.Errorf("Expected a single match but could not find one")
	}

	if C02ScrubberRating != expectedCO2ScrubberRating {
		t.Errorf("Expected to see an oxygen rating of #{expectedCO2ScrubberRating } but was #{oxygenGeneratorRating}")
	}
}

func TestConvertBinaryStringToInt(t *testing.T) {

	binaryString := "10110"
	var expectedInt int64
	expectedInt = 22

	binaryAsInt := ConvertBinaryStringToInt(binaryString)

	if binaryAsInt != expectedInt {
		t.Errorf("Expect to a binary value of %d but was %d", expectedInt, binaryAsInt)
	}

}
