import unittest
import sys
from collections import defaultdict

def create_fruit_counter():
    hashTable = defaultdict(int)
    hashTable["apple"] += 1
    hashTable["banana"] += 1
    hashTable["apple"] += 1
    return hashTable

class TestFruitCounter(unittest.TestCase):
    def test_apple_count(self):
        counter = create_fruit_counter()
        self.assertEqual(counter["apple"], 2)

    def test_banana_count(self):
        counter = create_fruit_counter()
        self.assertEqual(counter["banana"], 1)

    def test_orange_default(self):
        counter = create_fruit_counter()
        self.assertEqual(counter["orange"], 0)  # 存在しないキーでもdefaultdictなら0

def main():
    got = create_fruit_counter()
    for k, v in got.items():
        print(f"{k} {v}")

if __name__ == '__main__':
    if len(sys.argv) > 1 and sys.argv[1] == "test":
            unittest.main(argv=sys.argv[:1])
    else:
        main()

# テストを実行する場合
# % python3 template.py test
# ...
# ----------------------------------------------------------------------
# Ran 3 tests in 0.000s

# OK

# 通常の実行
# % python3 template.py
# apple 2
# banana 1
