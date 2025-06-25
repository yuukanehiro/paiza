import unittest

def getPrimers(n):
    """
    素数取得
    """
    res = []
    for i in range(n+1):
        if isPrimer(i):
            res.append(i)
    return res

def isPrimer(n):
    if n <= 1:
        return False

    for i in range(2, n+1):
        if n == i:
            return True

        if n % i == 0:
            return False

class TestgetPrimers(unittest.TestCase):
    def testGetPrimers1(self):
        self.assertEqual(getPrimers(1), [])

    def testGetPrimers2(self):
        self.assertEqual(getPrimers(2), [2])

    def testGetPrimers3(self):
        self.assertEqual(getPrimers(3), [2,3])

    def testGetPrimers4(self):
        self.assertEqual(getPrimers(3), [2,3])

    def testGetPrimers5(self):
        self.assertEqual(getPrimers(5), [2,3,5])

if __name__ == '__main__':
    unittest.main()
