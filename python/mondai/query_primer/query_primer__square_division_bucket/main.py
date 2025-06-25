import sys
import unittest
import math
from collections import defaultdict
from typing import List, Dict


def get_square_division_bucket(items: List[int]) -> List[int]:
    # """"""""""""""""""
    # 平方分割でそれぞれ最大値を格納しListを返却
    # """"""""""""""""""
    size = int(math.sqrt(len(items)))
    answer = [-1] * size

    for i in range(size):
        start, end = size * i, size * (i + 1)
        answer[i] = max(items[start:end])

    return answer


class TestItemPriceMap(unittest.TestCase):
    def test_get_square_division_bucket(self):
        # 16の平方根は4
        input_items = [
            10,
            19,
            17,
            11,
            # 1
            21,
            26,
            29,
            28,
            # 2
            33,
            39,
            38,
            32,
            # 3
            41,
            40,
            44,
            49
            # 4
        ]
        # 4つのグループでそれぞれの最大値のList
        expected = [
            19,
            29,
            39,
            49
        ]
        actual = get_square_division_bucket(input_items)
        self.assertEqual(expected, actual)


def main():
    n = 10000
    # item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    items = [int(input().strip()) for _ in range(n)]
    # items = [list(map(int, input().strip().split())) for _ in range(item_count)]
    # queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    for v in get_square_division_bucket(items):
        print(v)

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# paiza くんは、長さ N の数列のある区間に含まれる要素の最大値を K 回求めたいのですが、与えられる区間の要素をいちいち全て調べていては時間計算量にして最大で O(NK) かかってしまいます。
# そこで、paiza くんは 平方分割 と言われるアルゴリズムを用いることで、この計算量を減らそうと考えました。
# 平方分割とは、次のようなアルゴリズムです。

# 1. 長さ N の配列が与えられたとき、N の平方根 x を求め、配列を長さ x の配列に分割し、それぞれの配列について目的の値を調べておく。
# （分割で得られる最後の配列の長さは必ずしも x になるとは限りません）
# 2. 調べたい区間に完全に含まれている配列についての 1. で求めた値と、その配列以外の部分の値を全て調べて、目的の値を求める。



# この問題では、長さ 10,000 の数列 A について手順 1. を行ってみましょう。
# 10,000 の平方根は 100 なので、先頭から 100 要素ずつの最大値を求めましょう。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# A_1
# ...
# A_10000


# ・10000 行で数列 A の要素が先頭から順に与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ans_1
# ...
# ans_100


# ・i 行目 (1 ≦ i ≦ 100) に A を 100 要素ずつに区切ったときの先頭から i 番目の区間の最大値を出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・-100,000 ≦ A_i ≦ 100,000 (1 ≦ i ≦ 10,000)
