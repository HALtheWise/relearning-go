// main_bench
package main

import (
	"encoding/json"
	"os"
	"testing"
)

func BenchmarkEuler(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(eulerGrids); i++ {
			solveGrid(eulerGrids[i])
		}
	}
}

func TestEuler(t *testing.T) {
	eulerSolnsFile, err := os.Open("euler_solutions.json")
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
			t.Errorf("Euler problem #%v returned:\n%s\nbut should have returned:\n%s", i+1, soln, eulerSolns[i])
		}
	}

	if sum != eulerSum {
		t.Errorf("Combined Euler problem returned %d but should have returned %d", sum, eulerSum)
	}

}
