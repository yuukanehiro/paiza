
items1 = [11, 12, 13, 15, 16]
result = []
result.extend(items1[:3])
print(result) # [11, 12, 13]

result.append(14)
print(result) # [11, 12, 13, 14]

result.extend(items1[3:])
print(result) # [11, 12, 13, 15, 16]
