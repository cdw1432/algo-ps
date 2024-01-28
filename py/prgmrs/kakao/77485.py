"""
행렬 테두리 회전하기
https://school.programmers.co.kr/learn/courses/30/lessons/77485
"""
def solution(rows, columns, queries):
    ans = []
    #행렬 만들기
    graph = [[0 for x in range(columns)] for y in range(rows)]
    for i in range(rows):
        for j in range(columns):
            graph[i][j] = ((i) * columns + j+1)

    #회전
    for v in queries:
        graph,min = rotate(graph,[v[0]-1,v[1]-1], [v[2]-1,v[3]-1])
        ans.append(min)
    

    return ans

#테두리를 일자로 핀 다음 오른쪽으로 한칸씩 옮기고 Matrix에 반영하는 방법
def rotate(matrix, q1, q2):        
        di = [0,1,0,-1]
        dj = [1,0,-1,0]

        si,sj = q1[0],q1[1]
        c = 0
        arr =[]
        while(True):
            arr.append(matrix[si][sj])
            si = si + di[c]
            sj = sj + dj[c]
            if (si == q1[0] and sj == q2[1]) or (si == q2[0] and sj == q2[1]) or (si == q2[0] and sj == q1[1]):
                 c += 1
            if si == q1[0] and sj == q1[1]:
                break
        
        #가장 작은 숫자 찾기
        small = min(arr)
        #맨 뒤 요소를 맨 앞으로 (오른쪽으로 한칸 옮기는 효과)
        arr = arr[len(arr)-1:] + arr[:len(arr)-1]

        #matrix에 반영
        si,sj = q1[0],q1[1]
        c = 0
        i = 0
        while(True):
            matrix[si][sj] = arr[i]
            i += 1
            si = si + di[c]
            sj = sj + dj[c]
            if (si == q1[0] and sj == q2[1]) or (si == q2[0] and sj == q2[1]) or (si == q2[0] and sj == q1[1]):
                 c += 1
            if si == q1[0] and sj == q1[1]:
                break
        
        # for i in matrix:
        #     print(i) 
        return matrix, small

print(solution(6,6,[[2,2,5,4]]))
print(solution(100, 97,[[1,1,100,97]]))