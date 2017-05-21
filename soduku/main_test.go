package main

// Tests in this file serve as end-to-end tests that validate whether
// the algorithm can correctly solve a wide variety of 9x9 Sodoku grids.

import (
	"encoding/json"
	"os"
	"testing"
)

const eulerSolutionsFileName = "test_data/euler_solutions.json"

// BenchmarkEuler tests the code on the Project Euler repository of Sudoko grids.
// Note that each grid is treated as a single iteration for the purposes of
// computing averaged statistics, rather than a complete run through the dataset
// being one iteration in order to get more useful memory allocation statistics
func BenchmarkEuler(b *testing.B) {
	b.ReportAllocs()
	// Ensure that the number of benchmarks is a multiple of len(eulerGrids)
	numGrids := len(eulerGrids)
	b.N = int(b.N/numGrids) * numGrids

	// Run the test
	for i := 0; i < b.N; i++ {
		solveGrid(*eulerGrids[i%numGrids])
	}
}

func TestEuler(t *testing.T) {
	eulerSolnsFile, err := os.Open(eulerSolutionsFileName)
	if err != nil {
		t.Fatalf("Unable to load soltions file: %v", err)
		return
	}
	defer eulerSolnsFile.Close()
	eulerSolns := make([][9][9]GridValue, len(eulerGrids))
	json.NewDecoder(eulerSolnsFile).Decode(&eulerSolns)
	const eulerSum = 24702

	results, sum, err := solveEuler()
	if err != nil {
		t.Fatalf("Solving Euler problems failed with error %v", err)
	}

	for i, soln := range results {
		if soln != eulerSolns[i] {
			t.Errorf("Euler problem #%v returned:\n%v\nbut should have returned:\n%v", i+1, soln, eulerSolns[i])
		}
	}

	if sum != eulerSum {
		t.Errorf("Combined Euler problem returned %d but should have returned %d", sum, eulerSum)
	}

}
