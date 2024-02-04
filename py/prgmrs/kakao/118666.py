"""
성격 유형 검사하기
https://school.programmers.co.kr/learn/courses/30/lessons/118666
"""
def solution(survey, choices):
    answer = ''
    index = {
        "T":0,
        "R":0,
        "C":0,
        "F":0,
        "J":0,
        "M":0,
        "A":0,
        "N":0,
    }
    for i in range(len(survey)):
        disagree = survey[i][0]
        agree = survey[i][1]
        if choices[i] > 4:
            index[agree] += choices[i] - 4
        elif choices[i] < 4:
            index[disagree] += 4 - choices[i]

    answer += 'T' if index['T'] > index['R'] else 'R'
    answer += 'F' if index['F'] > index['C'] else 'C'
    answer += 'M' if index['M'] > index['J'] else 'J'
    answer += 'N' if index['N'] > index['A'] else 'A'
    return answer

print(solution(["AN", "CF", "MJ", "RT", "NA"],[5, 3, 2, 7, 5]))