from rich.progress import track

def unique_in_order(s):
    return [s[i] for i in range(len(s)-1) if s[i] != s[i + 1]] + [s[-1]] if s else []

def sol1(s):
    """
    it's an "abuse" of the boolean shortcut:

    check if x in r[-1:], if so, the or condition is fulfilled so the second statement isn't executed
    if x is not in the end part, the second statement is checked and x is appended to the r list.
    Note that the or condition returns a boolean, but this one isn't used.

    ! so if x in r[-1:-1] is True, the second part is alwasy ignored
    """

    r = []
    for x in s:
        x in r[-1:] or r.append(x)
    return r


def sol2(s):
    from itertools import groupby
    return [x for x, _ in groupby(s)]

def sol3(s):
    return [v for i, v in enumerate(s) if i == 0 or v != s[i - 1]]

# test case
test = [sol3('AAAABBBCCDAABBB') == ['A', 'B', 'C', 'D', 'A', 'B'],
        sol3('soddi') == ['s', 'o', 'd', 'i'],
        sol3('') == []]
for t in track(test, description='test case'):
    assert(t)