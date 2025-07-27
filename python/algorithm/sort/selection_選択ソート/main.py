import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# クイックソートでList[int]を返却
# ------------------------
def selection_sort(arr: List[int]) -> List[int]:
    n = len(arr)
    for i in range(n):
        min_index = i

        for j in range(i + 1, n):
            if arr[j] < arr[min_index]:
                min_index = j

        arr[i], arr[min_index] = arr[min_index], arr[i]
    return arr


# ------------------------
# テストコード（unittest）
# ------------------------
class TestSelectionSort(unittest.TestCase):
    def test_selection_sort(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = selection_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_selection_sort_case2(self):
        expected = [1]
        actual = selection_sort([1])
        self.assertEqual(expected, actual)
    def test_selection_sort_case3(self):
        expected = [1, 2]
        actual = selection_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_selection_sort_case4(self):
        expected = [1, 2, 3]
        actual = selection_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_selection_sort_case5_empty(self):
        expected = []
        actual = selection_sort([])
        self.assertEqual(expected, actual)
    def test_selection_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = selection_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_selection_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = selection_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(selection_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()