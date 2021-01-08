function maxProfit(prices: number[]): number {
    const n = prices.length;
    let buy1 = -prices[0], buy2 = -prices[0];
    let sell1 = 0, sell2 = 0;
    for (let i = 1; i < n; i++) {
        buy1 = Math.max(buy1, -prices[i]);
        sell1 = Math.max(sell1, buy1 + prices[i]);
        buy2 = Math.max(buy2, sell1 - prices[i]);
        sell2 = Math.max(sell2, buy2 + prices[i]);
    }
    return sell2;
};

/**
 * 123. 买卖股票的最佳时机 III
 */
test('leetcode_123', () => {
    console.log(maxProfit([3, 3, 5, 0, 0, 3, 1, 4]))
});