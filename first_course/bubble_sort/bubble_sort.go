package main

import (
	"fmt"
)

func main() {

	fmt.Println(`To perform bubble sort you should input the sequence up to 10 integers, let's Go:`)

	sliceToSort := getUserInput()
	bubbleSort(sliceToSort)
	fmt.Printf("Sorted sequence of your input integer is:\n %v", sliceToSort)

}

func getUserInput() []int {
	sliceLength := 10
	sliceToSort := make([]int, 0)
	for i := 0; i < sliceLength; i++ {
		fmt.Printf("%d int: ", i+1)
		var inputInt int
		_, err := fmt.Scan(&inputInt)
		if err != nil {
			fmt.Println("Aint int. Try again.")
			break
		}
		sliceToSort = append(sliceToSort, inputInt)
	}
	return sliceToSort
}

func bubbleSort(sliceToSort []int) {
	swapFlag := 0
	for index := 0; index < len(sliceToSort)-1; index++ {
		if sliceToSort[index] > sliceToSort[index+1] {
			swap(sliceToSort, index)
			swapFlag++
		}
	}
	if swapFlag > 0 {
		bubbleSort(sliceToSort)
	}
}

func swap(numbersSlice []int, indexToSwap int) {
	tempSwap := numbersSlice[indexToSwap]
	numbersSlice[indexToSwap] = numbersSlice[indexToSwap+1]
	numbersSlice[indexToSwap+1] = tempSwap
}
