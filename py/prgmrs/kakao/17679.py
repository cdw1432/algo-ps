"""
[1차] 프렌즈4블록
https://school.programmers.co.kr/learn/courses/30/lessons/17679
"""
def solution(m, n, board):
    answer = 0

    while True:
        remove = set()
        for i in range(m-1):
            for j in range(n-1):
                if board[i][j] == '$':
                    continue

                if board[i][j] == board[i][j+1] and board[i][j] == board[i+1][j] and board[i][j] == board[i+1][j+1]:
                    remove.add((i,j))
                    remove.add((i, j+1))
                    remove.add((i+1, j))
                    remove.add((i+1, j+1))
        if len(remove) == 0:
            break

        answer += len(remove)
        for j in range(n):
            stack = []
            #print(board)
            for i in range(m-1,-1, -1):
                if not ((i,j) in remove):
                    stack.append(board[i][j])
            while len(stack) < m:
                stack.append("$")
            
            #print(stack)
            for k in range(m):
                board[k] = board[k][:j] + stack[len(stack)-1-k] + board[k][j+1:]
            

        #print(board)
    return answer

print(solution(4,5,["CCBDE", "AAADE", "AAABF", "CCBBF"]))
print(solution(6,6,["TTTANT", "RRFACC", "RRRFCC", "TRRRAA", "TTMMMF", "TMMTTJ"]))