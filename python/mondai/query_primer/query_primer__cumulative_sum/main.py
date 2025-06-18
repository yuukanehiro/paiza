import sys
import unittest
from typing import List


def get_cumulative_sum_list(items: List[int]) -> List[int]:
    sums = [0] * (len(items) + 1)

    for i, v in enumerate(items):
        sums[i + 1] = sums[i] + v

    return sums


class TestItemPriceMap(unittest.TestCase):
    def test_get_cumulative_sum_list(self):
        expected = [
            0,
            69,
            81,
            109
        ]
        input_items = [
            69,
            12,
            28
        ]
        actual = get_cumulative_sum_list(input_items)
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    items = [int(input().strip()) for _ in range(item_count)]
    queries = [int(input().strip()) for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    sums = get_cumulative_sum_list(items)

    for q in queries:
        print(sums[q])


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 長さ N の数列 A と、K 個の整数 Q_1 ... Q_K が与えられるので、各整数 Q_i (1 ≦ i ≦ K) について A_1 ... A_{Q_i} の和を求めてください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# A_1
# ...
# A_N
# Q_1
# ...
# Q_K


# ・1 行目では、配列 A の要素数 N と与えられる整数の数 K が与えられます。
# ・続く N 行では、配列 A の要素が A_1 から順に与えられます。
# ・続く K 行では、整数 Q_1 ... Q_K が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_K


# ・i 行目に A_1 ... A_{Q_i} の和 ans_i (1 ≦ i ≦ K)を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N , K ≦ 100,000
# ・-100 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
# ・1 ≦ Q_i ≦ N (1 ≦ i ≦ K)

# 入力例1
# 3 1
# 69
# 12
# 28
# 3

# 出力例1
# 109

# 入力例2
# 10 3
# 45
# 74
# -94
# 68
# -63
# 19
# -47
# -69
# 38
# 60
# 9
# 5
# 5

# 出力例2
# -29
# 30
# 30
