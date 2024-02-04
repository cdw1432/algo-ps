"""
두 큐 합 같게 만들기
https://school.programmers.co.kr/learn/courses/30/lessons/118667
"""
def solution(queue1, queue2):
    whole = queue2 + queue1 + queue2 + queue1
    l = len(queue1+queue2)
    ans = 0
    trigger = False
    for i in range(l):
        for j in range(i+1,l):
            q1 = whole[i:j]
            q2 = whole[j:i+l]
            if sum(q1) == sum(q2):
                trigger = True
                print(q1,q2)
    if trigger == False:
        return -1
    return ans

print(solution([3, 2, 7, 2]	,[4, 6, 5, 1]))
print(solution([1, 2, 1, 2]		,[1, 10, 1, 2]	))
print(solution([1, 1]		,[1, 5]	))