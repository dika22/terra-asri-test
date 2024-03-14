package main

import "fmt"

func main()  {
	// selfNumber(5000)
	selfNumber(100)
}

func selfNumber(numbers int) int {
	sn := []int{}
	shortOfNumbers := make([]bool, numbers+1)
	for i := 1; i <= numbers; i++ {
		 n := i + getNumber(i)
		 if n <= numbers {
			shortOfNumbers[n] = true
		 }
	}

	for i, v := range shortOfNumbers {
		if !v && i > 0{
			sn = append(sn, i)
		}
	}

	fmt.Println(sn)
	return 0
}

func getNumber(i int) int {
	sum := 0
    for i > 0 {
        sum += i % 10
		i /= 10
    }
    return sum
}