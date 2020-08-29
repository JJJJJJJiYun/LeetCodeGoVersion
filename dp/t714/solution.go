package t714

import "math"

// 给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
// 你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
// 返回获得利润的最大值。
// 注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。

// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price[i])
// dp[i][1] = max(dp[i-1][1], dp[i-1][0]-price[i]-fee)
func maxProfit(prices []int, fee int) int {
	notPossess, possess := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		lastNotPossess := notPossess
		notPossess = max(notPossess, possess+prices[i])
		possess = max(possess, lastNotPossess-prices[i]-fee)
	}
	return notPossess
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}