"""
Note:
    {x}   Exactly x repetitions of a char/charclass/groups

    w{3}        matches 'www'
    [xyz]{5}    matches 'xxxxy' or 'zzzzz' or 'yyyyy' or 'xyxyz'
    \d{4}       matches digit Exactly 4 times

Task:
    S size = 45
    first 40 char: letters/digits
    last   5 char: odd digits/ whitespace
"""

import re
s = 'a' * 20 + '2' * 20 + '1  39'
regex = r'[a-zA-Z02468]{40}[\s13579]{5}'
match = re.match(regex, s)
print(match)

