def find_even_index(l: list):
    return ([i for i in range(len(l)) if sum(l[:i]) == sum(l[i+1:])] + [-1])[0]


if __name__ == '__main__':
    assert find_even_index([1,2,3,4,3,2,1]) == 3
    assert find_even_index([1,100,50,-51,1,1]) == 1
    assert find_even_index([1,2,3,4,5,6]) == -1
    assert find_even_index([20,10,30,10,10,15,35]) == 3
    assert find_even_index([20,10,-80,10,10,15,35]) == 0
    assert find_even_index([-1,-2,-3,-4,-3,-2,-1]) == 3
    assert find_even_index(list(range(-100,-1))) == -1
    assert find_even_index([-4, -2, 2, -6]) == 1
