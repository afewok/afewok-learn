function addDigits(num: number): number {
    let temp: string = ""
    while (num >= 10) {
        temp = String(num)
        num = 0
        for (let i = 0; i < temp.length; i++) {
            num = num + Number(temp[i])
        }
    }
    return num
};

/**
 * 258. 各位相加
 */
test('leetcode_258', () => {
    console.log(addDigits(10))
    console.log(addDigits(38))
});