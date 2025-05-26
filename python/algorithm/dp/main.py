def knapsack(weights, values, W):
    n = len(weights)
    dp = [[0] * (W + 1) for _ in range(n + 1)]

    for i in range(n): # ← iは「何番目のアイテムを使うか」のインデックス
        for w in range(W + 1):
            if w < weights[i]: # 重さが足りないなら入れない
                # 例)
                # i=2 のときは「アイテム2をこれから使うかどうか考える」
                # dp[1][w] ...「アイテム0までを使った状態」
                # dp[2][w] ...「アイテム0と1を使った場合の重さ w 以内で得られる最大の価値」

                # dp[i+1][w] ... i+1 個目までのアイテム（つまり「i番目のアイテム」まで）を見たときの、重さ w のときの最大価値
                dp[i + 1][w] = dp[i][w] # 入れらないので、前の最大価値を代入
            else:
                # 入れない場合 vs 入れる場合 // dp[i][w - weights[i]] ... 入れる場合の許容スペースが空いている場合の最大価値
                # を比べて最大価値を代入
                dp[i + 1][w] = max(dp[i][w], dp[i][w - weights[i]] + values[i])
    return dp[i][W]

if __name__ == '__main__':
    weights = [2, 1, 3]
    values = [3, 2, 4]
    W = 5

    print(knapsack(weights, values, W)) # Output: 5
