package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	output := make([]bool, len(candies))
	maxNumber := candies[0]
	for _, v := range candies[1:] {
		if maxNumber < v {
			maxNumber = v
		}
	}

	for i := 0; i < len(candies); i++ {
		output[i] = candies[i]+extraCandies >= maxNumber
	}
	return output
}
