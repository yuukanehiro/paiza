
# sequence[start:end:step]
# start：スライスの開始位置を指定します。この位置の要素を含みます。
# end：スライスの終了位置を指定します。この位置の要素は含みません。
# step（省略可）：スライスのステップ数を指定します。デフォルトでは1です。

numbers = numbers = [1, 2, 3, 4, 5]

print(numbers) # Output: [1, 2, 3, 4, 5]
print(numbers[:]) # Output: [1, 2, 3, 4, 5]

print(numbers[1:4]) # Output: [2, 3, 4]

print(numbers[:3]) # Output: [1, 2, 3]

print(numbers[2:]) # Output: [3, 4, 5]

print(numbers[::1]) # Output: [1, 2, 3, 4, 5]
print(numbers[::2]) # Output: [1, 3, 5]
print(numbers[::-1]) # Output: [5, 4, 3, 2, 1]
