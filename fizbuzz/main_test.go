// test.go
package main

import (
	"math/rand"
	"testing"
)

var testTable = []struct {
	input  int
	output string
}{
	{1, "1"},
	{3, "fizz"},
	{5, "buzz"},
	{15, "fizzbuzz"}}

func TestClassicFizzbuzz(t *testing.T) {
	for _, testcase := range testTable {
		result := handleNumber(testcase.input)
		if result != testcase.output {
			t.Errorf("Number %d failed: output was %s but should have been %s",
				testcase.input, result, testcase.output)
		}
	}
}

const FUZZ_COUNT = 1000

func TestFuzzing(t *testing.T) {
	for i := 0; i < FUZZ_COUNT; i++ {
		n := rand.Intn(100)
		result := handleNumber(n)

		// Check for a 0-length output
		if len(result) == 0 {
			t.Errorf("Number %d failed: output had length 0", n)
		}
	}
}
