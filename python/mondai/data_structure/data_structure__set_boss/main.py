import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_sorted_hash_set(a_items: List[int], b_items: List[int]) -> List[int]:
    hashSet = set()

    for v in a_items:
        hashSet.add(v)
    
    for v in b_items:
        hashSet.add(v)

    return sorted(hashSet)


class TestItemPriceMap(unittest.TestCase):
    def test_get_sorted_hash_set_case1(self):
        expected = [1, 2, 3, 4, 5]
        actual = get_sorted_hash_set([1, 2, 3], [3, 4, 5])
        self.assertEqual(expected, actual)
    def test_get_sorted_hash_set_case2(self):
        expected = [1]
        actual = get_sorted_hash_set([1, 1, 1], [1, 1, 1])
        self.assertEqual(expected, actual)
    def test_get_sorted_hash_set_case3(self):
        expected = [7, 8, 9]
        actual = get_sorted_hash_set([9, 8, 7], [7, 9, 8])
        self.assertEqual(expected, actual)

def main():
    _ = map(int, input().split())
    a_items = list(map(int, input().split()))
    b_items = list(map(int, input().split()))

    sortedHashSet = get_sorted_hash_set(a_items, b_items)
    sortedHashSetStr = [str(v) for v in sortedHashSet]

    print(" ".join(sortedHashSetStr))


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# N 個の要素からなる数列 A, B が与えられます。A または B に含まれる値をすべて列挙し、重複を取り除いて昇順に出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# A_1 A_2 ... A_N
# B_1 B_2 ... B_N

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# 答えとなる数列 C を半角スペース区切りで1行に出力してください。
# 末尾に改行を入れ、余計な文字、空行を含んではいけません。


# C_1 C_2 ...
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 入力はすべて整数
# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ A_i, B_i ≦ 1,000,000,000 (1 ≦ i ≦ N)

# 入力例1
# 3
# 1 2 3
# 3 4 5

# 出力例1
# 1 2 3 4 5

# 入力例2
# 3
# 1 1 1
# 1 1 1

# 出力例2
# 1

# 入力例3
# 3
# 9 8 7
# 7 9 8

# 出力例3
# 7 8 9