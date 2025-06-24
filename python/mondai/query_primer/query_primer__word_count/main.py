import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple


def get_item_price_map(items: List[str]) -> Dict[str, int]:
    item_price_map: Dict[str, int] = {}
    for line in items:
        name, price = line.split()
        item_price_map[name] = int(price)
    return item_price_map


class TestItemPriceMap(unittest.TestCase):
    def test_get_item_price_map(self):
        expected = {"eraser": 50, "pencil": 30}
        actual = get_item_price_map(["eraser 50", "pencil 30"])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    A = [0] + [int(input().strip()) for i in range(item_count)]
    # items = [list(map(int, input().strip().split())) for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    sum_array = [0] * len(A)
    for i, _ in enumerate(A):
        sum_array[i] = A[i] + sum_array[i - 1]

    for q in queries:
        a_s, a_e, b_s, b_e = map(int, q.split())
        limit_page = item_count / 3

        a_faul = False
        b_faul = False
        if a_e - a_s + 1 >= limit_page:
            a_faul = True
        if b_e - b_s + 1 >= limit_page:
            b_faul = True

        if a_faul and b_faul:
            print("DRAW")
            continue
        elif not a_faul and b_faul:
            print("A")
            continue
        elif not b_faul and a_faul:
            print("B")
            continue

        a_sum = sum_array[a_e] - sum_array[a_s - 1]
        b_sum = sum_array[b_e] - sum_array[b_s - 1]

        if a_sum == b_sum:
            print("DRAW")
            continue
        elif a_sum > b_sum:
            print("A")
            continue
        elif a_sum < b_sum:
            print("B")
        

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 英語の授業中に暇だった paiza 君は、N ページの教科書を使った次のようなゲームを思いつきました。

# 「2人のプレイヤーがそれぞれ教科書の 1 箇所を適当に掴んで、掴んだ範囲のページに含まれる 'I' の合計数が多い方が勝ち、少ない方が負け、同じだったら引き分け！」
# 「ただし、 N/3 ページ以上掴んだ人は反則負けで、 2 人とも反則したら引き分け！」
# (ここでの N/3 は整数であるとは限りません)

# 審判を任されたあなたは、各ページに含まれる 'I' の数を記録しておくことで、掴んだページの両端の番号を確認するだけで、掴んだ範囲のページに含まれる 'I' の合計数を求めることができることに気付きました。

# 教科書のページ数 N と試合の数 K , 各ページの 'I' の数 I_1 ... I_N と、
# i 番目の試合で対戦した A と B の 2 人が掴んだページの両端のページ番号 A_l_i, A_r_i , B_l_i, B_r_iが与えられるので、各試合のジャッジしてください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# A_1
# ...
# A_N
# A_l_1 A_r_1 B_l_1 B_r_1
# ...
# A_l_K A_r_K B_l_K B_r_K


# ・1 行目では、教科書のページ数 N とおこなわれる試合数 K が与えられます。
# ・続く N 行のうち、 i 行目では、教科書の i ページ目に含まれる 'I' の数が与えられます。
# ・続く K 行のうち、 i 行目では、i 試合目のプレイヤー A が掴んだページの両端のページ番号 A_l_i , A_r_i とプレイヤー B が掴んだページの両端のページ番号 B_l_i , B_r_i が与えられます。

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
# ・1 ≦ N ≦ 100,000
# ・1 ≦ K ≦ 100,000
# ・0 ≦ A_i ≦ 100 (1 ≦ i ≦ N)
# ・1 ≦ A_l_i ≦ A_r_i ≦ N (1 ≦ i ≦ K)
# ・1 ≦ B_l_i ≦ B_r_i ≦ N (1 ≦ i ≦ K)

# 入力例1
# 3 3
# 0
# 1
# 2
# 1 1 2 2
# 2 2 3 3
# 3 3 3 3

# 出力例1
# DRAW
# DRAW
# DRAW

# 入力例2
# 10 3
# 10
# 9
# 8
# 7
# 6
# 5
# 4
# 3
# 2
# 1
# 1 3 7 10
# 1 4 3 4
# 1 5 6 10

# 出力例2
# A
# B
# DRAW