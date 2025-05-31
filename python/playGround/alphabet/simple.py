import string
import unittest


def getAlphabetString():
    return string.ascii_lowercase

def getAlphabetDict():
    return {char: 0 for char in string.ascii_lowercase}


class TestAlphabet(unittest.TestCase):
    def test_getAlphabetString(self):
        self.assertEqual(getAlphabetString(), "abcdefghijklmnopqrstuvwxyz")

    def test_getAlphabetDict(self):
        self.assertEqual(getAlphabetDict(), {'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0})
        self.assertEqual(" ".join(getAlphabetDict().keys()), "a b c d e f g h i j k l m n o p q r s t u v w x y z")

if __name__ == '__main__':
    unittest.main()
