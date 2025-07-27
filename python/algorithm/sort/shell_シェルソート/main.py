import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# シェルソートでList[int]を返却
# gapをざっくり定義してソートしてgapを1まで狭めてソートする
# ------------------------
def shell_sort(arr: List[int]) -> List[int]:
    n = len(arr)

    gap = len(arr) // 2

    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            j = i
            # 挿入ソート ... 挿入位置を求めてtempを挿入
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap

            arr[j] = temp
        gap //= 2

    return arr



# ------------------------
# テストコード（unittest）
# ------------------------
class TestQuickSort(unittest.TestCase):
    def test_shell_sort(self):
        expected = [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]
        actual = shell_sort([2, 99, 5, 1, 3, 7, 8, 4, 9, 10])
        self.assertEqual(expected, actual)
    def test_shell_sort_case2(self):
        expected = [1]
        actual = shell_sort([1])
        self.assertEqual(expected, actual)
    def test_shell_sort_case3(self):
        expected = [1, 2]
        actual = shell_sort([1, 2])
        self.assertEqual(expected, actual)
    def test_shell_sort_case4(self):
        expected = [1, 2, 3]
        actual = shell_sort([3, 2, 1])
        self.assertEqual(expected, actual)
    def test_shell_sort_case5_empty(self):
        expected = []
        actual = shell_sort([])
        self.assertEqual(expected, actual)
    def test_shell_sort_case6_exists_zero(self):
        expected = [0, 1, 2, 3, 99]
        actual = shell_sort([99, 0, 2, 1, 3])
        self.assertEqual(expected, actual)
    def test_shell_sort_case7_duplicated(self):
        expected = [1, 2, 2, 3, 4, 4, 5]
        actual = shell_sort([5, 1, 2, 4, 4, 3, 2])
        self.assertEqual(expected, actual)

def main():
    arr = [2, 99, 5, 1, 3, 7, 8, 4, 9, 10]
    print(shell_sort(arr)) # Output: [1, 2, 3, 4, 5, 7, 8, 9, 10, 99]


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()