import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# 2分挿入ソートでList[int]を返却
# ------------------------
def binary_insertion_sort(arr: List[int]) -> List[int]:
    for i in range(1, len(arr)):
        key = arr[i]

        left = 0
        right = i - 1
        while left <= right:
            mid = (left + right) // 2
            if arr[mid] > key:
                right = mid - 1
            else:
                left = mid + 1

        # arr[i + 1:] ... arr[i]の値であるkeyが被らないようにiをずらして挿入
        arr = arr[:left] + [key] + arr[left:i] + arr[i + 1:]

    return arr


# # ------------------------
# # テストコード（unittest）
# # ------------------------
class TestBinaryInsertionSort(unittest.TestCase):
    def test_binary_insertion_sort_case1(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = binary_insertion_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case2(self):
        expected = [1]
        actual = binary_insertion_sort([1])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case3(self):
        expected = [1, 2]
        actual = binary_insertion_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case4(self):
        expected = [1, 2, 3]
        actual = binary_insertion_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case5_empty(self):
        expected = []
        actual = binary_insertion_sort([])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = binary_insertion_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_binary_insertion_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = binary_insertion_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(binary_insertion_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()