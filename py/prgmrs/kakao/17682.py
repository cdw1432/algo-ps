"""
[1차] 다트 게임
https://school.programmers.co.kr/learn/courses/30/lessons/17682
"""
def solution(dartResult):
    i = 0
    stack = []
    num = ''
    while i < len(dartResult):
        if dartResult[i] == 'S':
            stack.append(int(num))
            num = ''
            i += 1
            continue
        elif dartResult[i] == 'D':
            stack.append(int(num)**2)
            i += 1
            num = ''
            continue
        elif dartResult[i] == 'T':
            stack.append(int(num)**3)
            i += 1
            num = ''
            continue
        if dartResult[i] == '*':
            j = 0
            while j < len(stack):
                stack[len(stack)-1-j] *= 2
                j += 1
                if j == 2:
                    break
            i += 1
            continue
        elif dartResult[i] == '#':
            stack[len(stack)-1] *= -1
            i += 1
            continue
        
        num += dartResult[i]
        i += 1
    
    return sum(stack)

print(solution("1S2D*3T"))
print(solution("1D2S#10S"))