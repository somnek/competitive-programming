"""
Note:
    [] only matches one character in []

Task:
    S size = 6
    1st char: 1, 2 or 3
    2nd char: 1, 2 or 0
    3rd char: x, s or 0
    4th char: 3, 0 , A or a
    5th char: x, s or u
    6th char: . or ,
"""

import re
s = "12xAu.x"
regex = r'^[123][120][xs0][30A][xsu][.,]$'
match = re.match(regex, s)
print(match)
