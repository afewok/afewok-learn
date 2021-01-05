function twoSum(nums: number[], target: number): number[] {
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

test('leetcode_001', () => {
    let list: number[] = [2, 7, 11, 15];
    console.log(twoSum(list, 9))
    // expect()
});