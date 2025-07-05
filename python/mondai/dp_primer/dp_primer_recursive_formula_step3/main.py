import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# を返却
# ------------------------
def get_arithmetic_progression_odd_even(x: int, d_1: int, d_2: int, k: int) -> List[int]:
    dp = [0] * (k + 1)
    dp[1] = x

    for i in range(2, k + 1):
        if i % 2 == 0:
            dp[i] = dp[i - 1] + d_2
        else:
            dp[i] = dp[i - 1] + d_1

    return dp

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_get_arithmetic_progression_odd_even(self):
        expected = [0, 3, -1, 6, 2, 9, 5, 12, 8, 15, 11]
        actual = get_arithmetic_progression_odd_even(3, 7, -4, 10)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    x, d_1, d_2 = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    query_count = int(input())
    queries = [int(input().strip()) for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    dp = get_arithmetic_progression_odd_even(x, d_1, d_2, max(queries))

    for q in queries:
        print(dp[q])

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step3/edit?language_uid=python3

# 下記の問題をプログラミングしてみよう！
# 整数 x, d_1, d_2, Q と Q 個の整数 k_1, k_2, ... , k_Q が与えられます。

# 次のように定められた数列の k_i 項目の値を順に出力してください。


# ・ a_1 = x 
# ・ a_n = a_{n-1} + d_1 (n が奇数のとき、n ≧ 3) 
# ・ a_n = a_{n-1} + d_2 (n が偶数のとき)
# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# x d_1 d_2
# Q
# k_1
# k_2
# ...
# k_Q


# ・ 1行目では、数列の初項 x と公差 d_1, d_2 が与えられます。

# ・ 2行目では、3行目以降で与えられる入力の行数 Q が与えられます。

# ・ 続く Q 行のうち i 行目では、k_i が与えられます。


# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# Q 行出力してください。

# i 行目には、数列の k_i 項目の値を出力してください。

# また、末尾に改行を入れ、余計な文字、空行を含んではいけません。


# a_{k_1}
# a_{k_2}
# ...
# a_{k_Q}
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ -1,000 ≦ x ≦ 1,000

# ・ -1,000 ≦ d_1 ≦ 1,000

# ・ -1,000 ≦ d_2 ≦ 1,000

# ・ 1 ≦ Q ≦ 1,000

# ・ 1 ≦ k_i ≦ 1,000 (1 ≦ i ≦ Q)

# 入力例1
# 3 7 -4
# 5
# 1
# 2
# 3
# 4
# 10

# 出力例1
# 3
# -1
# 6
# 2
# 11