"""
MooTube
https://www.acmicpc.net/problem/15591
"""
N, Q = map(int, input().split())

graph  = [[] for _ in range(N+1)]  
for _ in range(N-1):
    p, q, r = map(int, input().split())
    graph[p].append((q, r))  
    graph[q].append((p, r))

def bfs(k,v):
    order = []
    visited = [0] * (N+1)
    queue = []
    queue.append(v)
    while len(queue) > 0:
        c = queue[0]
        queue = queue[1:]
        
        for i,val in graph[c]:
            if visited[i] == 0 and i != v and val >= k:
                visited[i] = 1
                order.append(c)
                queue.append(i)
    
    return len(order)

ans = []
for _ in range(Q):
    k,v = map(int, input().split())
    ans.append(bfs(k,v))

for count in ans:
    print(count)

