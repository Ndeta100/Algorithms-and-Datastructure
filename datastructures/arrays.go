package datastructures

func TwoDarray(numrows, numcols int) [][]int {
	cols := make([][]int, numrows)
	for i := 0; i < numrows; i++ {
		cols[i] = make([]int, numcols)
	}
	return cols
}

func ThreeDarray(numrows, numcols int) [][][]int {
	cols := make([][][]int, numrows)
	for i := 0; i < numrows; i++ {
		cols[i] = make([][]int, numcols)
	}
	for i := 0; i < numrows; i++ {
		for j := 0; j < numcols; j++ {
			cols[i][j] = []int{i, j}
		}
	}
	return cols
}
