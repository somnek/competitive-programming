# this is correct, but not the right way
# try to solve it without `set()`, it should not
# include dulicates


def permutations(string: str) -> list[str]:
    result = []
    helper(list(string), 0, len(string), result)
    return set(result)


def helper(chars: list[str], start_idx: int, length: int, res: list[str]):
    if start_idx == length:
        res.append("".join(chars))
    else:
        for i in range(start_idx, length):
            chars[start_idx], chars[i] = chars[i], chars[start_idx]
            helper(chars, start_idx + 1, length, res)
            chars[start_idx], chars[i] = chars[i], chars[start_idx]


string = "ABC"
final = permutations(string)
print(final)
assert len(final) == len(set(final))
