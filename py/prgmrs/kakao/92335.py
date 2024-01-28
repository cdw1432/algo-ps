"""
k진수에서 소수 개수 구하기
https://school.programmers.co.kr/learn/courses/30/lessons/92335
"""
def solution(n, k):
    answer = 0
    nums = conversion(n,k).split('0')
    #print(nums)
    for i in nums:
        if i and isPrime(int(i)):
            answer += 1
    return answer

def isPrime(n):
    #print(n)
    if n <= 1:
        return False
    for i in range(2, int(n**0.5)+1):
         if n % i == 0:
            return False
    return True
def conversion(n, b):
    e = n//b
    q = n%b
    if n == 0:
        return '0'
    elif e == 0:
        return str(q)
    else:
        return conversion(e, b) + str(q)
print(solution(437674,3))
print(solution(110011,10))