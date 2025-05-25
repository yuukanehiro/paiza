import unittest
from collections import defaultdict

def fibonacci(n):
    """
    n番目のフィボナッチ数列の値を返却する
    // フィボナッチ数列 1, 1, 2, 3, 5, 8, 13, 21, 34, ...
    """
    dp = [0] * (n+1) # dp[0]は0で、フィボナッチ数列はdp[1]の1から始まるからn+1にする
    dp[1] = 1
    for i in range(2, n+1):
        dp[i] = dp[i-1] + dp[i-2]
    return dp[n]

if __name__ == '__main__':
    print(fibo(10)) # Output: 55
