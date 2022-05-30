def find_the_odd_int(seq):
  return [d for d in seq if seq.count(d)%2][0]

def sol1(seq):
  res = 0
  for x in seq:
    res ^= x
  return res

def sol2(seq):
  for i in seq:
    if seq.count(i)%2:
      return i

x = [1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1]
print(find_the_odd_int(x))
