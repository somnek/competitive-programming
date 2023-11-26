def cakes(need, have):
    return int(all(n in have for n in need)) and min(
        [have[k] // v for k, v in need.items()]
    )


# using dict.get() = setting default value
def sol_1(need, have):
    return min(have.get(k, 0) / need[k] for k in need)


if __name__ == "__main__":
    need = {"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}
    have = {"sugar": 500, "flour": 2000, "milk": 2000}
    print(cakes(need, have))
    print(sol_1(need, have))
