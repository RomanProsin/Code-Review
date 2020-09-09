package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не верные аргументы")
		return
	}
	countArgs, _ := strconv.Atoi(os.Args[1])
	var unsortedArr []int
	for i := 2; i < countArgs; i++ {
		item, _ := strconv.Atoi(os.Args[i])
		unsortedArr = append(unsortedArr, item)
	}
	fmt.Printf("before: %v", unsortedArr)
	sortBubble(unsortedArr)
	fmt.Printf("after: %v", unsortedArr)
}

func sortBubble(arr []int) []int {
	count := len(arr)
	left := 0
	right := count - 1
	for {
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				//swap
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right -= 1
		for i := right; i > left; i-- {
			if arr[i] < arr[i-1] {
				//swap
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		left += 1
		if left >= right {
			break
		}
	}

	return arr
}
