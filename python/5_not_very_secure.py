alphanumeric = lambda x: 1 & x.isalnum() | 0

print(alphanumeric("PassW0rd"))
print(alphanumeric("hello world_"))
print(alphanumeric("  "))