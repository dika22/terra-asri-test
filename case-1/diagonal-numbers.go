package main

import "fmt"


func main()  {
	fmt.Println(calculateDiagonalSum(3))
	// fmt.Println(calculateDiagonalSum(5))
	// fmt.Println(calculateDiagonalSum(7))
}


func calculateDiagonalSum(numbers int) int {
	n := numbers
    length := (n * n) + 1
	islands := make([][]int, n)
	doSwitch := false
	row := 0
	column := n
	loopCount := n

	for i := range islands {
        islands[i] = make([]int, n)
        // for j := range islands[i] {
        //     islands[i][j] = n
        // }

		// fmt.Println(islands)
    }

	for {
		if doSwitch {
		} else {
			for i := 0; i < loopCount; i++ {
				length -= 1;
				column += 1;
				islands[row][column] = length;
				fmt.Println(islands)
			}
			// if (loopCount > 0) {
			// 	// loop row
			// 	for i := 0; i < loopCount; i++ {
			// 		length -= 1
			// 		column -= 1
			// 		islands[row][column] = length
			// 	}

			// 	loopCount = loopCount - 1
			// 	// loop column
			// 	for i := 0; i < loopCount; i++ {
			// 		length -= 1
			// 		row += 1
			// 		islands[row][column] = length
			// 	}
			// }
		}

		doSwitch = !doSwitch;
		if loopCount > 0 {
			break
		}

	} 

	count := 0

	return count
}