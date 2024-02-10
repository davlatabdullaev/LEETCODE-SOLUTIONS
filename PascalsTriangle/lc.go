package pascalsarray

func Generate(numRows int) [][]int {

	array := make([][]int, numRows)
	if numRows == 0 {
		return array
	}

	array[0] = []int{1}

	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		nRow := array[i-1]

		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = nRow[j-1] + nRow[j]
		}

		array[i] = row
	}

	return array

}
