function longestCommonPrefix(strs: string[]): string {
    if (strs.length == 0) {
        return "";
    } else if (strs.length == 1) {
        return strs[0];
    }

    let sub = 0, temp = null, result = "";
    while (true) {
        temp = null;
        for (let i = 0; i < strs.length; i++) {
            if (strs[i].length <= sub) {
                temp = null;
                break;
            }
            if (temp == null) {
                temp = strs[i][sub];
            } else if (temp != strs[i][sub]) {
                temp = null;
                break;
            }
        }
        if (temp == null) {
            break;
        }
        result = result + temp;
        sub++;
    }
    return result;
};

/**
 * 14. 最长公共前缀
 */
test('leetcode_014', () => {
    console.log(longestCommonPrefix(["flower", "flow", "flight"]))
    console.log(longestCommonPrefix(["dog", "racecar", "car"]))
    console.log(longestCommonPrefix(["ab", "a"]))
});