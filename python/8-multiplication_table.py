def multi_table(n):
    return "\n".join([f"{i} * {n} = {n * i}" for i in range(1, 11)])

    
    
print(multi_table(12))