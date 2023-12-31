class Heap {
    constructor() {
        this.heap = [];
    }
    isEmpty() {
        return this.heap.length === 0;
    }
    push() {

    }
    pop() {

    }
}
const tests = [
    ["I 16", "I -5643", "D -1", "D 1", "D 1", "I 123", "D -1"],
    ["I -45", "I 653", "D 1", "I -642", "I 45", "I 97", "D 1", "D -1", "I 333"]	
]

for(let t of tests) {
    console.log(solution(t));
}

function solution(operations) {
    //var heap = new Heap();
    let arr = [];
    for(let i = 0; i < operations.length; i++) {
        let op = operations[i].split(' ');
        if(op[0] === 'D') {
            if(arr.length > 0) {
                arr.sort((a,b) => a-b);
                if(op[1] == '1') {
                    arr.pop();
                } else {
                    arr.shift();
                }
            }
        } else {
            arr.unshift(parseInt(op[1]))
        }
    }
    if(arr.length > 0) {
        arr.sort((a,b) => a-b);
        return [arr[arr.length-1],arr[0]]
    } else {
        return [0,0]
    }
}
/*
힙 안쓰고 기본 sort로 풀음.
*/