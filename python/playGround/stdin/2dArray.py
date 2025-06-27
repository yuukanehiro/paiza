

stdin_input1 = """3 3 1
1 2 3
4 5 6
7 8 9
"""

lines1 = stdin_input1.strip().split("\n")
iter_lines_input1 = iter(lines1).__next__ # 次の行を読む関数
A = [list(map(int, iter_lines_input1().split())) for _ in range(len(lines1))]
print(A) # Output: [[3, 3, 1], [1, 2, 3], [4, 5, 6], [7, 8, 9], [2, 2, 3, 1]]


stdin_input2 = """3 3 1
11 12 13
14 15 16
17 18 19
"""

lines2 = stdin_input2.strip().split("\n")
iter_lines_input2 = iter(lines2).__next__ # 次の行を読む関数
for _ in range(len(lines2)):
    print(iter_lines_input2())

# Output:
# 3 3 1
# 11 12 13
# 14 15 16
# 17 18 19