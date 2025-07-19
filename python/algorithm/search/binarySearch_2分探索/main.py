import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 2分探索した結果をboolで返却
# ------------------------
def binary_search(arr: List[int], target: int) -> bool:
    # ソート
    sorted_arr: List[int] = sorted(arr)

    left_index: int = 0
    right_index: int = len(sorted_arr) - 1 # 0-indexed

    while left_index <= right_index:
        mid_index: int = (left_index + right_index) // 2

        if sorted_arr[mid_index] == target:
            return True

        if target > sorted_arr[mid_index]:
            left_index = mid_index + 1
        else:
            right_index = mid_index - 1

    return False


# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestBinarySearch(unittest.TestCase):
    def test_binary_search_simple(self):
        input_arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]
        input_target: int = 5
        expected = True
        actual = binary_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_binary_search_duplicated(self):
        input_arr: List[int] = [5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = binary_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_binary_search_zero(self):
        input_arr: List[int] = [0, 5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = binary_search(input_arr, input_target)
        self.assertEqual(expected, actual)

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]

    target_5: int = 5
    print(binary_search(arr, target_5)) # True
    target_11: int = 11
    print(binary_search(arr, target_11)) # False


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()