"""
크레인 인형뽑기 게임
https://school.programmers.co.kr/learn/courses/30/lessons/64061
"""
def solution(board, moves):
    answer = 0
    basket = []

    for move in moves:
        i = 0
        while i < len(board) and board[i][move-1] == 0:
            if i == len(board)-1:
                break
            i += 1

        #print(board[i][move-1])
        if board[i][move-1] != 0:
            if len(basket) != 0 and basket[len(basket)-1] == board[i][move-1]:
                basket.pop()
                answer += 2
            else:
                basket.append(board[i][move-1])
            board[i][move-1] = 0
    return answer

print(solution([[0,0,0,0,0],[0,0,1,0,3],[0,2,5,0,1],[4,2,4,4,2],[3,5,1,3,1]],[1,5,3,5,1,2,1,4]))