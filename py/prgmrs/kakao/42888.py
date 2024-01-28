"""
오픈채팅방
https://school.programmers.co.kr/learn/courses/30/lessons/42888
"""
def solution(record):
    answer = []
    user = dict()
    for command in record:
        arr = command.split(' ')
        if(arr[0] == 'Enter' or arr[0] == 'Change'):
            user[arr[1]] = arr[2]

    for command in record:
        arr = arr = command.split(' ')
        if(arr[0] == 'Enter'):
            answer.append(user[arr[1]]+"님이 들어왔습니다.")
        if(arr[0] == 'Leave'):
            answer.append(user[arr[1]]+"님이 나갔습니다.")
    #print(user)
    return answer

print(solution(["Enter uid1234 Muzi", "Enter uid4567 Prodo","Leave uid1234","Enter uid1234 Prodo","Change uid4567 Ryan"]))