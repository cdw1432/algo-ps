"""
양궁대회
https://school.programmers.co.kr/learn/courses/30/lessons/92342
"""
#어피치가 모든 점수에서 승리했을때 총 점수
def sum(arr):
    sum = 0
    for i in range(len(arr)):
        if(arr[i] >0):
            sum += (len(arr) - i) -1
    return sum

#경우의 수, 완전탐색
def search(r_arr, r_n, depth):
    if(depth == 11 or r_n == 0):
        return
    
    if(r_arr[depth] < r_n):
        search(r_arr, r_n - r_arr[depth], depth+1)
    
    search(r_arr, r_n, depth+1)

def solution(n, info):
    ans = []
    a_sum = sum(info)
    toWin = list(map(lambda x : x+1, info))
    
    search(toWin, n, 0)
    if len(ans) == 0:
        return [-1]
    return ans

print(solution(5, [2,1,1,1,0,0,0,0,0,0,0]))