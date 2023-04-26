"""
Note:
    {x, y}   Between x and y repetitions of a char/charclass/groups

    w{3, 5}     matches 'w' 3 or 4 or 5 times
    [xyz]{5,}   matches 'x' or 'y' or 'z' 5 or more times
    \d{1, 4}    matches digits between 1 and 4 times

Task:
    S
    begin with 1 or 2 digits
    after that, S should have 3 or more letters
    Then, S should end with up to 3 '.' symbol(s). 
    You can end with 0 to 3 '.' symbol(s), inclusively.
"""

import re
ss = [
    '99abc...',
    '99abc..',
    '9abcdef..',
    'abf',
    '3threeormorealphabets8',
]
regex = r'^\d{1,2}[a-zA-Z]{3,}\.{0,3}$'
for s in ss:
    match = re.match(regex, s)
    print(match)

