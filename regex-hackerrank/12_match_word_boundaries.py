"""
Note:
    \b           assert position at a word boundary
    
    three diferrent positions qualify for word boundaries :
    1. Before the first character in the string, if the first character is a word character.
    2. Between two characters in the string, where one is a word character and the other is not a word character.
    3. After the last character in the string, if the last character is a word character.

Task:
    S
    match word starting with vowel (a,e,i,o, u, A, E, I , O or U).
    consist of letters (lowercase and uppercase both) only.
"""

import re
ss = [
    'Found any match?',
    'A match is starting.',
    'An example is a match.',
    'that indie song',
    'Eggs are a good source of protein.',
]
regex = r'\b[aeiouAEIOU][a-zA-Z]*\b'

for s in ss:
    match = re.search(regex, s)
    print(match)

