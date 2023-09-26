def count_deaf_rats(town: str) -> int:
    total = 0
    l, r = "O~", "~O"
    ss = town.replace(" ", "")
    l_side, r_side = ss.split("P")[0], ss.split("P")[1]
    l_split = [l_side[i:i+2] for i in range(0, len(l_side), 2)]
    r_split = [r_side[i:i+2] for i in range(0, len(r_side), 2)]
    total += l_split.count(l)
    total += r_split.count(r)
    return total


# clever
def sol_1(town: str) -> int:
    return town.replace(' ', '')[::2].count('O')

# regex
def sol_2(town: str) -> int:
    import re
    return re.findall(r"O~|~O", town.split("P")[0]).count("O~") + \
           re.findall(r"O~|~O", town.split("P")[1]).count("~O") 


if __name__ == "__main__":
    assert count_deaf_rats("~O~O~O~OP~O~OO~") == 2
    assert sol_1("~O~O~O~OP~O~OO~") == 2
    assert sol_2("~O~O~O~OP~O~OO~") == 2

