"""
Note:
    [^] matches any char that is not in []

Task:
    S size = 6
    1st char: not a digit(0 till 9)
    2nd char: not lowercase (aeiou)
    3rd char: not in (bcDF)
    4th char: not white spcace char (\r, \n, \t, \f ,<space>)
    5th char: not upper (AEIOU)
    6th char: not (. ,)
"""

import re
s = "0aD\nA."
regex = r'^[^0-9][^aeiou][^bcDF][^\n\t\r\f\s][^AEIOU][^.,]$'
match = re.match(regex, s)
print(match)

