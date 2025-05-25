def outerFunction(x):
    def innerFunction(y):
        return x + y
    return innerFunction(5)

print(outerFunction(10)) # Output: 15
