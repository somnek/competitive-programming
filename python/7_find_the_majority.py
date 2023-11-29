def majority(l):
    if not l:
        return None

    d = {}
    for i in l:
        if i not in d:
            d[i] = 1
        else:
            d[i] += 1

    d = dict(sorted(d.items(), key=lambda x: x))
    mx = max(x for x in d.values())

    for k, v in d.items():
        if [x for x in d.values()].count(mx) == 1 and v == mx:
            return k

    return None


def sol(arr):
    d = {elem: arr.count(elem) for elem in arr}
    l = [k for k, v in d.items() if v == max(d.values())]

    return l[0] if len(l) == 1 else None


if __name__ == "__main__":
    print(majority([]))  # None
    print(majority([""]))  # None
    print(majority(["a", "b", "c"]))  # None
    print(majority(["a", "b", "c", "c"]))  # c
    print(majority(["d", "a", "b", "c", "a"]))  # a
    print(majority(["a", "b", "b", "a"]))  # None
