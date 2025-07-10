import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# を返却
# ------------------------
def get_fibonacci(k: int) -> int:
    dp = [0] * (k + 1)
    dp[1] = 1

    for i in range(2, k + 1):
        dp[i] = dp[i - 1] + dp[i - 2]

    return dp[k]


# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestItemPriceMap(unittest.TestCase):
    def test_get_fibonacci(self):
        expected = 13
        actual = get_fibonacci(7)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    l = list(map(int, input().split()))
    k = l[0]
    # item_count = int(input()) # 1つのintの場合
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    print(get_fibonacci(k))


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step4/edit?language_uid=python3

# 下記の問題をプログラミングしてみよう！
# 整数 k が与えられます。
# 次のように定められた数列の k 項目の値を出力してください。
# ちなみに、これはフィボナッチ数列と呼ばれる有名な数列です。

# ・ a_1 = 1 
# ・ a_2 = 1 
# ・ a_n = a_{n-2} + a_{n-1} (n ≧ 3)
# (ヒント)
# 漸化式に登場する項の数が2つから3つへ増えましたが、やはりやることはこれまでと同じです。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# k

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 数列の k 項目の値を出力してください。

# また、末尾に改行を入れ、余計な文字、空行を含んではいけません。


# a_k
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ k ≦ 40

# 入力例1
# 7

# 出力例1
# 13
