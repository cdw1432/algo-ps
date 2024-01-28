"""
[1차] 캐시
https://school.programmers.co.kr/learn/courses/30/lessons/17680
"""
def cacheHit(cache, ele):
    for i in range(len(cache)):
        if cache[i] == ele:
            return i
    return -1
def solution(cacheSize, cities):
    cache = []
    ans = 0
    for i in range(len(cities)):
        if cacheHit(cache, cities[i].lower()) >= 0:
            ans += 1
            x = cache[cacheHit(cache, cities[i].lower())]
            cache.pop(cacheHit(cache, cities[i].lower()))
            cache.append(x)
        else:
            ans += 5
            if cacheSize == 0:
                continue
            if len(cache) == cacheSize:
                cache = cache[1:]
                cache.append(cities[i].lower())
            else:
                cache.append(cities[i].lower())

    return ans

print(solution(0, ["Jeju", "Pangyo", "Seoul", "NewYork", "LA"]	))