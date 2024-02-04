"""
[3차] n진수 게임
https://school.programmers.co.kr/learn/courses/30/lessons/17687
"""
def conversion(number, base):
    base_string = "0123456789ABCDEF"
    result = ""
    while number:
        result += base_string[number % base]
        number //= base
    return result[::-1] or "0"

def solution(n, t, m, p):
    answer = ''
    arr = []
    for i in range(0, m*t):
            num_str = conversion(i, n)
            num_arr = list(map(str,num_str))
            for j in num_arr:
                arr.append(j)

    i = p-1
    while len(answer) < t:
        answer += arr[i]
        i += m
    return answer

print(solution(2,4,2,1))
print(solution(16,16,2,1))