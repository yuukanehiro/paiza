

twoDArray = [[0] * 2 for _ in range(3)]
print(twoDArray) # Output: [[0, 0], [0, 0], [0, 0]]

# ===========
# Bad
# ===========
# 次のように*演算子を使用して直接リストを複製すると、
# すべての内側のリストが同じ参照（つまり同じオブジェクト）を指すことになります。
# そのため、一つの内側のリストを変更すると、
# 他のすべての内側のリストも同様に変更されてしまいます。
# このため、2次元配列を作成する際にはこの方法を避けるべきです。
badTwoDArray = [[0] * 2] * 3
print(badTwoDArray) # Output: [[0, 0], [0, 0], [0, 0]]

print(id(twoDArray[0]) != id(twoDArray[1])) # Output: True
print(id(twoDArray[0]) == id(twoDArray[1])) # Output: False

# 全行が同じ1つのリストを共有してしまっている...
print(id(badTwoDArray[0]) == id(badTwoDArray[1])) # Output: True

badTwoDArray = [[0] * 2] * 3
badTwoDArray[0][0] = 10
badTwoDArray[0][1] = 11
print(badTwoDArray) # Output: [[10, 11], [10, 11], [10, 11]]

twoDArray = [[0] * 2 for _ in range(3)]
twoDArray[0][0] = 10
twoDArray[0][1] = 11
print(twoDArray) # Output: [[10, 11], [0, 0], [0, 0]]
