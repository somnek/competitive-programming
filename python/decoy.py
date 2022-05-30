(lambda o, i, b , a, p: (
    lambda s: [
        p(s), 
        (lambda u: p("\n".join([f"{o}{i*a:^{u}}{o}" for _ in " " * a])))(b-2),
        p(s[::-1])])(o*b+("\n"if b-a-2 else"")+"\n".join(o+" "*(b-2)+o for _ in" "*((b-a)//2-1))))("#","~",7,3,print)