def capitals(word: str) -> list:
    return [i for i, w in enumerate(word) if w.isupper()]

if __name__ == '__main__':
    test_case = ["CodEWaRs", "aAbBcC", "4ysdf4", "ABCDEF", "AB" ]
    for t in test_case:
        print(capitals(t))
