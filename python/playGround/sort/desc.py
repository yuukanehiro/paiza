
# sorted()で降順
numbers = [5, 3, 1, 4, 2]
sortedNumbers = sorted(numbers, reverse=True)
print(sortedNumbers) # Output: [5, 4, 3, 2, 1]

# sort() ... 破壊的にソート
numbers2 = [5, 3, 1, 4, 2]
numbers2.sort(reverse=True)
print(numbers2) # Output: [5, 4, 3, 2, 1]
