import unittest

def unpack_list():
    letters = ['a', 'b', 'c']
    return " ".join(letters)


class TestUnpack(unittest.TestCase):
    def test_unpack_list(self):
        self.assertEqual(unpack_list(), "a b c")

if __name__ == '__main__':
    unittest.main()
