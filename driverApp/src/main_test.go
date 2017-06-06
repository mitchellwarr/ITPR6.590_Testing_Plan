package main

import (
	"math/rand"
	"testing"
)

func TestGetRandomGen(t *testing.T) {
	//ITPR6.590A2-SYS001
	var input string
	var randGen *rand.Rand
	var n float64
	var err error

	input = "100"
	randGen, _ = getRandomGen(input)
	n = randGen.Float64()

	AssertEqual(t, "Random gen with seed 100 should be 0.8165026937796166", n, 0.8165026937796166)

	input = "one hundred"
	randGen, err = getRandomGen(input)
	n = randGen.Float64()

	AssertEqual(t, "Random gen with error gen should be 0.9451961492941164", n, 0.9451961492941164)
	AssertNotEqual(t, "Random gen should return an error", err, nil)
}

// ITPR6.590A2-DRV002
// Test that a driver stops moving once at an exit


// ITPR6.590A2-SYS002
// App should iterate through 5 drivers