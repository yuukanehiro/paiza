import sys
import unittest
from collections import defaultdict
from typing import List


def get_item_seen_results(items: List[int]) -> List[str]:
    seen = defaultdict(bool)
    results = []

    seen[items[0]] = True
    for i in range(1, len(items)):
        if seen[items[i]]:
            results.append("Yes")
        else:
            results.append("No")
            seen[items[i]] = True
    
    return results


class TestGetItemSeenResults(unittest.TestCase):
    def test_get_item_seen_results_exists1(self):
        expected = ["No", "No", "Yes", "No", "Yes", "Yes", "No", "Yes"]
        actual = get_item_seen_results([1, 2, 3, 2, 5, 3, 3, 10, 2])
        self.assertEqual(expected, actual)
    def test_get_item_seen_results_exists2(self):
        expected = ["Yes"]
        actual = get_item_seen_results([1, 1])
        self.assertEqual(expected, actual)
    def test_get_item_seen_results_some_exists(self):
        expected = ["No", "Yes", "Yes", "Yes", "No", "No"]
        actual = get_item_seen_results([1, 2, 1, 1, 2, 3, 4])
        self.assertEqual(expected, actual)
    def test_get_item_seen_results_not_exists(self):
        expected = ["No", "No"]
        actual = get_item_seen_results([1, 2, 3])
        self.assertEqual(expected, actual)

def main():
    _ = map(int, input().split())
    items = list(map(int, input().split()))

    for v in get_item_seen_results(items):
        print(v)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# N 個の要素からなる数列 A が与えられます。2 ≦ i ≦ N の各 i に対して、A_i と同じ値が A_1 から A_{i-1} の間にあるかどうかを判定してください。

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

# ・ 2 ≦ N ≦ 100
# ・ 1 ≦ A_i ≦ 10,000 (1 ≦ i ≦ N)

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
# 1 1

# 出力例2
# Yes