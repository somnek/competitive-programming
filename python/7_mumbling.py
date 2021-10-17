def accum(s):
    return "-".join((c * (i + 1)).capitalize() for i, c in enumerate(s))

# solution
def accum(s):
    return '-'.join((a * i).title() for i, a in enumerate(s, 1))

print(accum("hello"))