import sys
import unittest
from collections import defaultdict
from typing import List


def get_count(item_lines: List[str], query: str) -> int:
    count = 0
    seen = False
    for item_line in item_lines:
        count += 1
        if item_line == query:
            seen = True
            return count
    if seen == False:
        return -1

class Test(unittest.TestCase):
    def test_get_count_exist(self):
        expected = 2
        actual = get_count(["a", "b", "c"], "b")
        self.assertEqual(expected, actual)
    def test_get_count_some_exist(self):
        expected = 2
        actual = get_count(["a", "b", "c", "b"], "b")
        self.assertEqual(expected, actual)
    def test_get_count_not_exist(self):
        expected = -1
        actual = get_count(["a", "b", "c"], "x")
        self.assertEqual(expected, actual)

def main():
    item_count, query_count = map(int, input().split())
    item_lines = [input() for _ in range(item_count)]
    query_lines = [input() for _ in range(query_count)]

    for query_line in query_lines:
        print(get_count(item_lines, query_line))

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
#  下記の問題をプログラミングしてみよう！
# N 個の文字列 S_1, ... , S_N と、Q 個の文字列 T_1, ... , T_Q が与えられます。各 T_i について、以下の処理を行ってください。

# ・ S_j == T_i を満たす最小の j を出力する。ただし、そのような j が存在しない場合は -1 を出力する。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N Q
# S_1
# S_2
# ...
# S_N
# T_1
# T_2
# ...
# T_Q

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# Q 行出力してください。i 行目には、S_j == T_i を満たす最小の j を出力してください。ただし、そのような j が存在しない場合は -1 を出力してください。


# j_1
# ...
# j_Q


# 末尾に改行を入れ、余計な文字、空行を含んではいけません。
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ Q ≦ 100
# ・ S_i, T_j は英小文字からなる1文字以上3文字以下の文字列 (1 ≦ i ≦ N,1 ≦ j ≦ Q)

# 入力例1
# 3 2
# a
# b
# c
# b
# d

# 出力例1
# 2
# -1

# 入力例2
# 6 2
# pai
# za
# p
# pa
# pai
# za
# za
# pai

# 出力例2
# 2
# 1
