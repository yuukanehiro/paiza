import sys
import unittest
from collections import defaultdict
from typing import List


# def get_result_list(items: List[str]) -> List[str]:
#     superchat_sum_map = defaultdict(int)
#     memberships = []
#     response = []

#     for v in items:
#         arr = v.split()
#         name = arr[0]
#         query = arr[1]
#         if query == "give":
#             if name in superchat_sum_map:
#                 superchat_sum_map[name] += int(arr[2])
#             else:
#                 superchat_sum_map[name] = int(arr[2])
#         elif query == "join":
#             memberships.append(name)

#     # ソート条件:
#     # - superchat 合計金額の降順
#     # - 金額が同じならアカウント名の辞書降順
#     sorted_items = sorted(
#         superchat_sum_map.items(),
#         key=lambda x: (-x[1], tuple(-ord(c) for c in x[0]))
#     )
#     memberships.sort()

#     for k, _ in sorted_items:
#         response.append(k)

#     for v in memberships:
#         response.append(v)

#     return response


# class TestResultList(unittest.TestCase):
#     def test_get_result_list(self):
#         expected = [
#             "yoyo",
#             "aiueo",
#             "so_cute",
#             "coffee_addiction",
#             "kk"
#         ]
#         input = [
#             "aiueo give 2489 !",
#             "kk join membership!",
#             "coffee_addiction join membership!",
#             "so_cute give 837 !",
#             "yoyo give 9284 !"
#         ]
#         actual = get_result_list(input)
#         self.assertEqual(expected, actual)


def main():
    # item_count = int(input())
    H, W, N = map(int, input().strip().split())
    # items = [input().strip() for _ in range(item_count)]
    # queries: Dict[int, str] = {int(line.split()[0]): line.split()[1] for line in (input().strip() for _ in range(query_count))}
    # queries: List[Tuple[int, str]] = [(int(line.split()[0]), line.split()[1]) for line in (input().strip() for _ in range(query_count))]

    A = []
    for _ in range(H):
        row = list(map(int, input().strip().split()))
        A.append(row)

    


if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()
