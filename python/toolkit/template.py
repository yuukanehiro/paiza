import unittest
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

if __name__ == '__main__':
    unittest.main()
