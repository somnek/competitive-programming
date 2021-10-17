def strong(n):
  fact = lambda x: x * fact(x - 1) if x > 1 else 1
  return ["Not Strong !!", "Strong!!!"][sum(fact(c) for c in [int(i) for i in str(n)]) == n]

print(strong(144))