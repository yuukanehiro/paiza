import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# マージソートでList[int]を返却
# 要素1の[]をleft, rightで作ってそれをMergeしていくソート
# ------------------------
def merge_sort(arr: List[int]) -> List[int]:
    if len(arr) <= 1:
        return arr

    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])

    return merge(left, right)


def merge(left: List[int], right: List[int]) -> List[int]:
    result = []
    i = j = 0

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1

    result.extend(left[i:])
    result.extend(right[j:])

    return result



# # ------------------------
# # テストコード（unittest）
# # ------------------------
class TestMergeSort(unittest.TestCase):
    def test_merge_sort_case1(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = merge_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_merge_sort_case2(self):
        expected = [1]
        actual = merge_sort([1])
        self.assertEqual(expected, actual)
    def test_merge_sort_case3(self):
        expected = [1, 2]
        actual = merge_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_merge_sort_case4(self):
        expected = [1, 2, 3]
        actual = merge_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_merge_sort_case5_empty(self):
        expected = []
        actual = merge_sort([])
        self.assertEqual(expected, actual)
    def test_merge_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = merge_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_merge_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = merge_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(merge_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()