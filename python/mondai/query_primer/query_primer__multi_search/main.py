import sys
import unittest
from typing import List


def get_is_exists_list(items: List[int], query_items: List[int]) -> List[str]:
    items_set = set(items) # setにしてハッシュにする
    result = []
    for q in query_items:
        if q in items_set: # 平均して O(1) で動作
            result.append("YES")
        else:
            result.append("NO")

    return result


class TestIsExistsList(unittest.TestCase):
    def test_get_is_exists_list(self):
        expected = ["YES", "YES", "YES", "NO", "NO"]
        actual = get_is_exists_list([1, 2, 3, 4, 5], [1, 3, 5, 7, 9])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    items = [input().strip() for _ in range(item_count)]
    query_items = [input().strip() for _ in range(query_count)]

    for v in get_is_exists_list(items, query_items):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 長さ N の重複した要素の無い数列 A と Q 個の整数 K_1 ... K_Q が与えられるので、
# 各 K_i について、 A に K_i が含まれていれば "YES" を、そうでなければ "NO" を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N Q
# A_1
# ...
# A_N
# K_1
# ...
# K_Q


# ・1 行目では、配列 A の要素数 N と検索する値の個数 Q が半角スペース区切りで与えられます。
# ・続く N 行では、配列 A の要素が先頭から順に与えられます。
# ・続く Q 行では、検索する値 K_1 .. K_Q が順に与えられます

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・Q 行出力してください。i 行目には A に K_i が含まれていれば "YES" を、そうでなければ "NO" を出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N , Q ≦ 100,000
# ・0 ≦ A_i ≦ 1,000,000 (1 ≦ i ≦ N)
# ・0 ≦ K_i ≦ 1,000,000 (1 ≦ i ≦ Q)

# 入力例1
# 5 5
# 1
# 2
# 3
# 4
# 5
# 1
# 3
# 5
# 7
# 9

# 出力例1
# YES
# YES
# YES
# NO
# NO

# 入力例2
# 10 5
# 351051
# 62992
# 166282
# 497610
# 636807
# 678131
# 885162
# 81763
# 810110
# 943644
# 670661
# 463229
# 62992
# 1973
# 901393

# 出力例2
# NO
# NO
# YES
# NO
# NO