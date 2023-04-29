"""
Note:
    +           match 1 or more repetitions of char/char class/group

    w           match w 1 or more times
    [xyz]+      match x, y, or z 1 or more times
    \d+         match a digit 1 or more times

Task:
    S
    begin with 1 or more digits
    after that, 1 or more uppercase letters
    End with 0 or more lowercase letters
"""

import re
s = '1Q'
regex = r'^\d+[A-Z]{1,}[a-z]+$'
match = re.match(regex, s)
print(match)

