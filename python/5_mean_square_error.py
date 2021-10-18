def solution(a1, a2):
  return int(sum((z - x) ** 2 for z, x in zip(a1, a2))/len(a1))

a1 = [1,2,3]
a2 = [4,5,6]
print(solution(a1, a2))