import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def is_contains(items: List[int], target: int) -> str:
    if target in items:
        return "Yes"
    else:
        return "No"

class Test(unittest.TestCase):
    def test_is_contains_exist(self):
        expected = "Yes"
        actual = is_contains([1, 2, 3, 4, 5], 4)
        self.assertEqual(expected, actual)
    def test_is_contains_some_exist(self):
        expected = "Yes"
        actual = is_contains([1, 2, 3, 4, 5, 4], 4)
        self.assertEqual(expected, actual)
    def test_is_contains_not_exist(self):
        expected = "No"
        actual = is_contains([1, 2, 3, 4, 5], 9)
        self.assertEqual(expected, actual)

def main():
    _, target = map(int, input().split())
    items = list(map(int, input().split()))

    print(is_contains(items, int(target)))


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()


# Q
# 下記の問題をプログラミングしてみよう！
# N 個の要素からなる数列 A と、整数 B が与えられます。B が A に含まれているかどうかを判定してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N B
# A_1 A_2 ... A_N

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# B が A に含まれているなら「Yes」を、含まれていないなら「No」を出力してください。
# 末尾に改行を入れ、余計な文字、空行を含んではいけません。

# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ B ≦ 100
# ・ 1 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
# ・ i ≠ j ならば A_i ≠ A_j

# 入力例1
# 5 4
# 1 2 3 4 5

# 出力例1
# Yes

# 入力例2
# 5 6
# 1 2 3 4 5

# 出力例2
# No

# 入力例3
# 1 100
# 1

# 出力例3
# No
