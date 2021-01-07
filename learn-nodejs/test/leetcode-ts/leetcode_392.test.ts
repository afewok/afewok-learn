function isSubsequence(s: string, t: string): boolean {
    if (s.length == 0) {
        return true
    }
    let sub: number = 0
    for (let i = 0; i < t.length; i++) {
        if (s[sub] == t[i]) {
            sub++
            if (sub == s.length) {
                return true
            }
        }
    }
    return false
};

/**
 * 392. 判断子序列
 */
test('leetcode_392', () => {
    console.log(isSubsequence("abc", "ahbgdc"))
    console.log(isSubsequence("axc", "ahbgdc"))
});