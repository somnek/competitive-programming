def solution(s):
    #bool(len(s)%2) * '_'
    return s and [x for x in zip(s[::2], s[1::2])]



print(solution('cowdone'))



tests = (
    ("asdfadsf", ['as', 'df', 'ad', 'sf']),
    ("asdfads", ['as', 'df', 'ad', 's_']),
    ("", []),
    ("x", ["x_"]),
)