import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_interval_sum_list(items: List[int], queries: List[str]) -> List[int]:
    result = []
    sums = [0] * (len(items) + 1)

    for k, v in enumerate(items):
        sums[k + 1] = sums[k] + v

    for q in queries:
        start, end = map(int, q.split())
        result.append(sums[end] - sums[start-1])
    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_interval_sum_list(self):
        expected = [10]
        actual = get_interval_sum_list([2, 4, 6, 8, 10], ["2 3"])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    items = [int(input().strip()) for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    for v in get_interval_sum_list(items, queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 長さ N の数列 A と、K 個の区間 (l_1,r_1) ... (l_K,r_K) が与えられるので、各区間についての A の区間和 A_{l_i} + ... + A_{r_i} (1 ≦ i ≦ K) を求めてください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# A_1
# ...
# A_N
# l_1 r_1
# ...
# l_K r_K


# ・1 行目では、配列 A の要素数 N と与えられる整数の数 K が与えられます。
# ・続く N 行では、配列 A の要素が A_1 から順に与えられます。
# ・続く K 行では、和を求めるのに使う区間の値 l , r が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_K


# ・i 行目に A_{l_i} + ... + A_{r_i} の和 ans_i (1 ≦ i ≦ K)を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N , K ≦ 100,000
# ・-100 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
# ・1 ≦ l_i ≦ r_i ≦ N (1 ≦ i ≦ K)

# 入力例1
# 4 2
# 16
# 88
# 10
# -65
# 2 4
# 1 2

# 出力例1
# 33
# 104

# 入力例2
# 10 5
# 82
# -37
# 40
# -72
# -24
# -54
# 57
# -6
# 42
# -24
# 8 9
# 6 9
# 2 3
# 4 4
# 1 5

# 出力例2
# 36
# 39
# 3
# -72
# -11
