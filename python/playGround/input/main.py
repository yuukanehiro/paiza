import unittest
from unittest.mock import patch
from typing import List, Dict

class TestInputPatterns(unittest.TestCase):
    @patch('builtins.input', return_value='abcde') # returnValueだとエラーになる。
    def test_string_input(self, mockInput): # @patch('builtins.input', return_value=...) を使うと、テスト関数の引数にモックオブジェクトが自動で渡されるようになる
        s = input()
        self.assertEqual(s, 'abcde')

    @patch('builtins.input', return_value='abcde')
    def testListStringInput(self, mockInput):
        s = list(input())
        self.assertEqual(s, ['a', 'b', 'c', 'd', 'e'])

    @patch('builtins.input', return_value='5')
    def testSingleIntInput(self, mockInput):
        a = int(input())
        self.assertEqual(a, 5)

    @patch('builtins.input', return_value='1 2')
    def testTwoIntsInput(self, mockInput):
        x, y = map(int, input().split())
        self.assertEqual((x, y), (1, 2))

    @patch('builtins.input', return_value='1 2 3 4 5')
    def testSplitStringList(self, mockInput):
        li = input().split()
        self.assertEqual(li, ['1', '2', '3', '4', '5'])

    @patch('builtins.input', return_value='1 2 3 4 5')
    def testIntList(self, mockInput):
        li = list(map(int, input().split()))
        self.assertEqual(li, [1, 2, 3, 4, 5])

    @patch('builtins.input', return_value='FFFTFTTFF')
    def testSplitByT(self, mockInput):
        li = input().split('T')
        self.assertEqual(li, ['FFF', 'F', '', 'FF'])

    # 辞書型の受け取り
    @patch('builtins.input', side_effect=['2 Yuu', '3 Ayaka'])
    def testSplitByT(self, mockInput):
        li: Dict[str, str] = dict(input().split() for _ in range(2))
        self.assertEqual(li, {"2": "Yuu", "3": "Ayaka"})

if __name__ == '__main__':
    unittest.main()
