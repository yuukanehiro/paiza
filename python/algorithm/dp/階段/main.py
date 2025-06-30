import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# n段目に到達するパターンの総数を返却
# ------------------------
def count_ways(n):
    dp = [0] * (n + 1)
    dp[0] = 1 # 0段目は1通りとする

    for i in range(1, n + 1):
        if i - 1 >= 0:
            dp[i] += dp[i - 1]
        if i - 2 >= 0:
            dp[i] += dp[i - 2]
        if i - 3 >= 0:
            dp[i] += dp[i - 3]

    return dp[n]

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_count_ways(self):
        expected = 24
        actual = count_ways(6)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    # 「始めに階段の0段目にいる。あなたは足が長いので、1段上にいく、2段上にいく、
    # 3段上にいくの3パターンの登り方がある。6段目に登る方法は何通りか？」

    print(count_ways(6)) # Output: 24


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()