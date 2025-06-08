
# List[str] -> List[int]
# list & map
items1 = ["1", "2", "3", "39", "4", "5"]
int_items1 = list(map(int, items1))
print(int_items1) # Output: [1, 2, 3, 39, 4, 5]
# for
items2 = ["10", "20", "30", "390", "40", "50"]
int_items2 = [int(v) for v in items2]
print(int_items2) # Output: [10, 20, 30, 390, 40, 50]

# List[int] -> List[str]
# list & map
items3 = [1, 2, 3, 39, 4, 5]
str_items3 = list(map(str, items3))
print(str_items3) # Output: ['1', '2', '3', '39', '4', '5']
# for
items4 = [10, 20, 30, 390, 40, 50]
str_items4 = [str(v) for v in items4]
print(str_items4) # Output: ['10', '20', '30', '390', '40', '50']
