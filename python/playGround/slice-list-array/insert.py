

# OK
items1 = [11, 12, 13, 15, 16]
result1 = []

result1.extend(items1[:3])
result1.append(14)
result1.extend(items1[3:])
print(result1) # Output: [11, 12, 13, 14, 15, 16] // success

# NG
items2 = [11, 12, 13, 15, 16]
result2 = []

result2.append(items2[:3])
result2.append(14)
result2.append(items2[3:])
print(result2) # Output: [[11, 12, 13], 14, [15, 16]] // fail