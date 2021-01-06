function countAndSay(n: number): string {
    let sub: number = 0, resultTemp: string = "", result: string = "1";
    while (n > 1) {
        sub = 0;
        resultTemp = "";
        for (let i = 1; i < result.length; i++) {
            if (result[sub] != result[i]) {
                resultTemp = resultTemp + (i - sub) + result[sub]
                sub = i
            }
        }
        resultTemp = resultTemp + (result.length - sub) + result[sub]
        result = resultTemp;
        n--
    }
    return result;
};

/**
 * 38. 外观数列
 */
test('leetcode_001', () => {
    console.log(countAndSay(1))
    console.log(countAndSay(4))
    console.log(countAndSay(5))
});