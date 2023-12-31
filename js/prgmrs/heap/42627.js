class MinHeap {
    constructor() {
        this.heap = [];
    }

    push(value) {
        this.heap.push(value);
        this.heapifyUp();
    }

    pop() {
        if (this.isEmpty()) {
            return null;
        }
        const root = this.heap[0];
        const lastE = this.heap.pop();

        if (this.heap.length > 0) {
            this.heap[0] = lastE;
            this.heapifyDown();
        }
        return root;
    }

    isEmpty() {
        return this.heap.length === 0;
    }

    heapifyUp() {
        let cI = this.heap.length - 1;

        while (cI > 0) {
            const pI = Math.floor((cI - 1) / 2);

            if (this.heap[cI][1] < this.heap[pI][1]) {
                [this.heap[cI], this.heap[pI]] = [this.heap[pI], this.heap[cI]];
                cI = pI;
            } else {
                break;
            }
        }
    }

    heapifyDown() {
        let cI = 0;

        while (true) {
            const lCI = 2 * cI + 1;
            const rCI = 2 * cI + 2;
            let nI = cI;

            if (lCI < this.heap.length && this.heap[lCI][1] < this.heap[nI][1])
                nI = lCI;

            if (rCI < this.heap.length && this.heap[rCI][1] < this.heap[nI][1])
                nI = rCI;

            if (nI !== cI) {
                [this.heap[cI], this.heap[nI]] = [this.heap[nI], this.heap[cI]];
                cI = nI;
            } else {
                break;
            }
        }
    }
}

const tests = [[[0, 3], [1, 9], [2, 6]]];

for (let t of tests) {
    console.log(solution(t));
}

function solution(jobs) {
    let answer = 0;
    let i = 0;
    let timer = 0;
    let start = -1;
    let minheap = new MinHeap();

    while (i < jobs.length) {
        for ( j of jobs) {
            if (start < j[0] && j[0] <= timer) {
                minheap.push(j);
            }
        }
        if (!minheap.isEmpty()) {
            let c = minheap.pop();
            start = timer;
            timer += c[1];
            answer += (timer - c[0]);
            i++;
        } else {
            timer++;
        }
    }
    return Math.floor(answer/jobs.length)
}
/*
최소 힙을 사용함.
문제 풀이에 도움을 준 자료.
https://velog.io/@younge/Python-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%A8%B8%EC%8A%A4-%EB%94%94%EC%8A%A4%ED%81%AC-%EC%BB%A8%ED%8A%B8%EB%A1%A4%EB%9F%AC-%ED%9E%99
*/