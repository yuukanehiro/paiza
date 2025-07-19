import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# ハッシュ探索した結果をboolで返却
# ------------------------
def hash_search_by_set(arr: List[int], target: int) -> bool:
    s: set = set(arr)
    if target in s:
        return True
    else:
        return False

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestHashSearchBySet(unittest.TestCase):
    def test_hash_search_by_set_simple(self):
        input_arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]
        input_target: int = 5
        expected = True
        actual = hash_search_by_set(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_hash_search_by_set_duplicated(self):
        input_arr: List[int] = [5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = hash_search_by_set(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_hash_search_by_set_zero(self):
        input_arr: List[int] = [0, 5, 3, 6, 2, 1, 1, 8, 7, 10, 5]
        input_target: int = 5
        expected = True
        actual = hash_search_by_set(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_hash_search_by_set_not_found(self):
        input_arr = [1, 2, 3, 4, 5]
        input_target = 9
        expected = False
        actual = hash_search_by_set(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_hash_search_by_set_empty(self):
        input_arr = []
        input_target = 1
        expected = False
        actual = hash_search_by_set(input_arr, input_target)
        self.assertEqual(expected, actual)
    def test_hash_search_by_set_single_element_true(self):
        self.assertTrue(hash_search_by_set([7], 7))
    def test_hash_search_by_set_single_element_false(self):
        self.assertFalse(hash_search_by_set([7], 3))

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    arr: List[int] = [99, 3, 2, 1, 4, 5, 7, 6, 8, 9, 10, 0]

    target_5: int = 5
    print(hash_search_by_set(arr, target_5)) # True
    target_11: int = 11
    print(hash_search_by_set(arr, target_11)) # False


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()