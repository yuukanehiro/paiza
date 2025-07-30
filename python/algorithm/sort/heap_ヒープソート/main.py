import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple


# ------------------------
# ヒープ構造にしてList[int]を返却
# ------------------------
def heapfy(arr: List[int], end: int, i: int) -> List[int]:
    largest = i
    left = i * 2 + 1
    right = i * 2 + 2

    if left < end and arr[left] > arr[largest]:
        largest = left

    if right < end and arr[right] > arr[largest]:
        largest = right

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapfy(arr, end, largest)

    return arr


# ------------------------
# ヒープソートでList[int]を返却
# ------------------------
def heap_sort(arr: List[int]) -> List[int]:
    n = len(arr)

    # ヒープ作成
    for i in range(n // 2, -1, -1): # n // 2から0まで
        heapfy(arr, n, i)

    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapfy(arr, i, 0)

    return arr


# ------------------------
# テストコード（unittest）
# ------------------------
class TestQuickSort(unittest.TestCase):
    def test_heap_sort(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = heap_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_heap_sort_case2(self):
        expected = [1]
        actual = heap_sort([1])
        self.assertEqual(expected, actual)
    def test_heap_sort_case3(self):
        expected = [1, 2]
        actual = heap_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_heap_sort_case4(self):
        expected = [1, 2, 3]
        actual = heap_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_heap_sort_case5_empty(self):
        expected = []
        actual = heap_sort([])
        self.assertEqual(expected, actual)
    def test_heap_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = heap_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_heap_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = heap_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(heap_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()