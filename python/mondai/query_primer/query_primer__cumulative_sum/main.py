import sys
import unittest
from typing import List


def get_cumulative_sum_list(items: List[int]) -> List[int]:
    sums = [0] * (len(items) + 1)

    for i, v in enumerate(items):
        sums[i + 1] = sums[i] + v

    return sums


class TestItemPriceMap(unittest.TestCase):
    def test_get_cumulative_sum_list(self):
        expected = [
            0,
            69,
            81,
            109
        ]
        input_items = [
            69,
            12,
            28
        ]
        actual = get_cumulative_sum_list(input_items)
        self.assertEqual(expected, actual)


def main():
    item_count, query_count = map(int, input().split())
    # item_count = int(input()) # 1つのintの場合
    items = [int(input().strip()) for _ in range(item_count)]
    queries = [int(input().strip()) for _ in range(query_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    sums = get_cumulative_sum_list(items)

    for q in queries:
        print(sums[q])


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()
