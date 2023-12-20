/* 
의상
https://school.programmers.co.kr/learn/courses/30/lessons/42578
*/

const tests = [
        [["yellow_hat", "headgear"], ["blue_sunglasses", "eyewear"], ["green_turban", "headgear"]],
        [["crow_mask", "face"], ["blue_sunglasses", "face"], ["smoky_makeup", "face"]],
        [["glasses", "head"], ["hat", "head"], ["t-shirt", "mid"], ["hoodie", "mid"],["shorts","bottom"],["long","bottom"]]
    ]

for(let t of tests)
    console.log(solution(t));
function solution(clothes) {
    let answer = 1;
    let map = new Map();
    for(let item of clothes) {
        map.set(item[1],(map.get(item[1]) || 0) + 1);
    }
    map.forEach((v,k) => {
        answer *= (v+1);
    })

    return answer-1;
}
/* 
종류, 의상

머리: 안경, 모자 (2)    => 입었을때(2) + 안입었을때(1) = 3
상의: 티셔츠, 후드티 (2) => 입었을때(2) + 안입었을때(1) = 3
하의: 반바지, 긴바지 (2) => 입었을때(2) + 안입었을때(1) = 3

경우의 수
3 x 3 x 3 = 27

의상 하나도 안입었을때 제외
27 -1

26
*/