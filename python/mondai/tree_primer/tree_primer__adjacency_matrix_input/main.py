import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# を返却
# ------------------------
def get_adjacency_matrix(item_count: int, queries: List[str]) -> List[List[int]]:
    # 0-indexed
    matrix = [[0] * item_count for _ in range(item_count)]
    for q in queries:
        a, b = map(int, q.split())
        a -= 1
        b -= 1
        matrix[a][b] = 1
        matrix[b][a] = 1

    return matrix

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestAdjacencyMatrix(unittest.TestCase):
    def test_get_adjacency_matrix(self):
        expected = [
            [0, 1, 1, 1, 0],
            [1, 0, 0, 0, 0],
            [1, 0, 0, 0, 1],
            [1, 0, 0, 0, 0],
            [0, 0, 1, 0, 0]
        ]
        actual = get_adjacency_matrix(5, ['1 2', '1 3', '1 4', '5 3'])
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    item_count = int(input())
    # item_count = int(input()) # 1つのintの場合
    # inputList = list(map(int, input().split()))
    # item_count, query_count = inputList[0], inputList[1]
    # items = [input().strip() for _ in range(item_count)]
    # 2次元配列
    # items = [list(map(int, input().split())) for _ in range(item_count)]
    # 1 - indexed
    # items = [0] + [int(input().strip()) for _ in range(item_count)]
    queries = [input().strip() for _ in range(item_count - 1)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))
    
    matrix = get_adjacency_matrix(item_count, queries)

    for i in range(len(matrix)):
        strArray = list(map(str, matrix[i]))
        print(" ".join(strArray))

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# https://paiza.jp/works/mondai/tree_primer/tree_primer__adjacency_matrix_input/edit?language_uid=python3&t=69124839b678f723d40e02a9e5e6cd67

# 下記の問題をプログラミングしてみよう！
# この問題集では、グラフ理論における木を扱います。
# 木とは、n 個の頂点と、それら全てを連結する n-1 個の辺からなるグラフのことです。
# 例として、以下の図 1 は頂点数が 5 の木ですが、辺の数が 5 である図 2 と、5 つの頂点全てを連結していない図 3 は木ではありません。

# 図 1

# 図 2

# 図 3


# プログラミングで木を扱う際には、辺の情報を利用しやすい形で保持することが好まれるので、隣接行列や隣接リストと呼ばれる形式で辺の情報を管理します。
# この問題では、隣接行列を用いて辺の情報を管理してみましょう。

# グラフ (頂点数 N) の隣接行列とは、 N × N の行列 g であって i 行 j 列目の要素が
# ・ i 番目の頂点と j 番目の頂点が辺で結ばれているとき 1
# ・ 結ばれていないとき 0
# であるようなもののことをいいます。

# 木の頂点・辺についての情報が与えられるので、この木の隣接行列を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# a_1 b_1
# ...
# a_{N-1} b_{N-1}


# ・1 行目には、頂点の数 N が与えられます。
# ・続く N-1 行では、各辺の両端の頂点 a_i , b_i が与えられます。(1 ≦ i ≦ N-1)

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・与えられた入力に対応する隣接行列 g を以下の形式で N 行で出力してください。
# ・各要素の間には半角スペースを出力してください。


# g[1][1] g[1][2] ... g[1][N]            
# ...
# g[N][1] g[N][2] ... g[N][N]
# 条件
# すべてのテストケースにおいて、以下の条件をみたします。

# ・ 1 ≦ N ≦ 100
# ・ 1 ≦ a_i , b_i ≦ N (1 ≦ i ≦ N-1)

# 入力例1
# 5
# 1 2
# 1 3
# 1 4
# 5 3

# 出力例1
# 0 1 1 1 0
# 1 0 0 0 0
# 1 0 0 0 1
# 1 0 0 0 0
# 0 0 1 0 0