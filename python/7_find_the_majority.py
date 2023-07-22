def majority(arr):
    if not len(arr):
        return None
    d = {}
    for e in arr:
        count = arr.count(e)
        if count not in d:
            d[count] = e
        else:
            if e in d[count]:
                continue
            d[count] += e
    s = sorted(d.items())[-1]
    return s[1][0] if len(s[1]) == 1 else None



if __name__ == '__main__':
    print(majority([]))
    print(majority(['']))
    print(majority(['a', 'b', 'c']))
    print(majority(['d', 'a', 'b', 'c', 'a']))
    print(majority(['a', 'b', 'b', 'a']))
