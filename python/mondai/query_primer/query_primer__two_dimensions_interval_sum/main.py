import sys
import unittest
from collections import defaultdict
from typing import List, Dict


def get_item_price_map(items: List[str]) -> Dict[str, int]:
    item_price_map: Dict[str, int] = {}
    for line in items:
        name, price = line.split()
        item_price_map[name] = int(price)
    return item_price_map


class TestItemPriceMap(unittest.TestCase):
    def test_get_item_price_map(self):
        expected = {"eraser": 50, "pencil": 30}
        actual = get_item_price_map(["eraser 50", "pencil 30"])
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    items = [input().strip() for _ in range(item_count)]
    queries = [input().strip() for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    item_price_map = get_item_price_map(items)

    for item in queries:
        print(item_price_map.get(item, -1))


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()