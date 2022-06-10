find_missing_letter = lambda l: chr([ord(l[i]) for i in range(len(l)-1) if ord(l[i+1])-ord(l[i]) != 1][0]+1)

assert(find_missing_letter(['a','b','c','d','f']) == 'e')
assert(find_missing_letter(['O','Q','R','S']) == 'P')
