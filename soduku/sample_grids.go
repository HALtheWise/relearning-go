// sample_grids
package main

var (
	blankGrid,
	nearlyFull *Grid
)

func init() {
	blankGrid, _ = GridFromString(`*********
*********
*********
*********
*********
*********
*********
*********
*********`)
	nearlyFull, _ = GridFromString(`123456789
123456789
123456789
123456789
123456789
123456789
123456789
123456789
19345678*`)
}
