"""
Note:
    \w matches any alphanumeric character [a-zA-Z0-9_]
    \W will match any non-word character.

    this task didnt say u have to ignore newline, so i didnt add ^ and $
    can do \w instead of \w{1}

    [\w] or \w both works for this task

Task:
    match xxxXxxxxxxxxxxXxxx 
    (x is any word character, X is any non-word)

"""

import re
s = "abc$defghijklm*nop"

regex = r'\w{3}\W{1}\w{10}\W{1}\w{3}'
match = re.match(regex, s)
print(match)

