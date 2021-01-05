function largeGroupPositions(s: string): number[][] {
    let arr = new Array(), sub = 0, length = s.length;
    for (let i = 0; i < length; i++) {
        if (s[sub] != s[i]) {
            if (i - sub >= 3) {
                arr.push([sub, i - 1]);
            }
            sub = i;
        }
    }
    if (length - sub >= 3) {
        arr.push([sub, length - 1]);
    }
    return arr;
};

/**
 * 830. 较大分组的位置
 */
test('leetcode_830', () => {
    console.log(largeGroupPositions("abbxxxxzzy"))
    console.log(largeGroupPositions("abc"))
    console.log(largeGroupPositions("abcdddeeeeaabbbcd"))
    console.log(largeGroupPositions("aba"))
    console.log(largeGroupPositions("aaa"))
});