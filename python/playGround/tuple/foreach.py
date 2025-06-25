t = [("a", "val_a"), ("b", "val_b")]

for k, v in t:
    print(f"k: {k}, v: {v}")

# Output:
# k: a, v: val_a
# k: b, v: val_b


t2 = [("a1", "b1", "c1"), ("a2", "b2", "c2")]
for a, b, c in t2:
    print(f"a: {a}, b: {b}, c: {c}")

# Output:
# a: a1, b: b1, c: c1
# a: a2, b: b2, c: c2