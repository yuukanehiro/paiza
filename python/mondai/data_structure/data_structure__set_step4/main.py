import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_item_seen_results(items: List[int]) -> List[str]:
    seen = defaultdict(bool)
    results = []

    seen[items[0]] = True # 重要
    for v in range(1, len(items)):
        if seen[items[v]]:
            results.append("Yes")
        else:
            results.append("No")
            seen[items[v]] = True

    return results

class TestItemPriceMap(unittest.TestCase):
    def test_get_item_seen_results_case1(self):
        expected = ["No", "No", "Yes", "No", "Yes", "Yes", "No", "Yes"]
        actual = get_item_seen_results([1, 2, 3, 2, 5, 3, 3, 10, 2])
        self.assertEqual(expected, actual)
    def test_get_item_seen_results_case2(self):
        expected = ["Yes"]
        actual = get_item_seen_results(["1000000000", "1000000000"])
        self.assertEqual(expected, actual)

def main():
    _ = map(int, input().split())
    items = list(map(int, input().split()))

    for a in get_item_seen_results(items):
        print(a)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 重複の判定 2 Python3編（paizaランク C 相当）
# 問題にチャレンジして、ユーザー同士で解答を教え合ったり、コードを公開してみよう！

# シェア用URL:
# https://paiza.jp/works/mondai/data_structure/data_structure__set_step4
# 問題文のURLをコピーする
#  下記の問題をプログラミングしてみよう！
# N 個の要素からなる数列 A が与えられます。2 ≦ i ≦ N の各 i に対して、A_i と同じ値が A_1 から A_{i-1} の間にあるかどうかを判定してください。ただし、A_i は非常に大きくなることがあります。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# A_1 A_2 ... A_N

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# N-1 行出力してください。i (1 ≦ i ≦ N-1) 行目には、A_{i+1} と同じ値が A_1 から A_i の間にあるならば「Yes」を、ないならば「No」を 1 行に出力してください。
# 末尾に改行を入れ、余計な文字、空行を含んではいけません。

# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ A_i ≦ 1,000,000,000

# 入力例1
# 9
# 1 2 3 2 5 3 3 10 2

# 出力例1
# No
# No
# Yes
# No
# Yes
# Yes
# No
# Yes

# 入力例2
# 2
# 1000000000 1000000000

# 出力例2
# Yes