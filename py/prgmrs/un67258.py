"""
미풀
[카카오 인턴] 보석 쇼핑
https://school.programmers.co.kr/learn/courses/30/lessons/67258
"""
def solution(gems):
    answer = [0, len(gems)-1]
    start, end = 0, 0

    #보석 종류 수
    kind = len(set(gems))

    num = {gems[0] : True}
    while start < len(gems) and end < len(gems):
        if len(num) == kind:
            if end - start < answer[1] - answer[0]:
                answer = [start,end]
            else:
        else:
            end += 1
            if end == len(gems):
                break
            if gems[]

    print(kind)
    return answer
print(solution(["DIA", "RUBY", "RUBY", "DIA", "DIA", "EMERALD", "SAPPHIRE", "DIA"]))
print(solution(["D","D","R","R","R","D","D"]))