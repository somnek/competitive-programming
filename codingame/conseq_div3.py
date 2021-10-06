"""
calculate sum of 1 or 2 consequtive number that are divisible by 3
"""


# ------- solution 1---------
n = list(map(int,input()))
s = 0
p = 0

for i in n:
     if i % 3 == 0:
         s += i
     if p and (p + i) % 3 == 0:
         s += p + i
     p = i * 10

print(s)



