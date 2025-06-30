import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 等差数列のdpを返却
# ------------------------
def get_arithmetic_progression(x: int, d: int, k: int) -> List[int]:
    dp = [0] * k
    dp[0] = x
    for i in range(1, k):
        dp[i] = dp[i -1] + d
    
    return dp

# ------------------------
# を返却
# ------------------------
def get_result_list(dp: List[int], queries: List[int]) -> List[int]:
    result = []

    for q in queries:
        result.append(dp[q-1])

    return result

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_get_arithmetic_progression(self):
        expected = [
            0,
            7,
            14,
            21,
            28,
            35,
            42,
            49,
            56
        ]
        actual = get_arithmetic_progression(0, 7, 9)
        self.assertEqual(expected, actual)

    def test_get_result_list(self):
        input_dp = [
            0,
            7,
            14,
            21,
            28
        ]
        input_queries = [
            1,
            2,
            3,
            4,
            5
        ]
        expected = [
            0,
            7,
            14,
            21,
            28
        ]
        actual = get_result_list(input_dp, input_queries)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    x, d = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    query_count = int(input())

    dp = get_arithmetic_progression(x, d, query_count)

    qeueries = [int(input().strip()) for _ in range(query_count)]

    for v in get_result_list(dp, qeueries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step1/edit?language_uid=python3

# 2項間漸化式 2 Python3編（paizaランク C 相当）
# 問題にチャレンジして、ユーザー同士で解答を教え合ったり、コードを公開してみよう！

# シェア用URL:
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step1
# 問題文のURLをコピーする
#  下記の問題をプログラミングしてみよう！
# 整数 x, d, Q と Q 個の整数 k_1, k_2, ... , k_Q が与えられます。

# 次のように定められた数列の k_i 項目の値を順に出力してください。


# ・ a_1 = x
# ・ a_n = a_{n-1} + d (n ≧ 2)
# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# x d
# Q
# k_1
# k_2
# ...
# k_Q


# ・ 1行目では、数列の初項 x と公差 d が半角スペース区切りで与えられます。

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

# ・ -1,000 ≦ d ≦ 1,000

# ・ 1 ≦ Q ≦ 1,000

# ・ 1 ≦ k_i ≦ 1,000 (1 ≦ i ≦ Q)

# 入力例1
# 0 7
# 5
# 1
# 2
# 3
# 4
# 5

# 出力例1
# 0
# 7
# 14
# 21
# 28