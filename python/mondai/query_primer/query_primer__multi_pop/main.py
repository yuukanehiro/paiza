import sys
import unittest
from collections import deque
from typing import List

def get_merged_list(items: deque[str], queries: List[str]) -> List[str]:
    result = []
    for q in queries:
        if q == "pop":
            items.popleft()
        elif q == "show":
            result.extend(items)

    return result


class TestItemPriceMap(unittest.TestCase):
    def test_get_merged_list_case1(self):
        expected = ["2410", "9178", "7252"]
        actual = get_merged_list(deque(["7564", "4860", "2410", "9178", "7252"]), ["pop", "pop", "show"])
        self.assertEqual(expected, actual)
    def test_get_merged_list_case2(self):
        expected = ["1339", "4960", "3926", "9816", "3018", "4213", "9816", "3018", "4213"]
        actual = get_merged_list(deque(["1005", "2716", "7856", "8546", "1339", "4960", "3926", "9816", "3018", "4213"]), ["pop", "pop", "pop", "pop", "show", "pop", "pop", "pop", "show", "pop"])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    items = deque(input().strip() for _ in range(item_count))
    queries = [input().strip() for _ in range(query_count)]

    for v in get_merged_list(items, queries):
        print(v)


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()