def number_of_pairs(l):
    return sum(l.count(x)//2 for x in set(l))


assert number_of_pairs(["red","red"]) == 1
assert number_of_pairs(["red","green","blue"]) == 0
assert number_of_pairs(["gray","black","purple","purple","gray","black"]) == 3
assert number_of_pairs([]) == 0
assert number_of_pairs(["red","green","blue","blue","red","green","red","red","red"]) == 4