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

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    x, d, k = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    dp = get_arithmetic_progression(x, d, k)
    print(dp[k-1])

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# # Q
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step0/edit?language_uid=python3

# 2項間漸化式 1 Python3編（paizaランク C 相当）
# 問題にチャレンジして、ユーザー同士で解答を教え合ったり、コードを公開してみよう！

# シェア用URL:
# https://paiza.jp/works/mondai/dp_primer/dp_primer_recursive_formula_step0
# 問題文のURLをコピーする
#  下記の問題をプログラミングしてみよう！
# (はじめに)
# このメニューでは、動的計画法 (Dynamic Programming, 以下 DP と表記します) について扱います。

# DP は、一言でいうと「問題を部分問題に分割し、部分問題の答えを記録しながら、それらを利用することによって元の問題の答えを得る手法」です。

# 問題をどのように分割するか、部分問題の答えをどのように利用するかなどは問題により異なります。このメニューを通してさまざまなDPの問題に触れ、そのノウハウを身につけていきましょう。

# まずは、早速問題を見てみましょう。

# (問題)
# 整数 x, d, k が与えられます。
# 次のように定められた数列の k 項目の値を出力してください。

# ・ a_1 = x
# ・ a_n = a_{n-1} + d (n ≧ 2)
# (ヒント)
# 等差数列の公式を使えばDPを使わずとも答えを求めることができますが、練習だと思ってDPで解いてみましょう。

# DPを考える際には、まず漸化式を考えるとよいです。漸化式は、数列の各項をその前の項を用いて記述した式です。問題で与えられている a_n = a_{n-1} + d という式がこの問題における漸化式となっています。

# では、実際にこの問題を解いてみましょう。最終的に求めたいのは a_k です。a_1 ~ a_{k-1} がわかっているとして、a_k がどうなるかを考えてみましょう (a_1 ~ a_{k-1} が、「はじめに」の部分で述べた"部分問題"に相当しています) 。a_{k-1} がわかっているので、a_k = a_{k-1} + d とすればよいですね。今回は a_1 が x であることがわかっているので、順に a_2, a_3, a_4, ... を計算していけば a_k が求まることがわかるかと思います。

# 以下の疑似コードを参考にして、あなたの得意な言語で実装してみましょう。

# a[1] <- x

# for i = 2 to k
#     a[i] <- a[i-1] + d

# print a[k]
# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# x d k

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 数列の k 項目の値を出力してください。

# また、末尾に改行を入れ、余計な文字、空行を含んではいけません。


# a_k
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ -1,000 ≦ x ≦ 1,000

# ・ -1,000 ≦ d ≦ 1,000

# ・ 1 ≦ k ≦ 1,000

# 入力例1
# 0 7 9

# 出力例1
# 56