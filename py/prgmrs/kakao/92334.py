"""
신고 결과 받기
https://school.programmers.co.kr/learn/courses/30/lessons/92334
"""
def solution(id_list, report, k):
    ans = []
    users = dict() #각 유저가 누구(들) 신고했는지
    users_ = dict() #각 유저의 신고당한 횟수
    for id in id_list:
        users[id] = dict()
        users_[id] = 0

    #데이터 베이스
    for line in report:
        r = line.split(' ')
        if not (r[1] in users[r[0]]):
            users[r[0]][r[1]] = True
            users_[r[1]] += 1
    #각 유저가 신고한 유저(들)이 정지 당했는지 확인 후 메일 발송
    for id in id_list:
        cnt = 0
        for reported in users[id]:
            if users_[reported] >= k:
                cnt += 1
        ans.append(cnt)
    # print(users)
    # print(users_)
    return ans

print(solution(["muzi", "frodo", "apeach", "neo"],["muzi frodo","apeach frodo","frodo neo","muzi neo","apeach muzi"],2))

#다른 사람 풀이중, 좋은 답안
"""
def solution(id_list, report, k):
    answer = [0] * len(id_list)    
    reports = {x : 0 for x in id_list}

    for r in set(report):
        reports[r.split()[1]] += 1

    for r in set(report):
        if reports[r.split()[1]] >= k:
            answer[id_list.index(r.split()[0])] += 1

    return answer
"""