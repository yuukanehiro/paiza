
# List[str] -> List[int]
items = ["1", "2", "3", "39", "4", "5"]
int_items = list(map(int, items))
print(int_items) # Output: [1, 2, 3, 39, 4, 5]

# List[int] -> List[str]
items = [1, 2, 3, 39, 4, 5]
str_items = [str(v) for v in items]
print(str_items) # Output: ['1', '2', '3', '39', '4', '5']
