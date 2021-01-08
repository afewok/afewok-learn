function rotate1(nums: number[], k: number): number[] {
    let length: number = nums.length, num: number = 0
    if (length < k) {
        k = k % length
    }

    for (let i = 0; i < k; i++) {
        num = nums[length - 1]
        for (let j = length - 1; j > 0; j--) {
            nums[j] = nums[j - 1]
        }
        nums[0] = num
    }
    return nums
};

function rotate(nums: number[], k: number): number[] {
    let length: number = nums.length
    if (length < k) {
        k = k % length
    }
    let sub: number = 0, num: number = nums[0], count: number = 0
    let p: number = sub, q: number = p - k < 0 ? length + p - k : p - k
    while (count < length) {
        count++
        if (q == sub) {
            nums[p] = num
            sub = q + 1 == length ? 0 : q + 1
            num = nums[sub]
            p = sub
        } else {
            nums[p] = nums[q]
            p = q
        }
        q = p - k
        q = q < 0 ? length + q : q
    }
    return nums
}

/**
 * 189. 旋转数组
 */
test('leetcode_189', () => {
    console.log(rotate([1], 0))
    console.log(rotate([1, 2, 3, 4, 5, 6, 7], 3))
    console.log(rotate([-1, -100, 3, 99], 2))
    console.log(rotate([1, 2, 3, 4, 5, 6, 7], 15))
});