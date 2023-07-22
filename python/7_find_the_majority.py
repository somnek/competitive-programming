def majority(l):
    d = max(list(set([(l.count(x), x) for x in l])), key=lambda x: x[1])
    print(d)


if __name__ == '__main__':
    majority(['a', 'b', 'c', 'a'])
    majority(['a', 'b', 'b', 'a'])
