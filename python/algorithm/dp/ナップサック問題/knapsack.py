def knapsack(weights, values, W):
    n = len(weights)
    dp = [[0] * (W + 1) for _ in range(n + 1)]

    for i in range(1, n + 1): # i=1から始める。iは「何番目のアイテムを使うか」のインデックス
        for w in range(W + 1):
            if w >= weights[i - 1]:
                # 容量が足りている場合
                # 入れない場合 vs 入れる場合 // dp[i][w - weights[i]] ... 入れる場合の許容スペースが空いている場合の最大価値
                # を比べて最大価値を代入
                dp[i][w] = max(
                    dp[i - 1][w],
                    dp[i - 1][w - weights[i - 1]] + values[i - 1]
                )     
            else:
                # 容量が足りない場合
                dp[i][w] = dp[i - 1][w] # 入れらないので、前の最大価値を代入
    return dp[n][W]

if __name__ == '__main__':
    weights = [2, 1, 3]
    values = [3, 2, 4]
    W = 5

    print(knapsack(weights, values, W)) # Output: 7
