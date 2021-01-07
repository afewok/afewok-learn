function rotate(nums: number[], k: number): number[] {
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

/**
 * 189. 旋转数组
 */
test('leetcode_189', () => {
    console.log(rotate([1, 2, 3, 4, 5, 6, 7], 3))
    console.log(rotate([1, 2, 3, 4, 5, 6, 7], 15))
});