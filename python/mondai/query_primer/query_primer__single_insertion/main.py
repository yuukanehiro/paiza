import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_merged_slice(items: List[int], target_index: int, target_number: int) -> List[int]:
    result = []
    result.extend(items[:target_index])
    result.append(target_number)
    result.extend(items[target_index:])

    return result


class TestGetMergedSlice(unittest.TestCase):
    def test_get_merged_slice_case1(self):
        expected = [17, 57, 57, 83]
        actual = get_merged_slice([17, 57, 83], 1, 57)
        self.assertEqual(expected, actual)
    def test_get_merged_slice_case2(self):
        expected = [38, 83, 46, 57, 15, 30, 45, 51, 88, 96, 85]
        actual = get_merged_slice([38, 83, 46, 57, 15, 30, 51, 88, 96, 85], 6, 45)
        self.assertEqual(expected, actual)


def main():
    item_count, target_index, target_number = map(int, input().split())
    items = [input().strip() for _ in range(item_count)]

    for v in get_merged_slice(items, target_index, target_number):
        print(v)
    
if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()

# Q
# 下記の問題をプログラミングしてみよう！
# 整数 N, K, Q と、 長さ N の配列 A_1, A_2, ..., A_N が与えられるので、A_K の後ろに Q を挿入した後の長さ N+1 の配列について、先頭から改行区切りで出力してください。

# ▼　下記解答欄にコードを記入してみよう

# 入力される値
# N K Q
# A_1
# ...
# A_N


# ・1 行目では、配列 A の要素数 N と整数 K , Q が半角スペース区切りで与えられます。
# ・続く N 行では、配列 A の要素が先頭から順に与えられます。

# 入力値最終行の末尾に改行が１つ入ります。
# 文字列は標準入力から渡されます。 標準入力からの値取得方法はこちらをご確認ください
# 期待する出力
# A_1
# ...
# A_{N+1}


# ・Q を A_K の後ろに挿入した後の配列の各要素を先頭から改行区切りで出力してください。
# ・また、出力の末尾には改行を入れてください。
# 条件
# ・1 ≦ N ≦ 100,000
# ・1 ≦ K ≦ N
# ・0 ≦ Q ≦ 100
# ・0 ≦ A_i ≦ 100 (1 ≦ i ≦ N)

# 入力例1
# 3 1 57
# 17
# 57
# 83

# 出力例1
# 17
# 57
# 57
# 83

# 入力例2
# 10 6 45
# 38
# 83
# 46
# 57
# 15
# 30
# 51
# 88
# 96
# 85

# 出力例2
# 38
# 83
# 46
# 57
# 15
# 30
# 45
# 51
# 88
# 96
# 85
