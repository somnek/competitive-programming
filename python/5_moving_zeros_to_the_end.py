def move_zeros(l):
    return [x for x in l if x] + [0] * l.count(0)

def sol_1(l):
    # lambda x: not x -> put all False vals behind True vas
    return sorted(l, key=lambda x: not x)

if __name__ == '__main__':
    l = [1, 0, 1, 2, 0, 1, 3]
    want = [1, 1, 2, 1, 3, 0, 0]
    have = move_zeros(l)
    print(have)
    assert(have == want), f"expect {want}, got {have}"
