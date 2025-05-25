hashTable = {"apple": 1, "banana": 2, "orange": 3}
print(hashTable.keys()) # Output: dict_keys(['apple', 'banana', 'orange'])
print(hashTable.values()) # Output: dict_values([1, 2, 3])
print(hashTable.items()) # Output: dict_items([('apple', 1), ('banana', 2), ('orange', 3)])

print("start for key in hashTable.keys():")
for key in hashTable.keys():
    print(key)
print("end for key in hashTable.keys():")
# Output:
# start for key in hashTable.keys():
# apple
# banana
# orange
# end for key in hashTable.keys():


print("start for key in hashTable.values():")
for v in hashTable.values():
    print(v)
print("end for key in hashTable.values():")
# Output:
# start for key in hashTable.values():
# 1
# 2
# 3
# end for key in hashTable.values():


print("start for key in hashTable.items():")
for key, v in hashTable.items():
    print(key, v)
print("end for key in hashTable.items():")
# Output:
# start for key in hashTable.items():
# apple 1
# banana 2
# orange 3
# end for key in hashTable.items():
