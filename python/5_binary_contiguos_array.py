def binarray(l):
    count, max_l = 0, 0
    hash = {0: -1}

    for i, v in enumerate(l):
        count += 1 if v else -1
        if count not in hash:
            hash[count] = i
        else:
            max_l = max(i - hash[count], max_l)
    return max_l


a = [1, 1, 0, 1, 1, 0, 1, 1]
b = [1, 0]

print(binarray(a))
print(binarray(b))
