"""
[1차] 뉴스 클러스터링
https://school.programmers.co.kr/learn/courses/30/lessons/17677
"""
from collections import Counter
def solution(str1, str2):

    str1 = str1.lower()
    str2 = str2.lower()
    a_1 = list()
    a_2 = list()
    for i in range(0, len(str1)-1):
        if str1[i].isalpha() and str1[i+1].isalpha():
            a_1.append(str1[i:i+2])
        
    for i in range(0, len(str2)-1):
        if str2[i].isalpha() and str2[i+1].isalpha():
            a_2.append(str2[i:i+2])

    교집합 = list((Counter(a_2) & Counter(a_1)).elements())
    합집합 = list((Counter(a_2) | Counter(a_1)).elements())
    # print(교집합)
    # print(합집합)
    idx = 1
    if len(교집합) < len(합집합):
        idx = len(교집합) / len(합집합)
    
    return int(idx * 65536)


print(solution("FRANCE","french"))
print(solution("handshake","shake hands"))
print(solution("aa1+aa2","AAAA12"))
print(solution("E=M*C^2","e=m*c^2"))