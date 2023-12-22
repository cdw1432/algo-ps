# bascis
print('hello world')

변수입니다 = 0.3 + 0.6
print(변수입니다)
print(round(변수입니다))
print(round(변수입니다,1))

a = 5
b = 3
print( a % b)
print ( a / b)
print ( a // b)
print ( a ** b)

a = [0] * 10
print(a)
b = [1,2,3,4,5,6,7,8,9,10]
print(b[3:6])

mulOf3 = [i for i in range(10) if i%3 == 0 and i != 0]
print(mulOf3)

twoDArr = [[0] * 3 for _ in range(3)] # 3(m) X 3(n)
print(twoDArr)
'''
append()
sort()
reverse
insert
count
remove
'''

data = dict()
data['apple'] = 3
data['banana'] = 5

print(data)

set1 = set([1,2,3,3,4,5,5])
set2 = set([3,3,4,4,5,5,6,6,7,7])

print(set1)
print(set2)
print( set1 | set2) #합집합
print( set1 & set2) #교집합
print( set1 - set2) #차집합

#람다식
print((lambda a,b: a+b)(10,5))

#input
n = int(input())
inputArr = list(map(int, input().split()))
print(inputArr)

# 주요 라이브러리
'''
내장 함수
itertools
heapq
bisect
collections
math
'''
