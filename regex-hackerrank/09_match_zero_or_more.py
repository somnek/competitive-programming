"""
Note:
    *    Zero or more repetitions of a char/charclass/groups

    w*          matches 'w' 0 or more times
    [xyz]*      matches 'x' or 'y' or 'z' 0 or more times
    \d*         matches digits 0 or more times

Task:
    S
    begin with 2 or more digits
    after that, 0 or more lowercase letters
    End with 0 or more uppercase letters
"""

import re
s = '00000abcA'
regex = r'\d{2,}[a-z]*[A-Z]*'
match = re.match(regex, s)
print(match)

