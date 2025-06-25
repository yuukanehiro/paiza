import sys
import unittest
from typing import List, Tuple


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

def solve_query_sums(
    sum_matrix: List[List[int]],
    queries: List[Tuple[int, int, int, int]]
) -> List[int]:
    result = []
    for a, b, c, d in queries:
        val = sum_matrix[c][d]
        if a > 1:
            val -= sum_matrix[a - 1][d]
        if b > 1:
            val -= sum_matrix[c][b - 1]
        if a > 1 and b > 1:
            val += sum_matrix[a - 1][b - 1]
        result.append(val)
    return result

class TestSolveQuerySums(unittest.TestCase):
    def test_solve_query_sums(self):
        sum_matrix = [
            [0, 0, 0, 0],
            [0, 1, 3, 6],
            [0, 5, 12, 21],
            [0, 12, 27, 45]
        ]
        # A
        # [1, 2, 3],
        # [4, 5, 6],
        # [7, 8, 9]
        queries = [
            (1, 1, 3, 3),  # 全体 → 45
            (1, 2, 2, 3),  # 2+3+5+6 = 16
            (2, 2, 3, 3),  # 5+6+8+9 = 28
            (3, 3, 3, 3),  # 9
        ]

        expected = [45, 16, 28, 9]
        actual = solve_query_sums(sum_matrix, queries)
        self.assertEqual(expected, actual)


def main():
    H, W, N = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    A = [list(map(int, input().strip().split())) for _ in range(H)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    sum_matrix = get_two_dimensions_cumulative_sum_matrix(A, H, W)

    queries = [tuple(map(int, input().split())) for _ in range(N)]
    results = solve_query_sums(sum_matrix, queries)
    for val in results:
        print(val)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# H 行 W 列 の行列 A の 2 つの行・列番号の組 {a , b} , {c , d} における区間和 S({a,b} , {c,d}) (a ≦ c , b ≦ d) を以下の数式・図の通り定義します。以後 A の y 行 x 列の要素を A[y][x] と表すことにします。

# S({a,b} , {c,d}) = A[a][b] + A[a][b+1] + ... + A[a][d] + A[a+1][1] + ... + A[a+1][d] + ... + A[c][1] + ... + A[c][d]

# 例として、入力例 1 の A における S({2,2},{3,3}) は以下の通りになり、値は 28 となります。



# H 行 W 列 の行列 A と、区間和を求めたいペアについての情報が与えられるので、各ペアについて累積和を求めてください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# H W N
# A[1][1] ... A[1][W]
# ...
# A[H][1] ... A[H][W]
# a_1 b_1 c_1 d_1
# ...
# a_N b_N c_N d_N


# ・1 行目では、配列 A の行数 H と列数 W , 与えられる整数のペアの個数 N が与えられます。
# ・続く H 行のうち i 行目では、行列 A の i 行の要素が半角スペース区切りで A[i][1] から順に与えられます。
# ・続く N 行のうち i 行目では、累積和を求めるのに使う 2 つの行番号・列番号のペア {a_i , b_i} {c_i , d_i} が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_N


# ・i 行目に区間和 S({a_i,b_i} , {c_i,d_i}) の値 ans_i (1 ≦ i ≦ K)を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ H , W ≦ 1,000
# ・1 ≦ N ≦ 100,000
# ・-100 ≦ A[i][j] ≦ 100 (1 ≦ i ≦ H , 1 ≦ j ≦ W)
# ・1 ≦ a_i ≦ c_i ≦ H (1 ≦ i ≦ N)
# ・1 ≦ b_i ≦ d_i ≦ W (1 ≦ i ≦ N)

# 入力例1
# 3 3 2
# 1 2 3
# 4 5 6
# 7 8 9
# 1 1 3 3
# 1 2 2 3

# 出力例1
# 45
# 16

# 入力例2
# 10 10 4
# -74 -92 65 11 -96 66 17 33 -86 29
# 26 -83 100 -72 85 51 -29 8 49 72
# -47 52 -69 85 23 80 -59 79 92 -97
# 14 -26 -15 9 -22 -65 -29 -66 -30 -48
# -17 -68 -22 -50 -48 14 -29 96 -77 -23
# -96 -83 -31 46 8 -67 -94 92 -70 -49
# -97 8 42 -49 87 -72 -73 -80 68 66
# 100 94 -57 -62 -58 -18 -42 -80 55 47
# 18 -86 97 14 -37 -69 -19 56 58 -96
# 48 12 75 -68 19 52 -88 54 -24 -45
# 6 2 7 7
# 7 7 8 10
# 1 5 2 5
# 6 6 10 7

# 出力例2
# -278
# -39
# -11
# -490