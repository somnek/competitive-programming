def sol_1(l):
    return max([x for x in set(l) if l.count(x) == max(l.count(v) for v in l)])

def sol_2(l):
    return max(sorted(l, reverse=True), key=l.count)


# run tests
tests = [
    [[12, 10, 8, 12, 7, 6, 4, 10, 12], 12],
    [[12, 10, 8, 12, 7, 6, 4, 10, 12, 10], 12],
    [[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10], 3]
]

for t in tests:
    assert sol_1(t[0]) == t[1]
    assert sol_2(t[0]) == t[1]
