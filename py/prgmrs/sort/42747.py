"""
H-Index
https://school.programmers.co.kr/learn/courses/30/lessons/42747
"""

def solution(citations):
    citations.sort(reverse=True);
    h = 0
    for i, c in enumerate(citations, start=1):
        if c >= i:
            h = i
        else:
            break
    return h

print(solution([3, 0, 6, 1, 5]))

"""
h index 계산법
https://en.wikipedia.org/wiki/H-index
"""