import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 挿入ソートでList[int]を返却
# ------------------------
def insertion_sort(arr: List[int]) -> List[int]:
    for i in range(1, len(arr)):
        key: int = arr[i]

        j = i - 1
        while j >= 0 and arr[j] > key:
            arr[j + 1] = arr[j]
            j -= 1
        
        arr[j + 1] = key

    return arr


# ------------------------
# テストコード（unittest）
# ------------------------
class TestInsertionSort(unittest.TestCase):
    def test_insertion_sort_case1(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = insertion_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case2(self):
        expected = [1]
        actual = insertion_sort([1])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case3(self):
        expected = [1, 2]
        actual = insertion_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case4(self):
        expected = [1, 2, 3]
        actual = insertion_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case5_empty(self):
        expected = []
        actual = insertion_sort([])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = insertion_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_insertion_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = insertion_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(insertion_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()