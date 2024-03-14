package main

import "fmt"


func main()  {

	// schedules := [][][]int{
	// 	{{8, 10}}, {{12, 14}}, {{9,11}},
	// }

	schedules := [][][]int{
		{{9, 12}, {14, 16}},
		{{10, 12},{15, 17}},
		{{11, 13},{16, 18}},
	}

	// Print result
	fmt.Println(findCommonSlot(schedules))
}

func findCommonSlot(schedules [][][]int) interface{} {
	// Logic Here
	result := schedules[0]
	for i := 1; i < len(schedules); i++ { // loop count of player (wizard) schedules
		scheduleToCompare := [][]int{}
		for j := 0; j < len(result); j++ { // loop by player schedule
			for k := 0; k < len(schedules[i]); k++ { // loop by next player schedule
				from := result[j][0]
				to   := result[j][1]
				fromN := schedules[i][k][0]
				toN   := schedules[i][k][1]

				f1 := from
				if (from <= fromN) {
					 f1 = fromN
				}
				isFromValid := f1 <= toN && f1 <= to
				
				t1 := to
				if toN <= to {
					t1 = toN
				}
				isToValid := t1 >= fromN && t1 >= from

				if isFromValid && isToValid { 
					schedule := []int{}
					schedule = append(schedule, f1)
					schedule = append(schedule, t1)
					scheduleToCompare = append(scheduleToCompare, schedule)
				}
			}
		}
		result = scheduleToCompare
	}

	if len(result) == 0 {
		msg := []string{"No common slot available"}
		return msg
	}else{
		return result[0]
	}
}