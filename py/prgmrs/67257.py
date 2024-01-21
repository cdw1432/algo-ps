"""
[카카오 인턴] 수식 최대화
https://school.programmers.co.kr/learn/courses/30/lessons/67257?language=python3
"""
from itertools import permutations
def solution(expression):
    ans = -1
    exp,op = expToArr(expression)
    
    for i in range(len(op)):
        newArr = exp
        for j in range(len(op[i])):
            stack = []
            stack.append(newArr[0])
            for k in range(1, len(newArr)):
                if(stack[len(stack)-1] == op[i][j]):
                    co = stack.pop()
                    if(co == '*'):
                        stack.append(stack.pop() * newArr[k])
                    elif(co == '+'):
                        stack.append(stack.pop() + newArr[k])
                    elif(co == '-'):
                        stack.append(stack.pop() - newArr[k])
                else:
                    stack.append(newArr[k])
            newArr = stack
        if ans < abs(newArr[0]):
            ans = abs(newArr[0])         
    #print(exp,op)
    return ans


#반환: 문자열인 수식을 배열화, 연산자 경우의 수 배열화
def expToArr(str):
    li = 0
    arr = []
    op = set()
    for index, char in enumerate(str):
        if char == '*' or char == '-' or char == '+':
            op.add(char)
            arr.append(int(str[li:index]))
            arr.append(str[index])
            li = index + 1
        
    arr.append(int(str[li:]))
    return arr, list(permutations(op))
    

print(solution("100-200*300-500+20"))
print(solution("50*6-3*2"))
print(solution("3*2"))
