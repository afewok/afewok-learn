function isStrobogrammatic(num: string): boolean {
    let left = 0, right: number = num.length - 1
    while (left <= right) {
        if ((num[left] == "0" && num[right] == "0") ||
            (num[left] == "1" && num[right] == "1") ||
            (num[left] == "6" && num[right] == "9") ||
            (num[left] == "9" && num[right] == "6") ||
            (num[left] == "8" && num[right] == "8")) {
            left++
            right--
            continue
        }
        return false
    }
    return true
};

/**
 * 246. 中心对称数
 */
test('leetcode_246', () => {
    1234567890
    console.log(isStrobogrammatic("010"))
    console.log(isStrobogrammatic("69"))
    console.log(isStrobogrammatic("88"))
    console.log(isStrobogrammatic("962"))
    console.log(isStrobogrammatic("1"))
    console.log(isStrobogrammatic("2"))
});