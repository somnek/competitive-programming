#find_missing_letter = lambda l: chr([ord(l[i]) for i in range(len(l)-1) if ord(l[i+1])-ord(l[i]) != 1][0]+1)
def find_missing_letter(l):
    l = sum([ord(c) for c in l])/len(l)
    print(int(l)+2)
    

find_missing_letter(['a','b','c','d','f'])
#assert(find_missing_letter(['a','b','c','d','f']) == 'e')
#assert(find_missing_letter(['O','Q','R','S']) == 'P')
