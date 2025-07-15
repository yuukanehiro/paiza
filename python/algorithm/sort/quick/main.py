import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# クイックソートでList[int]を返却
# ------------------------
def quick_sort(arr: List[int]) -> List[int]:
    if len(arr) <= 1:
        return arr

    pivot = arr[0]
    left = [x for x in arr[1:] if x < pivot]
    right = [x for x in arr[1:] if x > pivot]

    return quick_sort(left) + [pivot] + quick_sort(right)


# ------------------------
# テストコード（unittest）
# ------------------------
class TestQuickSort(unittest.TestCase):
    def test_quick_sort(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = quick_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)


def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(quick_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()