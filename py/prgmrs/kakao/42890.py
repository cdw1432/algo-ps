"""
후보키
https://school.programmers.co.kr/learn/courses/30/lessons/42890
"""

from itertools import combinations as combination

def solution(relation):
    row = len(relation)
    col = len(relation[0])
    cases = []
    for i in range(1, col+1):
        cases.extend(combination(range(col),i))

    #유일성 체크
    unique = []
    for case in cases:
        hm = dict()
        promote = True
        for j in range(row):
            str = ''
            for i in case:
                str += relation[j][i]
            #print(str)
            if str in hm:
                promote = False
                break
            else:
                hm[str] = True
        if promote == True:
            unique.append(case)

    #최소성 체크
    minimal = set(unique)
    for i in range(len(unique)):
        for j in range(i+1, len(unique)):
            if len(unique[i]) == len(set(unique[i]) & set(unique[j])):
                minimal.discard(unique[j])
    #print(minimal)
    return len(minimal)


print(solution([["100","ryan","music","2"],["200","apeach","math","2"],["300","tube","computer","3"],["400","con","computer","4"],["500","muzi","music","3"],["600","apeach","music","2"]]))