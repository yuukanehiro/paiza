import sys
import unittest
from collections import deque
from typing import List, Deque

def get_merged_list(items: Deque[str], queries: List[str]) -> List[str]:
    result = []
    for q in queries:
        if q == "pop":
            items.popleft()
        elif q == "show":
            result.extend(items)

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_merged_list_case1(self):
        expected = ["2410", "9178", "7252"]
        actual = get_merged_list(deque(["7564", "4860", "2410", "9178", "7252"]), ["pop", "pop", "show"])
        self.assertEqual(expected, actual)
    def test_get_merged_list_case2(self):
        expected = ["1339", "4960", "3926", "9816", "3018", "4213", "9816", "3018", "4213"]
        actual = get_merged_list(deque(["1005", "2716", "7856", "8546", "1339", "4960", "3926", "9816", "3018", "4213"]), ["pop", "pop", "pop", "pop", "show", "pop", "pop", "pop", "show", "pop"])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    items = deque(input().strip() for _ in range(item_count))
    queries = [input().strip() for _ in range(query_count)]

    for v in get_merged_list(items, queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 先頭の要素の削除(query) Python3編（paizaランク C 相当）
# 問題にチャレンジして、ユーザー同士で解答を教え合ったり、コードを公開してみよう！

# シェア用URL:
# https://paiza.jp/works/mondai/query_primer/query_primer__multi_pop
# 問題文のURLをコピーする
#  下記の問題をプログラミングしてみよう！
# 数列 A と入力の回数 K が与えられるので、K 回の入力に応じて次のような処理をしてください。
# ・pop
# A の先頭の要素を削除する。既に A に要素が存在しない場合何もしない。
# ・show
# A の要素を先頭から順に改行区切りで出力する。A に要素が存在しない場合何も出力しない。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# A_1
# ...
# A_N
# S_1
# ...
# S_K


# ・1 行目では、配列 A の要素数 N と与えられる入力の数 K が与えられます。
# ・続く N 行では、配列 A の要素が先頭から順に与えられます。
# ・続く K 行では、"pop" または "show" が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・S_i で "show" が与えられる度に、A の全ての要素を先頭から順に改行区切りで出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ K ≦ N ≦ 100,000
# ・0 ≦ A_i ≦ 10,000 (1 ≦ i ≦ N)
# ・S_i (1 ≦ i ≦ K) は "pop" , "show" のいずれか
# ・S_i のうち、"show" であるものは 10 個以下であることが保証されている。

# 入力例1
# 5 3
# 7564
# 4860
# 2410
# 9178
# 7252
# pop
# pop
# show

# 出力例1
# 2410
# 9178
# 7252

# 入力例2
# 10 10
# 1005
# 2716
# 7856
# 8546
# 1339
# 4960
# 3926
# 9816
# 3018
# 4213
# pop
# pop
# pop
# pop
# show
# pop
# pop
# pop
# show
# pop

# 出力例2
# 1339
# 4960
# 3926
# 9816
# 3018
# 4213
# 9816
# 3018
# 4213