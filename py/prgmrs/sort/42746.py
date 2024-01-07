"""
가장 큰 수
https://school.programmers.co.kr/learn/courses/30/lessons/42746
"""
import functools

def comparator(a,b):
    x = int(''+str(a)+str(b))
    y = int(''+str(b)+str(a))
    return y - x
def solution(numbers):
    sort_ = sorted(numbers,key=functools.cmp_to_key(comparator))
    if sort_[0] == 0:
        return '0'
    return ''.join(map(str,sort_))
       

test = [[6, 10, 2], [3, 30, 34, 5, 9],[123,432,234]]

for i in test:
    print(solution(i))

"""
comparator는 연결된 값을 최대화하는 방식으로 숫자를 정렬하도록 한다.
"""

