
d1 = {"a": "val_a", "b": "val_b", "c": "val_c", "d": "val_d", "e": "val_e"}

for k, v in d1.items():
    print(f"i: {k}, v:{v}")

# Output:
# i: a, v:val_a
# i: b, v:val_b
# i: c, v:val_c
# i: d, v:val_d
# i: e, v:val_e

d2 = {"a2": "val_a2", "b2": "val_b2", "c2": "val_c2", "d2": "val_d2", "e2": "val_e2"}

for k in d2:
    print(f"i: {k}, v:{d2[k]}")

# Output:
# i: a2, v:val_a2
# i: b2, v:val_b2
# i: c2, v:val_c2
# i: d2, v:val_d2
# i: e2, v:val_e2
