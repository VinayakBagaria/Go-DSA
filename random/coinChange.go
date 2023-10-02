// https://leetcode.com/problems/coin-change/description/
package random

import "math"

func coinChangeTopDown(coins []int, amount int) int {
    dp := make(map[int]int)
    
    var dfs func(int) int
    dfs = func(required int) int {
        if required == 0 {
            return 0
        }

        if required < 0 {
            return math.MaxInt
        }

        if val, ok := dp[required]; ok {
            return val
        }

        result := math.MaxInt
        for _, c := range coins {
            subproblem := dfs(required - c)
            if subproblem != math.MaxInt {
                result = min(result, 1 + subproblem)
            }
        }
        dp[required] = result
        return result
    }

    result := dfs(amount)
    if result == math.MaxInt {
        return -1
    }
    return result
}


func coinChangeBottomUp(coins []int, amount int) int {
	if amount == 0 {
        return 0
    }

    dp := make([]int, amount + 1)
    for i := range dp {
        dp[i] = math.MaxInt
    }
    for _, c := range coins {
        if c <= amount {
            dp[c] = 1
        }
    }

    for i := 1; i <= amount; i++ {
        for _, c := range coins {
            if i - c >= 0 && dp[i - c] != math.MaxInt {
                dp[i] = min(dp[i], 1 + dp[i - c])
            }
        }
    }

    ans := dp[amount]
    if ans == math.MaxInt {
        return -1
    }
    return ans
}