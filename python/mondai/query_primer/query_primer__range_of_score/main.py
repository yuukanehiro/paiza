import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple
import math

# # ------------------------
# # itemと価格のmapを返却
# # ------------------------
def get_max_and_min(A: List[int], max_list: List[int], min_list: List[int], sqrt: int, start: int, end: int) -> Tuple[int, int]:
    start_block = start // sqrt
    end_block = end // sqrt

    max_value = -float('inf')
    min_value = float('inf')

    if start_block == end_block:
        max_value = max(A[start-1:end])
        min_value = min(A[start-1:end])
    else:
        # 左端
        end_l = (start_block + 1) * sqrt
        for i in range(start-1, end_l):
            max_value = max(max_value, A[i])
            min_value = min(min_value, A[i])
        # 中間ブロック
        for i in range(start_block + 1, end_block):
            max_value = max(max_value, max_list[i])
            min_value = min(min_value, min_list[i])
        # 右端
        start_r = end_block * sqrt
        for i in range(start_r -1, end):
            max_value = max(max_value, A[i])
            min_value = min(min_value, A[i])

    return max_value, min_value




# # ------------------------
# # テスト用コード（unittest）
# # ------------------------
# class TestItemPriceMap(unittest.TestCase):
#     def test_get_item_price_map(self):
#         expected = {"eraser": 50, "pencil": 30}
#         actual = get_item_price_map(["eraser 50", "pencil 30"])
#         self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    A = [int(input().strip()) for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    # print(items)
    # print(queries)

    size = int(math.sqrt(len(A)))
    if size * size < len(A):
        size += 1

    max_values = [None] * size
    min_values = [None] * size
    for i in range(size):
        start, end = i * size, (i + 1) * size
        max_values[i] = max(A[start:end])
        min_values[i] = min(A[start:end])

    for q in queries:
        a_l, a_r, b_l, b_r = map(int, q.split())

        a_max, a_min = get_max_and_min(A, max_values, min_values, size, a_l, a_r)
        b_max, b_min = get_max_and_min(A, max_values, min_values, size, b_l, b_r)

        if (a_max - a_min) > (b_max - b_min):
            print("A")
        elif (a_max - a_min) < (b_max - b_min):
            print("B")
        elif (a_max - a_min) == (b_max - b_min):
            print("DRAW")

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/query_primer/query_primer__range_of_score/edit?language_uid=python3

# 下記の問題をプログラミングしてみよう！
# テストの返却中に暇だった paiza 君は、また 2 人で遊ぶゲームを思いつきました。

# 「2 人はそれぞれ生徒番号 1 〜 N の全校生徒 N 人の中から生徒番号が連続するように好きな人数の生徒を選ぶ。その選んだ生徒達の得点の幅が大きい方、すなわちその生徒たちの (最高点 - 最低点) の値が大きい方が勝ち、同じだったら引き分け！」

# 「ただし、このルールだと人を多く選ぶ方が有利になってしまうから、選べる生徒の数はお互い N/2 人以下ね！」

# また審判を任されたあなたは、全ての生徒の得点を記録しておくことで、選んだ生徒たちの最小・最大の生徒番号を確認するだけで、その生徒たちの中の (最高点 - 最低点) の値をすぐに求めることができることに気付きました。

# 学校の生徒数 N と試合の数 K , 各生徒の得点 S_1 ... S_N と、
# i 番目の試合で対戦した A と B の 2 人が選んだ生徒の最小の生徒番号と最大の生徒番号が与えられるので、各試合のジャッジをしてください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# S_1
# ...
# S_N
# A_{l_1} A_{r_1} B_{l_1} B_{r_1}
# ...
# A_{l_K} A_{r_K} B_{l_K} B_{r_K}


# ・1 行目では、生徒数 N とおこなわれる試合数 K が与えられます。
# ・続く N 行のうち、i 行目では、生徒番号 i の生徒のテストの得点 S_i が与えられます。
# ・続く K 行のうち、i 行目では、i 試合目のプレイヤー A が選んだ生徒のうち、最小の生徒番号と最大の生徒番号 A_{l_i} , A_{r_i} とプレイヤー B が選んだ生徒のうち、最小の生徒番号と最大の生徒番号 B_{l_i} , B_{r_i} が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_K


# ・i 行目に i 試合目の結果を出力してください。
# A の勝ちの場合は 'A', B の勝ちの場合は 'B' , 引き分けの場合は 'DRAW' と出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N ≦ 10,000
# ・N は平方数である(ある整数 x > 0 を用いて N = x^2 と表すことができる)
# ・1 ≦ K ≦ 100,000
# ・0 ≦ S_i ≦ 100,000 (1 ≦ i ≦ N)
# ・1 ≦ Al_i ≦ Ar_i ≦ N (1 ≦ i ≦ K)
# ・1 ≦ Bl_i ≦ Br_i ≦ N (1 ≦ i ≦ K)
# ・各ゲームにおいて、プレイヤーの選ぶ生徒の数は N/2 以下であることが保証されている。

# 入力例1
# 4 2
# 1
# 3
# 2
# 4
# 1 2 2 3
# 1 2 3 4

# 出力例1
# A
# DRAW

# 入力例2
# 25 5
# 82336
# 23137
# 58263
# 78843
# 86854
# 90102
# 60652
# 35458
# 68587
# 20899
# 85950
# 20509
# 74628
# 71306
# 72676
# 70046
# 2492
# 20827
# 62047
# 4805
# 27067
# 5411
# 29873
# 37553
# 91148
# 6 11 11 14
# 13 15 12 13
# 2 3 17 19
# 1 8 13 18
# 6 10 15 16

# 出力例2
# A
# B
# B
# B
# A