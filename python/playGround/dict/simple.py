hashTable = {}
hashTable["apple"] = 1
hashTable["banana"] = 2
print(hashTable) # Output: {'apple': 1, 'banana': 2}
print(len(hashTable)) # Output: 2

hashTable["banana"] = 3
print(hashTable) # Output: {'apple': 1, 'banana': 3}


hashTable2 = {"apple": 1, "banana": 2, "orange": 3}
print("apple" in hashTable2)  # Output: True
print("banana" in hashTable2) # Output: True
print("grape" in hashTable2)  # Output: False

if "banana" in hashTable:
    print("'banana' exists in hash_table") # Output: 'banana' exists in hash_table
