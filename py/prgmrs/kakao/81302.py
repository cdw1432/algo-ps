"""
거리두기 확인하기
https://school.programmers.co.kr/learn/courses/30/lessons/81302
"""
def solution(places):
    answer = [1] * len(places)
    for i in range(len(places)):
        for j in range(len(places[i])):
            for k in range(len(places[i][j])):
                if places[i][j][k] == 'P':
                    if search(places[i], [j,k]) == 0:
                        answer[i] = 0
                        break
    return answer


def move(grid, init, n):
    x,y = init
    dx = [-1,1,0,0]
    dy = [0,0,-1,1]
    ret = 1
    for i in range(0,2):
        nx = x + dx[n]
        ny = y + dy[n]
        if nx < 0 or ny < 0 or nx > 4 or ny > 4:
            break
        
        #print(nx,ny, grid[nx][ny])
        if grid[nx][ny] == 'X':
            break
        elif grid[nx][ny] == 'P':
            ret = 0
        x = nx
        y = ny
    
    return ret

 

def search(grid, init):
    top = move(grid, init, 0)
    down = move(grid, init, 1)
    left = move(grid, init, 2)
    right = move(grid, init, 3)

    x,y = init
    dg = [[x-1,y-1], [x-1,y+1],[x+1,y+1],[x+1,y-1]]
    dig = 1
    for c in dg:
        destx, desty = c
        if destx < 0 or desty < 0 or destx > 4 or desty > 4:
            continue
        if grid[destx][desty] == 'P':
            if grid[destx][y] != 'X' or grid[x][desty] != 'X':
                dig = 0

    print(top,down,left,right,dig)
    if top == 0 or down == 0 or left == 0 or right == 0 or dig == 0:
        return 0
    return 1


print(solution([["POOOP", "OXXOX", "OPXPX", "OOXOX", "POXXP"], ["POOPX", "OXPXP", "PXXXO", "OXXXO", "OOOPP"], ["PXOPX", "OXOXP", "OXPOX", "OXXOP", "PXPOX"], ["OOOXX", "XOOOX", "OOOXX", "OXOOX", "OOOOO"], ["PXPXP", "XPXPX", "PXPXP", "XPXPX", "PXPXP"]]))
print(solution( [["POOOO", "XPOOO", "OOOOO", "OOOOO", "OOOOO"]]))