// main_bench
package main

import (
	"testing"
)

func BenchmarkEuler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(eulerGrids); i++ {
			solveGrid(eulerGrids[i])
		}
	}
}

func TestEuler()
