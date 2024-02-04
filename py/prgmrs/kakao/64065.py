"""
튜플
https://school.programmers.co.kr/learn/courses/30/lessons/64065
"""
def solution(s):
    s = s[2:-2]
    arr = s.split('},{')

    arr.sort(key=lambda s: len(s))
    #print(arr)
    order = dict()
    cnt = 0
    for a in arr:
        i = a.split(',')
        for ele in i:
            if not (ele in order):
                order[int(ele)] = cnt
                cnt +=1
    return list(order.keys())

print(solution("{{2},{2,1},{2,1,3},{2,1,3,4}}"))
print(solution("{{1,2,3},{2,1},{1,2,4,3},{2}}"))
