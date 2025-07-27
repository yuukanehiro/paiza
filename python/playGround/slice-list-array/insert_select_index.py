


items1 = [1, 2, 3, 4, 6, 7, 8, 9]
taraget = 5
target_index = 4

res = items1[0:target_index] + [5] + items1[target_index:]
print(res) # Output: [1, 2, 3, 4, 5, 6, 7, 8, 9]