/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
*/
package arr_string

func P121(prices []int) int {
	result := 0

	min := prices[0]
	for i := range prices {
		if prices[i] < min {
			min = prices[i]
		} else {
			result = max121(prices[i]-min, result)
		}
	}

	return result
}
func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}
