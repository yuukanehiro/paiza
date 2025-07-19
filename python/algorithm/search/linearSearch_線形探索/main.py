import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 線形探索した結果をboolで返却
# ------------------------
def linear_search(arr: List[int], target: int) -> bool:
    for v in arr:
        if target == v:
            return True

    return False


# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestLinearSearch(unittest.TestCase):
    def test_linear_search_simple(self):
        input_arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]
        input_target: int = 5
        expected = True
        actual = linear_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_linear_search_duplicated(self):
        input_arr: List[int] = [5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = linear_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_linear_search_zero(self):
        input_arr: List[int] = [0, 5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = linear_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_linear_search_not_found(self):
        input_arr = [1, 2, 3, 4, 5]
        input_target = 9
        expected = False
        actual = linear_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_linear_search_empty(self):
        input_arr = []
        input_target = 1
        expected = False
        actual = linear_search(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_linear_search_single_element_true(self):
        self.assertTrue(linear_search([7], 7))
    def test_linear_search_single_element_false(self):
        self.assertFalse(linear_search([7], 3))

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]

    target_5: int = 5
    print(linear_search(arr, target_5)) # True
    target_11: int = 11
    print(linear_search(arr, target_11)) # False


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()