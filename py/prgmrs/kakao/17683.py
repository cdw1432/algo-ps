"""
방금그곡
https://school.programmers.co.kr/learn/courses/30/lessons/17683
"""
#시간을 분으로 표현
def TimeConversion(start, end):
    s_time = start.split(':')
    e_time = end.split(':')

    start_time = int(s_time[0])*60 + int(s_time[1])
    end_time = int(e_time[0])*60 + int(e_time[1])

    return end_time - start_time

#멜로디를 배열화 함 #이 붙은 음표는 한 요소로 취급
# 그 후 네오가 기억하는 멜로디 배열이 음악 멜로디 배열의 subarr인지 확인
def Check(dest, src):
    i = 0
    arr = []
    while i < len(dest):
        str = dest[i]
        while i != len(dest)-1 and dest[i+1] == '#':
            str += dest[i+1]
            i += 1
        arr.append(str)
        i += 1
    
    j = 0
    arr2 = []
    while j < len(src):
        str = src[j]
        while j != len(src)-1 and src[j+1] == '#':
            str += src[j+1]
            j += 1
        arr2.append(str)
        j += 1
    return isSubArray(arr,arr2)

def isSubArray(arr1, arr2):
    for i in range(0, len(arr1)-len(arr2)+1):
        j = 0
        while j < len(arr2):
            if(arr1[i+j] != arr2[j]):
                break
            j += 1
        if(j == len(arr2)):
            return True
    return False

def solution(m, musicinfos):
    answer = ''
    maxPlayed = 0
    for i in range(len(musicinfos)):
        line = musicinfos[i].split(',')
        song_played_time = TimeConversion(line[0],line[1])
        song_name = line[2]
        song_melody = ''

        i = 0
        arr = []
        while i < len(line[3]):
            str = line[3][i]
            while i != len(line[3])-1 and line[3][i+1] == '#':
                str += line[3][i+1]
                i += 1
            arr.append(str)
            i += 1
        #재생시간에 맞추어 멜로디 만들기
        for j in range(song_played_time):
            song_melody += arr[j % (len(arr))]
        
        if Check(song_melody, m) and maxPlayed < song_played_time:
            answer = song_name
            maxPlayed = song_played_time
    if answer == '':
        return "(None)"
    return answer

print(solution("ABC",["12:00,12:14,HELLO,C#", "13:00,13:05,WORLD,ABCDEF"]	))

#스파게티코드 ㅜ