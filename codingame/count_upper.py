"""
i wont btw: (my solution)
    print(sum([int(x.isupper()) for x in _input]))
"""
# solution from stackoverflow 1
len([letter for letter in input_s if letter.isupper()])

# solution from stackoverflow 2
import re
len(re.findall(r'[A-Z]', string)
