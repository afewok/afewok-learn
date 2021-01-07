function containsNearbyDuplicate1(nums: number[], k: number): boolean {
    for (let i = 0; i < nums.length; i++) {
        for (let j = i + k < nums.length ? i + k : nums.length - 1; j > i; j--) {
            if (nums[i] == nums[j]) {
                return true
            }
        }
    }
    return false
};

function containsNearbyDuplicate(nums: number[], k: number): boolean {
    let indexMap: Map<number, number> = new Map();
    for (let i: number = 0; i < nums.length; i++) {
        let temp = indexMap.get(nums[i])
        if (temp != undefined && i - temp <= k) {
            return true
        }
        indexMap.set(nums[i], i)
    }
    return false
};

/**
 * 219. 存在重复元素 II
 */
test('leetcode_219', () => {
    console.log(containsNearbyDuplicate([1, 2, 3, 1], 3))
    console.log(containsNearbyDuplicate([1, 0, 1, 1], 1))
    console.log(containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2))
});