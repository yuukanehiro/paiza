import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# クイックソートでList[int]を返却
# ------------------------
def bubble_sort(arr: List[int]) -> List[int]:
    n = len(arr)

    for i in range(n):
        for j in range(n - 1 - i):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]

    return arr


# ------------------------
# テストコード（unittest）
# ------------------------
class TestQuickSort(unittest.TestCase):
    def test_bubble_sort(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = bubble_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case2(self):
        expected = [1]
        actual = bubble_sort([1])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case3(self):
        expected = [1, 2]
        actual = bubble_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case4(self):
        expected = [1, 2, 3]
        actual = bubble_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case5_empty(self):
        expected = []
        actual = bubble_sort([])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = bubble_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_bubble_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = bubble_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(bubble_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()