function findCircleNum(isConnected: number[][]): number {
    const provinces = isConnected.length;
    const visited = new Set();
    let circles = 0;
    const queue = new Array();
    for (let i = 0; i < provinces; i++) {
        if (!visited.has(i)) {
            queue.push(i);
            while (queue.length) {
                const j = queue.shift();
                visited.add(j);
                for (let k = 0; k < provinces; k++) {
                    if (isConnected[j][k] === 1 && !visited.has(k)) {
                        queue.push(k);
                    }
                }
            }
            circles++;
        }
    }
    return circles;
};

/**
 * 547. 省份数量
 */
test('leetcode_001', () => {
    console.log(findCircleNum([[1, 1, 0], [1, 1, 0], [0, 0, 1]]))
    console.log(findCircleNum([[1, 0, 0], [0, 1, 0], [0, 0, 1]]))
});