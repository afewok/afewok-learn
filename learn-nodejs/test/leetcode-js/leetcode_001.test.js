var twoSum = function(nums, target) {
    let myMap = new Map();
    for (let i = 0; i < nums.length; i++) {
        const find = target - nums[i];
        if (myMap.get(find) != undefined) {
            return [myMap.get(find), i];
        }
        myMap.set(nums[i], i);
    }
    return [];
};

/**
 * 1. 两数之和
 */
test('leetcode_001', () => {
    console.log(twoSum([2, 7, 11, 15], 9))
});