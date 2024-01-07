"""
K번째수
https://school.programmers.co.kr/learn/courses/30/lessons/42748?language=python3
"""
def solution(array, commands):
    answer = []
    for a in range (len(commands)):
        i = commands[a][0]-1
        j = commands[a][1]
        k = commands[a][2]-1
        sliced = array[i:j]
        sliced.sort()
        answer.append(sliced[k])
    return answer

print(solution([1, 5, 2, 6, 3, 7, 4],[[2, 5, 3], [4, 4, 1], [1, 7, 3]]))

"""
매우 간단한 문제.
"""