import sys
import unittest
from typing import List
from bisect import bisect_left

def get_insert_count(items: List[int], target: int) -> int:
    sorted_items = sorted(items)
    return bisect_left(sorted_items, target) + 1


class TestItemPriceMap(unittest.TestCase):
    def test_get_item_price_map(self):
        expected = 2
        actual = get_insert_count([181, 177, 113, 188], 174)
        self.assertEqual(expected, actual)


def main():
    item_count, addStudentHeight, targetStudentHeight = map(int, input().split())
    items = [int(input().strip()) for _ in range(item_count)]
    items.append(addStudentHeight)

    print(get_insert_count(items, targetStudentHeight))


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# paiza 君のクラスには paiza 君を含めて N + 1 人の生徒がいます。paiza 君の身長は P cm です。paiza 君以外の N 人の生徒の身長は A_1, ... ,A_N です。
# 今日、クラスに身長 X cm の転校生が 1 人やってきました。転校生が入ってきた後 N + 2 人のクラス全員で背の順で並んだ時、 paiza 君は前から何番目に並ぶことになるでしょうか。

# なお、背の順の先頭の生徒を前から 1 番目の生徒とします。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N X P
# A_1
# ...
# A_N


# ・1 行目では、クラスの paiza 君以外の生徒数 N と転校生の身長 X と paiza君の身長 P が与えられます。
# ・続く N 行では、N 人の生徒の身長が与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# ・転校生を含む N+2 人が背の順に並んだときに paiza 君が前から何番目になるかを 1 行で出力してください。
# ・また、出力の末尾には改行を入れてください。

# 条件
# ・1 ≦ N ≦ 100,000
# ・100 ≦ X , P ≦ 200
# ・100 ≦ A_i ≦ 200 (1 ≦ i ≦ N)
# ・転校生を含め、クラスの中で身長が P cm の生徒は paiza 君のみであることが保証されている

# 入力例1
# 3 188 174
# 181
# 177
# 113

# 出力例1
# 2

# 入力例2
# 10 139 146
# 165
# 159
# 144
# 195
# 188
# 118
# 118
# 141
# 199
# 124

# 出力例2
# 7
