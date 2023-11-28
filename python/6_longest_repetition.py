def longest_repetition(s):
    mx, c, alp = 0, 1, ""
    for i in range(len(s) - 1):
        c = (s[i] == s[i + 1]) and c + 1 or 1
        if c > mx:
            mx, alp = c, s[i]
    return (alp, mx)


if __name__ == "__main__":
    tests = [
        "aaaabb",  # ('a', 4)
        "bbbaaabaaaa",  # ('a', 4)
        "cbdeuuu900",  # ('u', 3)
        "abbbbb",  # ('b', 5)
        "aabb",  # ('a', 2)
        "ba",  # ('b', 1)
        "",  # ('', 0)
    ]

    for t in tests:
        print(longest_repetition(t))
