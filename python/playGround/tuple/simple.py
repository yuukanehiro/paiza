
tuple1 = (1, 2, 3)
print(tuple1) # Output: (1, 2, 3)

print(tuple1[0]) # Output: 1
print(tuple1[1]) # Output: 2
print(tuple1[2]) # Output: 3
print(tuple1[-1]) # Output: 3


tuple2 = {(1, 2): 3}
print(tuple2[(1, 2)]) # Output: 3


tuple3 = {(1, 2)}
print((1,2) in tuple3) # Output: True
