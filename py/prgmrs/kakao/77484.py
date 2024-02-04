"""
The Highest and Lowest Lotto Prizes
https://school.programmers.co.kr/learn/courses/30/lessons/77484
"""
def solution(lottos, win_nums):
    win_map = dict()
    for i in win_nums:
        win_map[i] = True

    scribble = 0
    cnt = 0
    for i in lottos:
        if i == 0:
            scribble += 1
        if i in win_map:
            cnt +=1

    print(win_map)
    print(scribble)
    print(cnt)
    max_ = 6 - (scribble + cnt) + 1
    min_ = 6 - (cnt) + 1
    if max_ >= 6:
        max_ = 6
    if min_ >= 6:
        min_ = 6
    return [max_,min_]

print(solution([44, 1, 0, 0, 31, 25],[31, 10, 45, 1, 6, 19]	))