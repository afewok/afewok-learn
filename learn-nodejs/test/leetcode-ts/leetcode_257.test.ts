function binaryTreePaths(root: TreeNode | null): string[] {
};


/**
 * 257. 二叉树的所有路径
 */
test('leetcode_257', () => {
    console.log(binaryTreePaths(new TreeNode(1, new TreeNode(2, new TreeNode(5, null, null), null), new TreeNode(3, null, null))))
});


class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}