"""
Note:
    by default, the . (dot) matches any character except a newline.
    had to add ^ and $ to match the entire string

    \s\S matches any character (except for a newline)

Task:
    match string in the form: abc.def.ghi.jkx (can be any character/number)
        except newline

"""

import re
# s = "123.456.a,c.*ef"
s = "123.123.123.132.123.123" # this should return False
regex = r'^[\s\S]{3}\.[\s\S]{3}\.[\s\S]{3}\.[\s\S]{3}$'
match = re.match(regex, s)
print(match)
