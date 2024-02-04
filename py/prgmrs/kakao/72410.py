"""
신규 아이디 추천
https://school.programmers.co.kr/learn/courses/30/lessons/72410
"""
def solution(new_id):
    answer = ''
    #1
    new_id = new_id.lower()
    #2
    second = ''
    for i in new_id:
        if  (i >= '0' and i <= '9') or  (i >= 'a' and i <= 'z') or  (i == '-') or  (i == '_') or  (i == '.'): 
            second += i
    #3
    i = 1
    third = second[0]
    while i < len(second):
        if third[len(third)-1] == '.' and second[i] == '.':
            i += 1
            continue
        third += second[i]
        i += 1
    #4
   
    if third[0] == '.':
        third = third[1:]
    if len(third) != 0 and third[len(third)-1] == '.':
        third = third[:len(third)-1]
    #5
    if len(third) ==  0:
        third = 'a'
    #6
    if len(third) >= 16:
        third = third[:15]
    if len(third) != 0 and third[len(third)-1] == '.':
        third = third[:len(third)-1]

    #7
    if len(third) <= 2:
        while len(third) != 3:
            third += third[len(third)-1]

    return third

print(solution("...!@BaT#*..y.abcdefghijklm"))
print(solution("=.="))
print(solution("123_.def"		))
print(solution("abcdefghijklmn.p"	))