function countAndSay(n: number): string {
    let sub: number = 0, resultTemp: string = "", result: string = "1";
    while (n > 1) {
        for (let i = 1; i < result.length; i++) {
            if(result[temp]!=result[i]){

            }

        }
        result = resultTemp;
        resultTemp = "";
        sub=0;
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
});