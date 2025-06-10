import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def is_exists(items: List[int], target_number: int) -> bool:
    return target_number in items


class TestIsExists(unittest.TestCase):
    def test_is_exists_true(self):
        expected = True
        actual = is_exists([1, 3, 5], 5)
        self.assertEqual(expected, actual)
    def test_is_exists_false(self):
        expected = False
        actual = is_exists([1, 2, 3, 5, 6], 4)
        self.assertEqual(expected, actual)

def main():
    item_count, target_number = map(int, input().split())
    items = [int(input().strip()) for _ in range(item_count)]
    
    if is_exists(items, target_number):
        print("YES")
    else:
        print("NO")


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 長さ N の重複した要素の無い数列 A と整数 K が与えられるので、
# A に K が含まれていれば "YES" を、そうでなければ "NO" を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K
# A_1
# ...
# A_N


# ・1 行目では、配列 A の要素数 N と検索する値 K が半角スペース区切りで与えられます。
# ・続く N 行では、配列 A の要素が先頭から順に与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・A に K が含まれていれば "YES" を、そうでなければ "NO" を 1 行で出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N ≦ 100,000
# ・0 ≦ A_i , K ≦ 1,000,000 (1 ≦ i ≦ N)

# 入力例1
# 3 5
# 1
# 3
# 5

# 出力例1
# YES

# 入力例2
# 5 4
# 1
# 2
# 3
# 5
# 6

# 出力例2
# NO