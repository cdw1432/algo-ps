/* 
폰켓몬
https://school.programmers.co.kr/learn/courses/30/lessons/1845
*/
tests = [[3,1,2,3],
        [3,3,3,2,2,4],
        [3,3,3,2,2,2]]

for(let i = 0; i < tests.length; i++) {
    console.log(solution(tests[i]));
}
function solution(nums) {
    let t = nums.length / 2;
    let map = {};
    for(let i = 0; i < nums.length; i++) {
        if(map[nums[i]])
            map[nums[i]]++;
        else
            map[nums[i]] = 1;
    }
    if(Object.keys(map).length > t) {
        return t;
    }
    return Object.keys(map).length;
}
// 위 방법 말고 Set을 이용하면 코드가 간결해짐.