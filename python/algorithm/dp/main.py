def knapsack(weights, values, W):
    n = len(weights)
    dp = [[0] * (W + 1) for _ in range(n + 1)]

    for i in range(n):
        for w in range(W + 1):
            if w < weights[i]:
                dp[i + 1][w] = dp[i][w] # 入れらない
            else:
                # 入れない場合 vs 入れる場合
                # 最大の価値の場合を表に入れる
                dp[i + 1][w] = max(dp[i][w], dp[i][w - weights[i]] + values[i])
    return dp[i][W]

if __name__ == '__main__':
    weights = [2, 1, 3]
    values = [3, 2, 4]
    W = 5

    print(knapsack(weights, values, W)) # Output: 5
