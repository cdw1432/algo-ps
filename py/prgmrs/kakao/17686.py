"""
[3차] 파일명 정렬
https://school.programmers.co.kr/learn/courses/30/lessons/17686
"""
from functools import cmp_to_key

def compare_files(a, b):
    def extract_head_number(s):
        head, number = "", ""
        i = 0
        while i < len(s) and not s[i].isdigit():
            head += s[i]
            i += 1
        while i < len(s) and s[i].isdigit():
            number += s[i]
            i += 1
        return head.lower(), int(number)

    a_head, a_number = extract_head_number(a)
    b_head, b_number = extract_head_number(b)

    if a_head < b_head:
        return -1
    elif a_head > b_head:
        return 1
    else:
        return a_number - b_number

letter_cmp_key = cmp_to_key(compare_files)
def solution(files):
    files.sort(key=letter_cmp_key)
    return files

print(solution( ["F-5 Freedom Fighter", "B-50 Superfortress", "A-10 Thunderbolt II", "F-14 Tomcat"]))
print(solution(["img12.png", "img10.png", "img02.png", "img1.png", "IMG01.GIF", "img2.JPG"]))
