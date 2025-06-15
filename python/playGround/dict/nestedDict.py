from collections import defaultdict

input = [
    "A 0000 10000",
    "B 1234 23456",
    "C 5678 98765"
]

bank = defaultdict(dict)

for v in input:
    name, password, amount = v.split()
    bank[name][password] = int(amount)

print(bank.items()) # Output: dict_items([('A', {'0000': 10000}), ('B', {'1234': 23456}), ('C', {'5678': 98765})])

for name in bank:
    for password in bank[name]:
        print(f"name: {name}, password: {password}, amount: {amount}")

# name: A, password: 0000, amount: 98765
# name: B, password: 1234, amount: 98765
# name: C, password: 5678, amount: 98765
