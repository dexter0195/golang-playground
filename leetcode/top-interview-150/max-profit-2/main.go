package main

import "fmt"

func maxProfit(prices []int) int {
	maximumProfit := 0
	// iterate over the array and keep track of the buyDay
	// we start by always buying, as we can also resell on the same day if we see that this is not going well
	// profits := make([]int, len(prices)/2)
	// profits[0] = prices[1] - prices[0]
	for i := 0; i < len(prices)-1; i++ {
		profit := prices[i+1] - prices[i]
		if profit > 0 {
			maximumProfit += profit
		}
	}
	return maximumProfit

}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
