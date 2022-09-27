package leetcode

/**
Given the root of a binary tree, return the inorder traversal of its nodes'
values.


 Example 1:


Input: root = [1,null,2,3]
Output: [1,3,2]


 Example 2:


Input: root = []
Output: []


 Example 3:


Input: root = [1]
Output: [1]



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100



Follow up: Recursive solution is trivial, could you do it iteratively?

 Related Topics 栈 树 深度优先搜索 二叉树 👍 1575 👎 0

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
var res = []int{}
func inorderTraversal(root *TreeNode) []int {
	res=[]int{}
	inorder(root)
	return res
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	res=append(res,root.Val)
	inorder(root.Right)
}

//leetcode submit region end(Prohibit modification and deletion)
