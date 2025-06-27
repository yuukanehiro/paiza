import sys
import unittest
from typing import List


# ------------------------
# 累積和テーブルを構築して返却
# ------------------------
def build_cumulative_sum_matrix(items: List[List[int]], H: int, W: int) -> List[List[int]]:
    sum_matrix = [[0] * (W + 1) for _ in range(H + 1)]
    for y in range(1, H + 1):
        for x in range(1, W + 1):
            sum_matrix[y][x] = (
                sum_matrix[y - 1][x]
                + sum_matrix[y][x - 1]
                - sum_matrix[y - 1][x - 1]
                + items[y - 1][x - 1]
            )
    return sum_matrix


# ------------------------
# チョコ数を範囲合計で取得して返却
# ------------------------
def calc(sum_matrix: List[List[int]], top: int, left: int, bottom: int, right: int) -> int:
    ans = sum_matrix[bottom][right]
    if top > 0 and left > 0:
        ans += sum_matrix[top - 1][left - 1]
    if top > 0:
        ans -= sum_matrix[top - 1][right]
    if left > 0:
        ans -= sum_matrix[bottom][left - 1]
    return ans


# ------------------------
# ドーナツのチョコの数を返却
# 外側のチョコ - 内側のチョコ = ドーナツのチョコの数
# ------------------------
def solve_from_input(items: List[List[int]], H: int, W: int, queries: List[str]) -> List[int]:
    sum_matrix = build_cumulative_sum_matrix(items, H, W)
    results = []

    for q in queries:
        pivotY, pivotX, outer_width, inner_width = map(int, q.split())
        outer_radius = outer_width // 2
        inner_radius = inner_width // 2

        outer = calc(sum_matrix,
                     pivotY - outer_radius, pivotX - outer_radius,
                     pivotY + outer_radius, pivotX + outer_radius)
        inner = calc(sum_matrix,
                     pivotY - inner_radius, pivotX - inner_radius,
                     pivotY + inner_radius, pivotX + inner_radius)
        results.append(outer - inner)

    return results

def main():
    H, W, query_count = map(int, input().split())
    items = [list(map(int, input().split())) for _ in range(H)]
    queries = [input() for _ in range(query_count)]

    results = solve_from_input(items, H, W, queries)
    for v in results:
        print(v)


# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestDonutChoco(unittest.TestCase):
    def test_cumulative_sum_matrix(self):
        items = [
            [1, 2, 3],
            [4, 5, 6],
            [7, 8, 9]
        ]
        H, W = 3, 3

        expected = [
            [0, 0, 0, 0],
            [0, 1, 3, 6],
            [0, 5, 12, 21],
            [0, 12, 27, 45]
        ]

        actual = build_cumulative_sum_matrix(items, H, W)
        self.assertEqual(actual, expected)

    def test_sample1(self):
        stdin = """3 3 1
1 2 3
4 5 6
7 8 9
2 2 3 1"""
        lines = stdin.strip().split("\n")
        H, W, N = map(int, lines[0].split())
        items = [list(map(int, line.split())) for line in lines[1:H+1]]
        queries = lines[H+1:]
        expected = [40]
        actual = solve_from_input(items, H, W, queries)
        self.assertEqual(actual, expected)

    def test_sample2(self):
        stdin = """5 5 4
7 8 9 8 5
4 2 6 2 1
2 5 3 9 1
1 3 3 2 3
2 3 2 6 6
3 4 3 1
2 3 3 1
2 4 3 1
3 3 5 3"""
        lines = stdin.strip().split("\n")
        H, W, N = map(int, lines[0].split())
        items = [list(map(int, line.split())) for line in lines[1:H+1]]
        queries = lines[H+1:]
        expected = [21, 46, 42, 68]
        actual = solve_from_input(items, H, W, queries)
        self.assertEqual(actual, expected)


# ------------------------
# 起動時分岐
# ------------------------
if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/query_primer/query_primer__dount/edit?language_uid=python3&t=01fbc8b4bbaf5aff1e238a91c96c58f3

# 下記の問題をプログラミングしてみよう！
# あなたはドーナツ屋 paiza の店員です. この店では, チョコが埋め込まれた四角のドーナツが名物となっています.

# H x W cm^2 の生地があります. これは 1 cm^2 ごとに区画されており, H x W 個の区画を持っています.
# また, それぞれの区間には既にいくつかのチョコが散りばめられています.

# この 1 枚の生地から四角いドーナツの形を 1 つ切り出すことができます.
# 具体的には次の工程をおこないます.

# 1. 上から y 番目、 左から x 番目の区間を中心に一辺が B cm の正方形の切れ目を入れる。
# 2. 上から y 番目、 左から x 番目の区間を中心に一辺が S cm の正方形の切れ目を入れることでドーナツ状の切れ目を完成させる。（S < B であることが保証されています）

# 生地を成す H*W 個の各区画について、その区画に含まれるチョコの数と、作る N 個のドーナツについての情報が与えられるので、各ドーナツにチョコがいくつ含まれることになるかを求めてください。
# なお、 N 枚の生地について、含まれるチョコの分布は全て同じであることがわかっています。

# 例として、入力例 1 では、生地は 3 × 3 cm^2 であり、次の通り区画されています。



# また、入力の情報から、生地には次の通りチョコが乗っていることがわかります。



# 左から 2 番目、上から 2 番目の区画を中心とする 3 × 3 の正方形をくり抜くと、次の図の水色の部分になります。



# また、そこからさらに左から 2 番目、上から 2 番目の区画を中心とする 1 × 1 の正方形をくり抜くと、次のようなドーナツができあがります。



# このドーナツに乗っているチョコの数は 40 個であるので、答えとして 40 を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# H W N
# C[1][1] ... C[1][W]
# ...
# C[H][1] ... C[H][W]
# y_1 x_1 B_1 S_1
# ...
# y_N x_N B_N S_N


# ・1 行目では、ドーナツの生地の縦の長さ H(cm) と横の長さ W(cm) と作るドーナツの数 N が与えられます。
# ・続く H 行では、生地の 1 cm^2 に含まれるチョコの数が左上から順に見た目の通りに半角スペース区切りで与えられます。
# ・続く N 行のうち、i 行目では、i 番目に作るドーナツの中心の生地の上からの距離 y_i (cm) と左からの距離 x_i (cm) とドーナツの外側の一辺の長さ B_i (cm) と 内側の一辺の長さ S_i (cm) が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_N


# ・i 行目に i 個目のドーナツに含まれるチョコの数 ans_i を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・3 ≦ H , W ≦ 1,000
# ・1 ≦ N ≦ 100,000
# ・0 ≦ C[i][j] < 10 (1 ≦ i ≦ H , 1 ≦ j ≦ W)
# ・1 ≦ y_i ≦ H (1 ≦ i ≦ N)
# ・1 ≦ x_i ≦ W (1 ≦ i ≦ N)
# ・0 ≦ S_i < B_i ≦ min(H,W) (1 ≦ i ≦ N)
# ・S_i , B_i は奇数
# ・くり抜けないようなドーナツの入力は与えられないことが保証されている

# 入力例1
# 3 3 1
# 1 2 3
# 4 5 6
# 7 8 9
# 2 2 3 1

# 出力例1
# 40

# 入力例2
# 5 5 4
# 7 8 9 8 5
# 4 2 6 2 1
# 2 5 3 9 1
# 1 3 3 2 3
# 2 3 2 6 6
# 3 4 3 1
# 2 3 3 1
# 2 4 3 1
# 3 3 5 3

# 出力例2
# 21
# 46
# 42
# 68