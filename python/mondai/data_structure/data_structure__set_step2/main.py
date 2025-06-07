import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_sorted_unique_items(items: List[str]) -> list[str]:
    unique_items = set(items)
    return sorted(unique_items)

class TestGetSortedUniqueItems(unittest.TestCase):
    def test_get_sorted_unique_items(self):
        expected = ["1", "2", "3", "4", "5"]
        actual = get_sorted_unique_items(["1", "2", "3", "3", "4", "5"])
        self.assertEqual(expected, actual)

def main():
    _ = map(int, input())
    items = input().split()

    for v in get_sorted_unique_items(items):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# N 個の要素からなる数列 A が与えられます。数列 A は昇順にソートされています。A の重複した要素を取り除いて昇順に出力してください。

# 入力される値
# N
# A_1 A_2 ... A_N

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# A の重複した要素を取り除き、半角スペース区切りで昇順に1行で出力してください。
# 末尾に改行を入れ、余計な文字、空行を含んではいけません。

# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
# ・ i < j ならば A_i ≦ A_j

# 入力例1
# 6
# 1 2 3 3 4 5

# 出力例1
# 1 2 3 4 5

# 入力例2
# 3
# 1 1 1

# 出力例2
# 1

# 入力例3
# 3
# 1 2 3

# 出力例3
# 1 2 3