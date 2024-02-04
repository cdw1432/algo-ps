"""
메뉴 리뉴얼
https://school.programmers.co.kr/learn/courses/30/lessons/72411
"""
def solution(orders, course):
    arr = []
    for i in range(len(orders)):
        menu = set(list(orders[i]))
        for j in range(i+1, len(orders)):
            tmp = set(list(orders[j]))
            if len(menu & tmp) >= 2:
                arr.append(sorted(menu & tmp))
    
    print(arr)
    ans = []
    for c in course:
        cnter = dict()
        for e in arr:
            if len(e) == c:
                if "".join(e) in cnter:
                    cnter["".join(e)] += 1
                else:
                    cnter["".join(e)] = 1
        print(cnter)
        max_num = max(cnter.values(),default=0)
        for k in cnter:
            if cnter[k] == max_num:
                ans.append("".join(k))
    return sorted(ans)

print(solution(["ABCFG", "AC", "CDE", "ACDE", "BCFG", "ACDEH"],[2,3,4]))
print(solution(["ABCDE", "AB", "CD", "ADE", "XYZ", "XYZ", "ACD"],[2,3,5]))