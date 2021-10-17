def even_or_odd(n):
    return ["Even", "Odd"][int("{0:b}".format(1)) & 1]
    

# sol 1
def sol1(n):
    return ["Even", "Odd"][n % 2]
 

print(even_or_odd(2))
print(sol1(2))