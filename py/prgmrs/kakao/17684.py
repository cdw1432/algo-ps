"""
[3차] 압축
https://school.programmers.co.kr/learn/courses/30/lessons/17684
"""
def solution(msg):
    ans = []
    cnt = 27
    words = {
        "A":1, "B":2, "C":3, "D":4, "E":5,
        "F":6, "G":7, "H":8, "I":9, "J":10,
        "K":11,"L":12,"M":13,"N":14,"O":15,
        "P":16,"Q":17,"R":18,"S":19,"T":20,
        "U":21,"V":22,"W":23,"X":24,"Y":25,
        "Z":26,
    }
    i = 0
    while(i < len(msg)):
        str = msg[i]
        while i != len(msg)-1 and str in words:
            next = msg[i+1]
            if str+next in words:
                str = str + next
                i += 1
            else:
                words[str+next] = cnt
                cnt += 1
                break
        
        if str in words:
            ans.append(words[str])
        i += 1
        
    #print(words)
    return ans

print(solution("KAKAO"))
print(solution("TOBEORNOTTOBEORTOBEORNOT"))