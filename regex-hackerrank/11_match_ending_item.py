"""
Note:
    $           "Boundary Matcher", match an occurrence of char/char class/group  at the end of line

Task:
    S
    only have letters(upper/lower)
    end in 's'
"""

import re
ss = [
    'This is a test string',
    '012 a test string',
    '012s',
    'Abcs',
]
regex = r'^[a-zA-Z]*s$'
for s in ss:
    match = re.match(regex, s)
    print(match)

