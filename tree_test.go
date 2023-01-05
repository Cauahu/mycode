package mycode

import (
	"fmt"
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Preorder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Val, " ")
	Preorder(n.Left)
	Preorder(n.Right)
}

func levelorder(root *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	i := 0
	last := root
	for i < len(queue) {
		n := queue[i]
		if n == nil {
			fmt.Print("NULL ")
		} else {
			fmt.Print(n.Val, " ")
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
			if n.Left != nil {
				last = n.Left
			}
			if n.Right != nil {
				last = n.Right
			}

			if last == n {
				return
			}
		}

		i++
	}
}

func Test_bstFromPreorder(t *testing.T) {
	preorder := []int{8, 5, 1, 7, 10, 12}
	head := bstFromPreorder(preorder)
	Preorder(head)
}

/*
 给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。

 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。

 二叉搜索树 是一棵二叉树，其中每个节点，Node.left的任何后代的值 严格小于 Node.val,Node.right的任何后代的值 严格大于 Node.val。

 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。
输入：preorder = [8,5,1,7,10,12]
输出：[8,5,10,1,7,null,12]
*/
// 1008
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}
	i := 1
	for i < len(preorder) && preorder[i] < preorder[0] {
		i++
	}

	node.Left = bstFromPreorder(preorder[1:i])
	if i < len(preorder) {
		node.Right = bstFromPreorder(preorder[i:])
	}

	return node
}

/*
给定二叉树的根节点root，找出存在于 不同 节点A 和B之间的最大值 V，其中V = |A.val - B.val|，且A是B的祖先。
（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
输出：7
*/
// 1080
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	var traverse func(*TreeNode, int, func(int, int) bool) int
	traverse = func(node *TreeNode, juge int, compare func(int, int) bool) int {
		if node == nil {
			return juge
		}

		maxleft := traverse(node.Left, juge, compare)
		maxright := traverse(node.Right, juge, compare)

		if compare(maxleft, maxright) {
			maxleft = maxright
		}

		if maxleft != juge {
			j := int(math.Abs(float64(node.Val - maxleft)))
			if j > res {
				res = j
				fmt.Println(res, node.Val, maxleft)
			}
		}

		if compare(node.Val, maxleft) {
			return maxleft
		}

		return node.Val
	}

	greater := func(i, j int) bool {
		if i > j {
			return true
		}
		return false
	}

	smaller := func(i, j int) bool {
		if i < j {
			return true
		}
		return false
	}

	traverse(root, math.MinInt, smaller)
	traverse(root, math.MaxInt, greater)

	return res
}

func Test_maxAncestorDiff(t *testing.T) {
	//[8,3,10,1,6,null,14,null,null,4,7,13] 7
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:  10,
			Left: nil,
			Right: &TreeNode{
				Val: 14,
				Left: &TreeNode{
					Val:   13,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
	}
	t.Log(maxAncestorDiff(root))
}

/* 1080
给定一棵二叉树的根 root，请你考虑它所有从根到叶的路径：从根到任何叶的路径。（所谓一个叶子节点，就是一个没有子节点的节点）
假如通过节点 node 的每种可能的 “根-叶” 路径上值的总和全都小于给定的 limit，则该节点被称之为「不足节点」，需要被删除。
请你删除所有不足节点，并返回生成的二叉树的根。
输入：root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
输出：[1,2,3,4,null,null,7,8,9,null,14]
*/
func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var traverse func(node *TreeNode, sum int) *TreeNode
	traverse = func(node *TreeNode, sum int) *TreeNode {
		if node == nil {
			return nil
		}

		sum += node.Val
		if node.Left == nil && node.Right == nil {
			if sum < limit {
				return nil
			} else {
				return node
			}
		}

		node.Left = traverse(node.Left, sum)
		node.Right = traverse(node.Right, sum)

		if node.Left == nil && node.Right == nil {
			return nil
		}

		return node
	}

	return traverse(root, 0)

	//[1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
	//[1,2,3,4,null,null,7,8,9,null,14]
}

func Test_sufficientSubset(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 19,
				},
			},
			Right: &TreeNode{
				Val: -99,
				Left: &TreeNode{
					Val: -99,
				},
				Right: &TreeNode{
					Val: -99,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: -99,
				Left: &TreeNode{
					Val: 12,
				},
				Right: &TreeNode{
					Val: 13,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: -99,
				},
				Right: &TreeNode{
					Val: 14,
				},
			},
		},
	}

	r := sufficientSubset(root, 1)
	levelorder(r)
}

/* 1110
给出二叉树的根节点root，树上每个节点都有一个不同的值。
如果节点值在to_delete中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案。
输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
输出：[[1,2,null,4],[6],[7]]
*/

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res := make([]*TreeNode, 0)
	dump := &TreeNode{
		Val:  1001,
		Left: root,
	}

	deletemap := make(map[int]struct{}, len(to_delete))
	for _, n := range to_delete {
		deletemap[n] = struct{}{}
	}

	var traverse func(node *TreeNode) *TreeNode
	traverse = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = traverse(node.Left)
		node.Right = traverse(node.Right)

		if _, ok := deletemap[node.Val]; ok {
			if node.Left != nil {
				res = append(res, node.Left)
			}

			if node.Right != nil {
				res = append(res, node.Right)
			}
			return nil
		}

		return node
	}
	traverse(dump)
	if dump.Left != nil {
		res = append(res, root)
	}
	return res
}

/* 113
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
[[5,4,11,2],[5,8,4,5]]
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	copysli := func(s []int) {
		dump := make([]int, 0, len(s))
		for _, n := range s {
			dump = append(dump, n)
		}
		res = append(res, dump)
	}

	var traverse func(node *TreeNode, sum int, s []int)
	traverse = func(node *TreeNode, sum int, s []int) {
		if node == nil {
			return
		}

		sum += node.Val
		s = append(s, node.Val)

		if node.Left == nil && node.Right == nil && sum == targetSum {
			copysli(s)
			return
		}

		traverse(node.Left, sum, s)
		traverse(node.Right, sum, s)

		s = s[:len(s)-1]
		return
	}
	traverse(root, 0, []int{})
	return res
}

func Test_pathSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	t.Log(pathSum(root, 22))

}

/* 114
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
*/
func flatten(root *TreeNode) {
	var traverse func(node *TreeNode) *TreeNode
	traverse = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		left := traverse(node.Left)   // 2-3-4
		right := traverse(node.Right) // 5-6

		if left != nil {
			node.Right = left

			for left.Right != nil {
				left = left.Right
			}

			left.Right = right
			node.Left = nil
		} else {
			node.Right = right
		}

		return node
	}

	root = traverse(root)
}

func Test_flatten(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)
	levelorder(root)
}

/*116
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL。
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	i, levelnode := 0, root

	for i != len(queue) {
		n := queue[i]

		if n.Left != nil {
			queue = append(queue, n.Left)
		}

		if n.Right != nil {
			queue = append(queue, n.Right)
		}

		if levelnode == n {
			n.Next = nil
			levelnode = queue[len(queue)-1]
		} else {
			n.Next = queue[i+1]
		}

		i++
	}

	return root
}

func Test_connect(t *testing.T) {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}
	connect(root)
	n := root
	for n != nil {
		nextlevel := n.Left
		for n != nil {
			t.Log(n.Val)
			n = n.Next
		}
		t.Log("NULL")
		n = nextlevel
	}
}

/*124
路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中至多出现一次。该路径至少包含一个节点，且不一定经过根节点。
路径和是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其最大路径和。
*/
func maxPathSum(root *TreeNode) int {
	res := root.Val

	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftmax := traverse(node.Left)
		rightmax := traverse(node.Right)

		if res < node.Val+leftmax+rightmax {
			res = node.Val + leftmax + rightmax
		}

		if leftmax < rightmax {
			leftmax = rightmax
		}

		if leftmax > 0 {
			node.Val += leftmax
		}

		if res < node.Val {
			res = node.Val
		}

		return node.Val
	}

	traverse(root)
	return res
}

/*129
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
*/
func sumNumbers(root *TreeNode) int {
	res := 0
	var traverse func(node *TreeNode, pre int)
	traverse = func(node *TreeNode, pre int) {
		pre = pre*10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += pre
			return
		}

		if node.Left != nil {
			traverse(node.Left, pre)
		}

		if node.Right != nil {
			traverse(node.Right, pre)
		}
	}
	traverse(root, 0)
	return res
}

func Test_sumNumbers(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}

	t.Log(sumNumbers(root))
}

/*1315
给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
如果不存在祖父节点值为偶数的节点，那么返回0 。
*/
func sumEvenGrandparent(root *TreeNode) int {
	var traverse func(node *TreeNode, father, grandfather int) int
	traverse = func(node *TreeNode, father, grandfather int) int {
		if node == nil {
			return 0
		}

		sum := traverse(node.Left, node.Val, father) + traverse(node.Right, node.Val, father)
		if grandfather%2 == 0 {
			return node.Val + sum
		}

		return sum
	}
	return traverse(root, -1, -1)
}

func Test_sumEvenGrandparent(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}
	t.Log(sumEvenGrandparent(root))
}

/*1325
给你一棵以root为根的二叉树和一个整数target，请你删除所有值为target 的叶子节点 。
注意，一旦删除值为target的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是target ，那么这个节点也应该被删除。
也就是说，你需要重复此过程直到不能继续删除。
*/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Val == target && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}

func Test_removeLeafNodes(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	levelorder(removeLeafNodes(root, 2))
}

/* 1339
给你一棵二叉树，它的根为root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
*/
func maxProduct(root *TreeNode) int {
	sum := 0
	var sumnode func(node *TreeNode) int
	sumnode = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return node.Val + sumnode(node.Left) + sumnode(node.Right)
	}
	sum = sumnode(root)
	fmt.Println("sum=", sum)

	res := 0
	div := math.Pow10(9) + 7
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		left := traverse(node.Left)
		right := traverse(node.Right)

		tl := left * (sum - left)
		if res < tl {
			res = tl
		}

		tr := right * (sum - right)
		if res < tr {
			res = tr
		}

		return node.Val + left + right
	}
	traverse(root)
	fmt.Println(res, div)
	return res % int(div)
}

func Test_maxProduct(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	t.Log(maxProduct(root))
}

/*1367
给你一棵以root为根的二叉树和一个head为第一个节点的链表。
如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以head为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。
一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
*/
func isSubPath(head *ListNode, root *TreeNode) bool {
	var traverse func(l *ListNode, n *TreeNode) bool
	traverse = func(l *ListNode, n *TreeNode) bool {
		if l == nil {
			return true
		}

		if n == nil {
			return false
		}

		if l.Val == n.Val {
			return traverse(l.Next, n.Left) || traverse(l.Next, n.Right)
		}

		return false
	}

	if root == nil {
		return false
	}

	return traverse(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func Test_isSubPath(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 8,
			},
		},
	}
	t.Log(isSubPath(head, root))
}

/*1372
给你一棵以root为根的二叉树，二叉树中的交错路径定义如下：
选择二叉树中 任意节点和一个方向（左或者右）。
如果前进方向为右，那么移动到当前节点的的右子节点，否则移动到它的左子节点。
改变前进方向：左变右或者右变左。
重复第二步和第三步，直到你在树中无法继续移动。
交错路径的长度定义为：访问过的节点数目 - 1（单个节点的路径长度为 0 ）。
请你返回给定树中最长 交错路径的长度。
*/
func longestZigZag(root *TreeNode) int {
	res := 0
	var zigzag func(node *TreeNode, direct, len int)
	zigzag = func(node *TreeNode, direct, len int) {
		if node == nil {
			if len > res {
				res = len
			}
			return
		}

		if direct == 0 {
			zigzag(node.Right, 1, len+1)
			zigzag(node.Left, 0, 1)
		} else {
			zigzag(node.Left, 0, len+1)
			zigzag(node.Right, 1, 1)
		}
	}

	zigzag(root, 0, 0)
	zigzag(root, 1, 0)

	return res - 1
}

func Test_longestZigZag(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	t.Log(longestZigZag(root))

	root = &TreeNode{
		Val: 1,
	}
	t.Log(longestZigZag(root))

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	t.Log(longestZigZag(root))
}

/*1373
给你一棵以root为根的二叉树，请你返回 任意二叉搜索子树的最大键值和。
二叉搜索树的定义如下：
任意节点的左子树中的键值都小于此节点的键值。
任意节点的右子树中的键值都 大于此节点的键值。
任意节点的左子树和右子树都是二叉搜索树。
*/
func maxSumBST(root *TreeNode) int {
	res := math.MinInt
	var traverse func(node *TreeNode) (bool, int, int, int)
	traverse = func(node *TreeNode) (bool, int, int, int) {
		istrue := true
		max := node.Val
		min := node.Val
		sum := node.Val

		if node.Left != nil {
			isleft, sumleft, maxleft, minleft := traverse(node.Left)
			if !isleft || node.Val <= maxleft {
				istrue = false
			} else {
				if min > minleft {
					min = minleft
				}

				sum += sumleft
			}
		}

		if node.Right != nil {
			isright, sumright, maxright, minright := traverse(node.Right)
			if !isright || node.Val >= minright {
				istrue = false
			} else {
				if max < maxright {
					max = maxright
				}
				sum += sumright
			}
		}

		if istrue && sum > res {
			res = sum
		}

		return istrue, sum, max, min
	}

	traverse(root)
	if res == math.MinInt || res < 0 {
		res = 0
	}

	return res
}

func Test_maxSumBST(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
	}
	t.Log(maxSumBST(root))

	root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	}
	t.Log(maxSumBST(root))

	root = &TreeNode{
		Val: -4,
		Left: &TreeNode{
			Val: -2,
		},
		Right: &TreeNode{
			Val: -5,
		},
	}
	t.Log(maxSumBST(root))

	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t.Log(maxSumBST(root))

	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	t.Log(maxSumBST(root))

	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: -5,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	t.Log(maxSumBST(root))
}

/*1448
给你一棵根为root的二叉树，请你返回二叉树中好节点的数目。
「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。
*/
func goodNodes(root *TreeNode) int {
	res := 0
	var traverse func(node *TreeNode, maxval int)
	traverse = func(node *TreeNode, maxval int) {
		if node == nil {
			return
		}

		if node.Val >= maxval {
			maxval = node.Val
			res++
		}

		traverse(node.Left, maxval)
		traverse(node.Right, maxval)
	}

	traverse(root, math.MinInt)
	return res
}

/*1457
给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
请你返回从根到叶子节点的所有路径中伪回文路径的数目。
*/
func pseudoPalindromicPaths(root *TreeNode) int {
	res := 0
	var traverse func(node *TreeNode, a []int, od int)
	traverse = func(node *TreeNode, a []int, od int) {
		if node == nil {
			return
		}

		a[node.Val] ^= 1
		if a[node.Val] == 1 {
			od++
		} else {
			od--
		}

		if node.Left == nil && node.Right == nil {
			if od <= 1 {
				res++
			}
		} else {
			traverse(node.Left, a, od)
			traverse(node.Right, a, od)
		}

		a[node.Val] ^= 1
	}

	a := make([]int, 10)
	traverse(root, a, 0)

	return res
}

func Test_pseudoPalindromicPaths(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	t.Log(pseudoPalindromicPaths(root))
}

/*508
给你一个二叉树的根结点root，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
*/
func findFrequentTreeSum(root *TreeNode) []int {
	summap := make(map[int]int, 0)
	maxnum := 0
	var traverse func(node *TreeNode) int
	traverse = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := traverse(node.Left)
		right := traverse(node.Right)
		sum := node.Val + left + right

		summap[sum]++
		if summap[sum] > maxnum {
			maxnum = summap[sum]
		}

		return sum
	}

	traverse(root)

	res := make([]int, 0)
	for i, n := range summap {
		if n == maxnum {
			res = append(res, i)
		}
	}

	return res
}

func Test_findFrequentTreeSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}

	t.Log(findFrequentTreeSum(root))
}

/*513
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lastnode, leftnode := root, root
	i := 0
	for i < len(queue) {
		node := queue[i]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		i++
		if node == lastnode {
			lastnode = queue[len(queue)-1]
			if i < len(queue) {
				leftnode = queue[i]
			}
		}
	}

	return leftnode.Val
}

func Test_findBottomLeftValue(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	t.Log(findBottomLeftValue(root))
}
