"""
주차 요금 계산
https://school.programmers.co.kr/learn/courses/30/lessons/92341
"""

import math
# 기본요금 + (ceil(누적 주차시간 - 기본 시간)/단위시간) * 단위요금
# 
# 누적주차시간이 기본시간 이하라면 기본요금만.
def solution(fees, records):
    ans = []
    db = dict()
    for event in records:
        arr = event.split(' ')
        inNout = arr[2]
        car_num = arr[1]
        time = conversion(arr[0][:2],arr[0][3:])

        if not (car_num in db):
            if(inNout == 'IN'):
                db[car_num] = [inNout, -1*time, int(car_num)]
            else:
                db[car_num] = [inNout, time, int(car_num)]
        else:
            if(inNout == 'IN'):
                db[car_num][0] = 'IN'
                db[car_num][1] -= time
            else:
                db[car_num][0] = 'OUT'
                db[car_num][1] += time
    
    sorted_dict = dict(sorted(db.items(), key=lambda item: item[1][2]))
    for i in sorted_dict:
        status = db[i][0]
        time = db[i][1]

        if(status == 'IN'):
            time += conversion("23","59")
        if time <= fees[0]:
            ans.append(fees[1])
        else:
            ans.append(fees[1] + math.ceil((time-fees[0])/fees[2])*fees[3])
    return ans

def conversion(time_h, time_m):
    return  int(time_h)*60 + int(time_m)

print(solution([180, 5000, 10, 600],["05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT"]))