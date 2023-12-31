const tests = [
    [[1, 2, 3, 9, 10, 12],7]
]
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
        const lastElement = this.heap.pop();

        if (this.heap.length > 0) {
            this.heap[0] = lastElement;
            this.heapifyDown();
        }

        return root;
    }

    isEmpty() {
        return this.heap.length === 0;
    }

    heapifyUp() {
        let currentIndex = this.heap.length - 1;

        while (currentIndex > 0) {
            const parentIndex = Math.floor((currentIndex - 1) / 2);

            if (this.heap[currentIndex] < this.heap[parentIndex]) {
                [this.heap[currentIndex], this.heap[parentIndex]] = [this.heap[parentIndex], this.heap[currentIndex]];
                currentIndex = parentIndex;
            } else {
                break;
            }
        }
    }

    heapifyDown() {
        let currentIndex = 0;

        while (true) {
            const leftChildIndex = 2 * currentIndex + 1;
            const rightChildIndex = 2 * currentIndex + 2;
            let nextIndex = currentIndex;

            if (leftChildIndex < this.heap.length && this.heap[leftChildIndex] < this.heap[nextIndex]) {
                nextIndex = leftChildIndex;
            }

            if (rightChildIndex < this.heap.length && this.heap[rightChildIndex] < this.heap[nextIndex]) {
                nextIndex = rightChildIndex;
            }

            if (nextIndex !== currentIndex) {
                [this.heap[currentIndex], this.heap[nextIndex]] = [this.heap[nextIndex], this.heap[currentIndex]];
                currentIndex = nextIndex;
            } else {
                break;
            }
        }
    }
}

for(let t of tests) {
    console.log(solution(t[0],t[1]))
}
function solution(scoville, K) {
    const minHeap = new MinHeap();
    let ans = 0;
    for(let i of scoville) 
        minHeap.push(i);
    console.log(minHeap)
    while(minHeap.heap[0] < K) {
        minHeap.push(minHeap.pop() + 2*minHeap.pop());
        ans++;

        if(minHeap.heap.length == 1 && minHeap.heap[0] < K)
            return -1;
    }    
    return ans;
}
/*
파이썬과 달리 자바스크립트에는 minheap 라이브러리가 없기 때문에 구현할줄 알아야한다.
힙 자료구조를 잘 알아야한다.
*/