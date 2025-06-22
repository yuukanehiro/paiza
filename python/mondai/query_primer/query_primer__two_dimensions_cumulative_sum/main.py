import sys
import unittest
from collections import defaultdict
from typing import List


def get_two_dimensions_cumulative_sum_matrix(A: List[List[int]], H: int, W: int) -> List[List[int]]:
    sum_matrix = [[0] * (W + 1) for _ in range(H + 1)]
    for y in range(1, H + 1):
        for x in range(1, W + 1):
            sum_matrix[y][x] = (
                sum_matrix[y - 1][x]
                + sum_matrix[y][x-1]
                - sum_matrix[y - 1][x - 1]
                + A[y-1][x-1]
            )

    return sum_matrix
    



class TestGetTwoDimensionsCumulativeSumMatrix(unittest.TestCase):
    def test_get_two_dimensions_cumulative_sum_matrix(self):
        expected = [
            [0, 0, 0, 0],
            [0, 1, 3, 6],
            [0, 5, 12, 21],
            [0, 12, 27, 45]
        ]
        input_matrix = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]
        ]
        input_h = 3
        input_w = 3
        actual = get_two_dimensions_cumulative_sum_matrix(input_matrix, input_h, input_w)
        self.assertEqual(expected, actual)


def main():
    H, W, N = map(int, input().strip().split())

    A = []
    for _ in range(H):
        row = list(map(int, input().strip().split()))
        A.append(row)

    sum_matrix = get_two_dimensions_cumulative_sum_matrix(A, H, W)

    queries = [input().strip() for _ in range(N)]

    for q in queries:
        y, x = map(int, q.split())
        print(sum_matrix[y][x])


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# H 行 W 列 の行列 A の y 行 x 列における累積和 S(y,x) を以下の数式・図の通り定義します。以後 A の y 行 x 列の要素を A[y][x] と表すことにします。

# S(y,x) = A[1][1] + A[1][2] + ... + A[1][x] + A[2][1] + ... + A[2][x] + ... + A[y][1] + ... + A[y][x]

# H 行 W 列 の二次元配列 A と、累積和を求めたい行・列番号についての情報が与えられるので、各ペアについて累積和を求めてください。
# 例として、入力例 1 の行列における累積和 S(2,2) は次のピンクの部分の和となり、S(2,2) = 12 となります。



# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# H W N
# A[1][1] ... A[1][W]
# ...
# A[H][1] ... A[H][W]
# y_1 x_1
# ...
# y_N x_N


# ・1 行目では、配列 A の行数 H と列数 W , 与えられる整数のペアの個数 N が与えられます。
# ・続く H 行のうち i 行目では、行列 A の i 行の要素が半角スペース区切りで A[i][1] から順に与えられます。
# ・続く N 行では、累積和を求めるのに使う整数のペア y , x が N 個与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_N


# ・i 行目に S(y_i , x_i) の値 ans_i (1 ≦ i ≦ N) を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ H , W ≦ 1,000
# ・1 ≦ N ≦ 100,000
# ・-100 ≦ A[i][j] ≦ 100 (1 ≦ i ≦ H , 1 ≦ j ≦ W)
# ・1 ≦ y_i ≦ H , 1 ≦ x_i ≦ W (1 ≦ i ≦ N)

# 入力例1
# 3 3 3
# 1 2 3
# 4 5 6
# 7 8 9
# 1 1
# 2 2
# 3 3

# 出力例1
# 1
# 12
# 45

# 入力例2
# 3 3 3
# -76 -44 61
# -72 30 95
# 99 42 82
# 3 3
# 1 1
# 3 1

# 出力例2
# 217
# -76
# -49