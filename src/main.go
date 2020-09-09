package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := flag.String("data", "", "data for sort (ex. 3,2,1")
	flag.Parse()
	arrayData := strings.Split(*data, ",")
	var dataToSort []int
	for id := range arrayData {
		i, err := strconv.Atoi(arrayData[id])
		if err != nil {
			fmt.Printf("Data %s is not int", arrayData[id])
			os.Exit(2)
		}
		dataToSort = append(dataToSort, i)
	}
	fmt.Println(sortBubble(dataToSort))
}

func sortBubble(data []int) []int {
	length := len(data)
	if length == 0 {
		return nil
	}
	res := make([]int, length)
	copy(res, data)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return res
}
