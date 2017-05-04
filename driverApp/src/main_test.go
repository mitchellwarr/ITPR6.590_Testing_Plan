package main

import (
	"math/rand"
	"testing"
)

func TestGetRandomGen(t *testing.T) {
	var input string
	var randGen *rand.Rand
	var n float64
	var err error

	input = "100"
	randGen, _ = getRandomGen(input)
	n = randGen.Float64()
	if n != 0.8165026937796166 {
		t.Error("Expected 0.8165026937796166, got ", n)
	}

	input = "one hundred"
	randGen, err = getRandomGen(input)
	n = randGen.Float64()
	if n != 0.9451961492941164 {
		t.Error("Expected 0.9451961492941164, got ", n)
	}
	if err == nil {
		t.Error("Expected error, got ", err)
	}
}

func TestGetInput(t *testing.T) {

}
