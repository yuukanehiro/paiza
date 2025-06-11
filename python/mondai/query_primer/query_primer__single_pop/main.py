import sys
import unittest
from typing import List


def get_poped_items(items: List[int], start_index: int) -> List[int]:
    return items[start_index:]


class TestItemPriceMap(unittest.TestCase):
    def test_get_poped_items(self):
        expected = [2, 3, 4, 5]
        actual = get_poped_items([1, 2, 3, 4, 5], 1)
        self.assertEqual(expected, actual)
    def test_get_poped_items(self):
        expected = [4, 5]
        actual = get_poped_items([1, 2, 3, 4, 5], 3)
        self.assertEqual(expected, actual)

def main():
    item_count = int(input())
    items = [input().strip() for _ in range(item_count)]

    for v in get_poped_items(items, 1):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 数列 A が与えられるので、A の先頭の要素を削除した後の A を出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N
# A_1
# ...
# A_N


# ・1 行目では、配列 A の要素数 N が与えられます。
# ・続く N 行では、配列 A の要素が先頭から順に与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# A_2
# ...
# A_N

# ・A の先頭の要素を削除した後の A の各要素を先頭から順に改行区切りで出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・2 ≦ N ≦ 100,000
# ・0 ≦ A_i ≦ 10,000 (1 ≦ i ≦ N)

# 入力例1
# 10
# 5980
# 1569
# 5756
# 9335
# 9680
# 4571
# 5309
# 8696
# 9680
# 8963

# 出力例1
# 1569
# 5756
# 9335
# 9680
# 4571
# 5309
# 8696
# 9680
# 8963

# 入力例2
# 2
# 6963
# 9374

# 出力例2
# 9374