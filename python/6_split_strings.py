def solution(s):
  return s and [''.join(x) for x in zip(s[::2], s[1::2])] + len(s)%2 * [f"{s[-1]}_"] or []


def sol1(s):
  import re
  return re.findall(".{2}", s + "_")

def sol2(s):
  result = []
  if len(s)%2:
    s += '_'
  for i in range(0, len(s), 2):
    result.append(s[i:i+2])
  return result


def sol3(s):
  # "🤯"
  return [f'{s}_'[i:i + 2] for i in range(0, len(s), 2)]



# ------------ test ------------
try:
  for func in [sol1, sol2, sol3, solution]:
    assert func("cow_eat_grass") == ['co', 'w_', 'ea', 't_', 'gr', 'as', 's_']
    assert func("asdfadsf") == ['as', 'df', 'ad', 'sf']
    assert func("asdfads") == ['as', 'df', 'ad', 's_']
    assert func("") == []
    assert func("x") == ["x_"]
  print("All tests passed 🍳")
except AssertionError:
  print("WRONG")
