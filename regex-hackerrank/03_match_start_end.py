"""
Note:
    ^ starts
    $ ends

Task:
    string S match Xxxxx, x = words, X = digit
    S must starts with X and end with .
    S size = 6

"""

import re
s = "1qazs."
regex = r'^\d\w{4}\.$'
match = re.match(regex, s)
print(match)
