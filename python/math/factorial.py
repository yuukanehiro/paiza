import unittest


def factorial(n):
    """
    階乗
    例) 5! = 5*4*3*2*1
    """
    if n == 1:
        return 1
    return n * factorial(n -1)


class TestFruitCounter(unittest.TestCase):
    def testFactorial2(self):
        self.assertEqual(factorial(2), 2) # 2! = 2*1 = 2

    def testFactorial3(self):
        self.assertEqual(factorial(3), 6) # 3! = 3*2*1 = 6

    def testFactorial4(self):
        self.assertEqual(factorial(4), 24) # 4! = 4*3*2*1 = 24

    def testFactorial5(self):
        self.assertEqual(factorial(5), 24) # 5! = 5*4*3*2*1 = 120

if __name__ == '__main__':
    unittest.main()
