package leetcode

/**
Given an integer n, return all the structurally unique BST's (binary search
trees), which has exactly n nodes of unique values from 1 to n. Return the answer
in any order.


 Example 1:


Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]


 Example 2:


Input: n = 1
Output: [[1]]



 Constraints:


 1 <= n <= 8


 Related Topics 树 二叉搜索树 动态规划 回溯 二叉树 👍 1313 👎 0

*/
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTree(1, n)
}

func generateBSTree(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBSTree(start, i-1)
		right := generateBSTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}

//leetcode submit region end(Prohibit modification and deletion)
