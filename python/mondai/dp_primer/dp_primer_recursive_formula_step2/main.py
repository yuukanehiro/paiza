import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 等差数列のdpを返却
# ------------------------
def get_arithmetic_progression_odd_even(x: int, d_1: int, d_2: int, k: int) -> List[int]:
    # 1-basedの形にする
    dp = [0] * (k + 1)
    dp[1] = x

    for i in range(2, k + 1):
        # 偶数
        if i % 2 == 0:
            dp[i] = dp[i - 1] + d_2
        # 奇数
        else:
            dp[i] = dp[i - 1] + d_1
    return dp

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_get_arithmetic_progression_odd_even(self):
        expected = [0, 5, 15, 8, 18, 11]
        actual = get_arithmetic_progression_odd_even(5, -7, 10, 5)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    x, d_1, d_2, k = map(int, input().split())

    dp = get_arithmetic_progression_odd_even(x, d_1, d_2, k)
    print(dp[k])


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step2/edit?language_uid=python3&t=b84fa977ad0a34849b2a26cb52c5ed1e

# 下記の問題をプログラミングしてみよう！
# 整数 x, d_1, d_2, k が与えられます。
# 次のように定められた数列の k 項目の値を出力してください。

# ・ a_1 = x 
# ・ a_n = a_{n-1} + d_1 (n が奇数のとき、n ≧ 3) 
# ・ a_n = a_{n-1} + d_2 (n が偶数のとき)
# (ヒント)
# 添字の偶奇によって漸化式の形が変わっていますが、やることはこれまでと同じです。a_1 ~ a_{k-1} が求まっているとして、a_k をどのように計算すればよいかを考えてみましょう。計算するときに、添字の偶奇による場合分けを行えばよいです。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# x d_1 d_2 k

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 数列の k 項目の値を出力してください。

# また、末尾に改行を入れ、余計な文字、空行を含んではいけません。


# a_k
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ -1,000 ≦ x ≦ 1,000

# ・ -1,000 ≦ d_1 ≦ 1,000

# ・ -1,000 ≦ d_2 ≦ 1,000

# ・ 1 ≦ k ≦ 1,000

# 入力例1
# 5 -7 10 5

# 出力例1
# 11