/*
프로세스
https://school.programmers.co.kr/learn/courses/30/lessons/42587
*/
const tests = 2
const tests_arg1 = [[2, 1, 3, 2],[1, 1, 9, 1, 1, 1]]
const tests_arg2 = [2,0]

for(let t = 0; t < tests; t++) {
    console.log(solution(tests_arg1[t],tests_arg2[t]));
}

function solution(priorities, location) {
    var answer = 0;
    while (priorities.length > 0) {
        let c = priorities.shift();

        if(c < Math.max(...priorities)) {
            priorities.push(c);
            if(location == 0)
                location = priorities.length-1;
            else
                location--;
        }
        else {
            answer++;
            if(location <= 0) {
                break;
            }
            location--;
        }

    }
    return answer;
}
/*
규칙에 따라 프로세스를 관리하여 실행 되는 큐를 구현하는 것은
쉬웠지만 특정 인덱스가 몇번째로 실행되는지 추적하기가 까다롭다.

첫번째 생각은 location을 반복마다 업데이트 시키는 것.
하지만 잘 안됐다.

두번째는 배열의 요소가 두개의 정보를 담게 하는것 [val, idx]. 이러면 요소가 실행될때
이 요소가 우리가 원하는 요소인지 알수있다. 하지만 저장공간이 늘어나기 때문에 적합하지 않다고 판단.

첫번째로 돌아와 location을 업데이트해주는 경우를 계산.
location의 요소가 차례가 아닐때.
    location이 0일때, 배열 인덱스로 업데이트.
              0이 아닐때,  1빼기;

location의 요소가 차례일때. (location == 0)
    없어진 요소들의 숫자 반환. 
*/