import sys
import unittest
from collections import defaultdict
from typing import List, Dict, Tuple

# ------------------------
# ハッシュ探索(dict)した結果をboolで返却
# ------------------------
def hash_search_by_dict(d: Dict[str, int], target: str) -> bool:
    if target in d:
        return True
    else:
        return False

# ------------------------
# テスト用コード（unittest）
# ------------------------
class TestHashSearchByDict(unittest.TestCase):
    def test_hash_search_by_dict_found(self):
        input_dict: Dict[str, int] = {"apple": 10, "banana": 20, "cherry": 30}
        input_target: str = "banana"
        expected = True
        actual = hash_search_by_dict(input_dict, input_target)
        self.assertEqual(expected, actual)

    def test_hash_search_by_dict_not_found(self):
        input_dict: Dict[str, int] = {"apple": 10, "banana": 20, "cherry": 30}
        input_target: str = "grape"
        expected = False
        actual = hash_search_by_dict(input_dict, input_target)
        self.assertEqual(expected, actual)

    def test_hash_search_by_dict_empty(self):
        input_dict: Dict[str, int] = {}
        input_target: str = "apple"
        expected = False
        actual = hash_search_by_dict(input_dict, input_target)
        self.assertEqual(expected, actual)

    def test_hash_search_by_dict_single_element_true(self):
        input_dict: Dict[str, int] = {"mango": 42}
        input_target: str = "mango"
        self.assertTrue(hash_search_by_dict(input_dict, input_target))

    def test_hash_search_by_dict_single_element_false(self):
        input_dict: Dict[str, int] = {"mango": 42}
        input_target: str = "melon"
        self.assertFalse(hash_search_by_dict(input_dict, input_target))

# ------------------------
# main()
# 起動時分岐
# ------------------------
def main():
    d: Dict[str, int] = {
        "satou": 80,
        "suzuki": 99,
        "tanaka": 50
    }

    target_suzuki: str = "suzuki"
    print(hash_search_by_dict(d, target_suzuki)) # True
    target_tanaka: str = "tanaka"
    print(hash_search_by_dict(d, target_tanaka)) # True
    target_unknown: str = "hoge"
    print(hash_search_by_dict(d, target_unknown)) # False

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
        unittest.main(argv=sys.argv[:1])
    else:
        main()