"""
최소직사각형
https://school.programmers.co.kr/learn/courses/30/lessons/86491
"""

def solution(sizes):
    vertical = []
    horizontal = []
    for i in sizes:
        if i[0] > i[1]:
            vertical.append(i[0])
            horizontal.append(i[1])
        else:
            vertical.append(i[1])
            horizontal.append(i[0])
    return max(vertical) * max(horizontal)


tests = [
    [[60, 50], [30, 70], [60, 30], [80, 40]],
    [[10, 7], [12, 3], [8, 15], [14, 7], [5, 15]],
    [[14, 4], [19, 6], [6, 16], [18, 7], [7, 11]]
]
for i in tests:
    print(solution(i))

"""
큰 수, 작은수 정리 후
max num(가로) * max num(세로)
"""