"""
[1차] 비밀지도
https://school.programmers.co.kr/learn/courses/30/lessons/17681?language=python3
"""
def solution(n, arr1, arr2):
    answer = []
    for i in range(n):
        map1 = bin(arr1[i])[2:].zfill(n)
        map2 = bin(arr2[i])[2:].zfill(n)
        str = ""
        for j in range(n):
            if(map1[j] == '0' and map2[j] == '0'):
                str += ' '
            else:
                str += '#'
        answer.append(str)

    return answer

print(solution(5,[9, 20, 28, 18, 11]
,[30, 1, 21, 17, 28]))