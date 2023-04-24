"""
Note:
    a `Class` = set of char in []
    - is range [a-z] etc

Task:
    S size >= 5
    1st char: lowercase alphabet
    2nd char: positive digit (0 aint positive)
    3rd char: not lowercase alpha
    4th char: not uppercase alpha
    5th char: uppercase alpha
"""

import re
s = "a0$?ZWe"
regex = r'^[a-z][1-9][^a-z][^A-Z][A-Z].*$'
match = re.match(regex, s)
print(match)

