package coinchange

import "fmt"

// claude notes:
// 0   1   2   3   4   5   6   7   8   9   10
// []      1   0   0   0   0   0   0   0   0   0   0
// [1]     1   1   1   1   1   1   1   1   1   1   1
// [1,2]   1   1   2   2   3   3   4   4   5   5   6
// [1,2,3] 1   1   2   3   4   5   7   8   10  12  14
// [1,2,3,5] 1 1   2   3   4   6   9   11  15  18  ?
// reference implementation:
// # Process each coin
// for coin in coins:
//     # Update all amounts we can make with this coin
//     for j in range(coin, amount + 1):
//         dp[j] += dp[j - coin]

// return dp[amount]
func makeChange(coins []int, target int) (numberOfWays int) {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, coin := range coins {
		if coin > target {
			continue
		}
		for j := range dp[coin:] {
			dp[j+coin] += dp[j]
		}
	}
	fmt.Println(dp)
	return dp[target]
	// This is the implementation that passed the tests in hackerrank.
	// Write your code here
	// this is basically finding the divisors, for each coin less than
	// the target
	// target := n
	// coins := c
	// dp := make([]int64, target+1)
	// dp[0] = 1
	// for _, coin := range coins {
	//     if coin > int64(target) {
	//         continue
	//     }
	//     for j := range dp[coin:] {
	//         dp[int64(j)+coin] += dp[j]
	//     }
	// }
	// fmt.Println(dp)
	// return dp[int64(target)]
}
