function isMatch(s: string, p: string): boolean {

};

/**
 * 10. 正则表达式匹配
 */
test('leetcode_010', () => {
    console.log(isMatch( "aa", "a"))
    console.log(isMatch( "aa", "a*"))
    console.log(isMatch( "ab", ".*"))
    console.log(isMatch( "aab", "c*a*b"))
    console.log(isMatch( "mississippi", "mis*is*p*."))
});