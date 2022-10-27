package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func main()  {
	//nums := []int{3,3}
	//target := 6
	//fmt.Println(twoSum(nums, target))
	//l1 := &ListNode{
	//	Val:  2,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//l2 := &ListNode{
	//	Val:  5,
	//	Next: &ListNode{
	//		Val:  6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	//fmt.Println(addTwoNumbers(l1, l2))
	//s:="bb"
	//fmt.Println(longestPalindrome(s))
	//height := []int{1,8,6,2,5,4,8,3,7}
	//fmt.Println(maxArea(height))
	//nums := []int{-1,0,1,2,-1,-4}
	//fmt.Println(threeSum(nums))
	//digits := "23"
	//fmt.Println(letterCombinations(digits))
	//fmt.Println(removeNthFromEnd(l1,1))
	//s := "(])"
	//fmt.Println(isValid(s))
	//fmt.Println(generateParenthesis(3))
	//arr := []int{5, 7, 7, 8, 8, 10}
	//taget := 6
	//fmt.Println(searchRange(arr,taget))
	//arr := []int{5, 7, 7, 8, 8, 10}
	//taget := 6
	//fmt.Println(combinationSum(arr,taget))
	//sortArr := []int{2,33,4,55,6,77,34,25,67,87,24,12,23,45,68,90}
	//fmt.Println(bucketSort(sortArr, 100))
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//fmt.Println(groupAnagrams(strs))
	nums := []int{2,0,2,1,1,0}
	sortColorsNew(nums)
	fmt.Println(nums)
}
/*
43. 从前序与中序遍历序列构造二叉树


*/
func buildTree(preorder []int, inorder []int) *TreeNode {

}

/*
42. 二叉树的最大深度

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Right), maxDepth(root.Left)) + 1
}

/*
41. 二叉树的层序遍历

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return levelFind([]*TreeNode{root})
}
func levelFind(levelNodes []*TreeNode) [][]int {
	resp := make([][]int, 0)
	temp := make([]int, 0)
	nextLevelNodes := make([]*TreeNode, 0)
	if len(levelNodes) <= 0 {
		return resp
	}

	for _, node := range levelNodes {
		temp = append(temp, node.Val)
		if node.Left != nil {
			nextLevelNodes = append(nextLevelNodes, node.Left)
		}
		if node.Right != nil {
			nextLevelNodes = append(nextLevelNodes, node.Right)
		}
	}
	resp = [][]int{temp}
	return append(resp, levelFind(nextLevelNodes)...)
}
/*
40. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。

输入：root = [1,2,2,3,4,4,3]
输出：true
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left ,right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}

	if right == nil && left != nil {
		return false
	}

	if right == left && left == nil {
		return true
	}

	if right.Val != left.Val {
		return false
	}

	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

/*
39. 验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

输入：root = [2,1,3]
输出：true

中序遍历，可以输出该树的一个数组,判断该树的数组是否是递增数组
 */
func isValidBSTNew(root *TreeNode) bool {
	return helper(root, root.Left.Val, root.Right.Val)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
func isValidBST(root *TreeNode) bool {

	// 中序遍历
	rootArr := midTraverse(root)
	// 判断该数字是否是递增数组
	for i := 1; i< len(rootArr); i++ {
		if rootArr[i] <= rootArr[i-1] {
			return false
		}
	}
	return true
}
func midTraverse(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := midTraverse(root.Left)
	left = append(left, root.Val)
	right := midTraverse(root.Right)
	return  append(left, right...)
}
/*
38. 不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

输入：n = 3
输出：5

动态规划
给定一个有序序列1...n，为了构建出一颗二叉树，我们可以遍历每个数字i，将数字欧威树根，将[1...i-1]序列作为左子树，将[i+1, n]作为右子树。
在上述构建过程中由于根值的不同，因此我们能保证每颗二叉搜索树都是唯一的。因此，原问题可以被分解成规模较小的两个子问题

算法
为此我们定义两函数：
	g(n) : 表示长度为n的序列能构成的不同二叉搜索树的个数
	f(i,n): 表示以i为根节点的序列长度为n不同二叉搜索树的个数
可见，g(n) 是我们求解需要的函数。

不同二叉树搜索树的总和g(n), 是对遍历所有i (0<=i<=n)的f(i,n)之和。对于边界情况，当序列长度为 1（只有根）或为 0（空树）时，只有一种情况
g(0) = g(1) = 1
给定序列 1⋯n，我们选择数字 i作为根，则根为 i的所有二叉搜索树的集合是左子树集合和右子树集合的笛卡尔积。举例而言，创建以 3为根、
长度为 7 的不同二叉搜索树，整个序列是 [1,2,3,4,5,6,7]，我们需要从左子序列
[1,2] 构建左子树，从右子序列 [4,5,6,7] 构建右子树，然后将它们组合（即笛卡尔积）。

对应方程为：F(i,n)=G(i−1)⋅G(n−i)

因此 g[n] = ∑ f(i,n) (0<=i<=n)

因此转移方程就是
	g[n] = ∑ g(i-1) * g(n-i)   (0<=i<=n)

*/
func numTrees(n int) int {
	G := make([]int, n + 1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i ; j++ {
			G[i] += G[j] * G[i-j]
		}
	}
	return G[n]
}

/*
37.
 */
/*
36. 二叉树的后序遍历

给定一个二叉树的根节点 root ，返回 它的 后序 遍历 。

输入：root = [1,null,2,3]
输出：[1,3,2]
*/

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 优先遍历左子树
	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)
	return append(left, append(right, root.Val)...)
}


/*
35. 二叉树的前序遍历

给定一个二叉树的根节点 root ，返回 它的 前序 遍历 。

输入：root = [1,null,2,3]
输出：[1,3,2]
*/

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 优先遍历左子树
	temp := make([]int, 0)
	temp = append(temp, root.Val)
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	return append(temp, append(left, right...)...)
}

/*
34. 二叉树的中序遍历

给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

输入：root = [1,null,2,3]
输出：[1,3,2]
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 优先遍历左子树
	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	// 遍历右子树
	right := inorderTraversal(root.Right)
	return append(left, right...)
}
/*
33. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

回溯法，需要在探索的过程中需要返回
 */
func exist(board [][]byte, word string) bool {
	trace := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		trace[i] = make([]bool, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			trace[i][j] = true
		}
	}
	var  dfs func(length int, i int ,j int) bool
	dfs = func(length int, i int, j int) bool {
		if length == len(word){
			return true
		}

		if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
			return false
		}

		if !trace[i][j]  || word[length] != board[i][j]{
			return false
		}

		trace[i][j] = false
		// 往上走
		if dfs(length+1,i-1,j) {
			return true
		}
		// 往下走
		if dfs(length+1,i+1,j) {
			return true
		}
		// 往左走
		if dfs(length+1,i,j+1) {
			return true
		}
		// 往右走
		if dfs(length+1,i,j-1) {
			return true
		}
		trace[i][j] = true
		return  false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] && dfs( 0,i,j) {
				return true
			}
		}
	}

	return false
}


/*
32. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集不能包含重复的子集。你可以按任意顺序返回解集。

输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

dfs 深度优先遍历搜索
*/
func subsets(nums []int) [][]int {
	set := []int{}
	ans := make([][]int, 0)
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

/*
31. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库的sort函数的情况下解决这个问题。

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]

两次循环+原地的排序算法；插入排序

双指针
*/
func sortColors(nums []int)  {
	if len(nums) <= 0 {
		return
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0  && nums[j] < nums[j-1] ;j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return
}
func sortColorsNew(nums []int)  {
	index0,index1 := 0,0
	for i, num := range nums {
		if num == 0 {
			nums[index0] ,nums[i] =  nums[i], nums[index0]
			if index0 < index1 {
				nums[i], nums[index1] = nums[index1], nums[i]
			}
			index0++
			index1++
		} else if num == 1 {
			nums[index1] ,nums[i] = nums[i], nums[index1]
			index1++
		}
	}
}
/*
30. 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

动态规划
dp[i][j]表示以下标i-1为结尾的字符串word1，和以下标j-1为结尾的字符串word2，最近编辑距离为dp[i][j]

在确定递推公式的时候，首先需要考虑清楚编辑的几种操作：
	if word1[i-1] == word2[j-1]
		不操作
	if word1[i-1] != word2[j-1]
		增
		删
		换
	以上四种情况中，
	if word1[i-1] == word2[j-1]说明不用做任何编辑，dp[i][j]就是dp[i-1][j-1]，即dp[i][j]=dp[i-1][j-1]， world1[i-1]和word2[j-1]相等了，那么就不用再编辑了。

	if word1[i-1] != word2[j-1]就需要编辑了，但是如何编辑：
		* 操作一：word1删除一个元素，那么就是以下标i-2结尾的word1与j-1为结尾的word2的最近编辑距离再加上一个操作，即dp[i][j] = dp[i-1][j]+1
		* 操作二：word2删除一个元素，那么就是以下标i-1为结尾的word1与j-2为结尾的word2的最近编辑距离再加上一个操作，即dp[i][j] = dp[i][j-1]+1
		* 操作三：替换元素，word1替换word1[i - 1]，使其与word2[j - 1]相同，此时不用增加元素，那么以下标i-2为结尾的word1 与 j-2为结尾的word2的最近编辑距离 加上一个替换元素的操作。即 dp[i][j] = dp[i - 1][j - 1] + 1;

	综上，当 if (word1[i - 1] != word2[j - 1]) 时取最小的，即：dp[i][j] = min({dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1]}) + 1;

dp数组如何初始化
	dp[i][0] ：以下标i-1为结尾的字符串word1，和空字符串word2，最近编辑距离为dp[i][0]。
	那么dp[i][0]就应该是i，对word1里的元素全部做删除操作，即：dp[i][0] = i;
	同理dp[0][j] = j;
 */
func minDistance(word1 string, word2 string) int {
	m,n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		dp[i][0] = i // word1[i] 变成 word2[0], 删掉 word1[i], 需要 i 部操作
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = j // word1[0] 变成 word2[j], 插入 word1[j]，需要 j 部操作
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else { // Min(插入，删除，替换)
				dp[i][j] = min(dp[i][j-1], min(dp[i-1][j], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

/*
29.爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

动态规划
f[i] 表示到达当前楼梯需要的方法
那么到达f[i]，可以从f[i-1] + 1,或者f[i-2]+2 转移过来
因此，转移方程为：f[i] =  f[i-1] + f[i-2]
f[1]= 1
f[2] = 2
 */
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
/*
28. 最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

动态规划
和下面的不同路径比较类似，f[i,j]可以从f[i-1,j]或者f[i,j-1]转移过来，
不过在转移时需要按照最小的去转移，所以转移方程为：f[i, j] = min(f[i-1, j] , f[i,j-1]) + grid[i, j]
但是这里的f[0.j]和f[i,0]是由grid转移过来，所以需要进行累加：
f[i,0] = grid[i, 0] + f[i-1, 0]
f[0,j] = grid[0, j] + f[0, j-1]

 */

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i<len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i :=1 ; i< len(grid); i++ {
		for j:= 1;j < len(grid[0]);j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}


/*
27. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

输入：m = 3, n = 7
输出：28

动态规划
我们用f[i][j]表示从左上角走到[i,j]的路径数量，其中i，j的范围分别是[0,m)和[0,n)
由于我们每一个步只能向下或者向右移动，因此，我们要想走到(i, j),只能从(i-1,j)或者(i，j-1)转移过来，
因此可知道转移方程为
f[i,j] = f[i-1, j] + f[i, j-1]
如果了i = 0或者j=0，那么f[i-1,j],f[i,j-1]这一项是非法的。并且最开始的f[0,0] = 1
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i:= 0; i< m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j< n ; j++ {
		dp[0][j] = 1
	}
	for i:= 1; i<m ;i++ {
		for j:= 1;j<n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

/*
26. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 */

func mergeNew(intervals [][]int) [][]int {
	resp := make([][]int, 0)
	if len(intervals) <= 0 {
		return resp
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	resp = append(resp, append([]int{}, intervals[0][0], intervals[0][1]))
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= resp[len(resp)-1][1] {
			if intervals[i][1] > resp[len(resp)-1][1] {
				resp[len(resp)-1][1] = intervals[i][1]
			}
		} else {
			resp = append(resp, intervals[i])
		}
	}
	return resp
}

/*
25. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

动态规划
当前能跳跃的最远距离，这个距离内的所有点的新距离可以更新最远距离，否则无法到最远
*/

func canJump(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}

	dp := make([]bool,len(nums))
	dp[0] = true
	for i := 1;i<len(nums); i++ {
		for j := i-1; j>=0; j-- {
			// nums[j] + j表示你可以跳跃的最远距离，从当前位置开始
			if dp[j] && nums[j] + j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}
/*
24. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。

动态规划

f[i]表示以第i个元素结尾的连续子数组的最大和，那么很明显我们需要求的是max(f[i])

转移方式：f[i] = max(nums[i], f[i-1]+nums[i])

初始条件：max ：= nums[0]
		dp[0] = nums[0]
 */
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i< len(nums); i++  {
		dp[i] = nums[i]
		if dp[i-1] + nums[i] > dp[i] {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

/*
23. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

排序+哈希
*/
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _,str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _,v := range mp {
		ans = append(ans, v)
	}
	return ans
}

/*
22.
冒泡排序
选择排序
快速排序
归并排序
插入排序
 */
func sortTemp(arr []int) {
	// 冒泡排序
	// 时间复杂度：o(n^2)
	// 比较相邻的元素，如果第一个比第二个大，责交换他们，
	tempArr0 := make([]int, len(arr))
	copy(tempArr0, arr)
	for i := len(arr)-1; i >= 0 ; i-- {
		for j := 0; j<i; j++ {
			if tempArr0[j] > tempArr0[j+1] {
				tempArr0[j] , tempArr0[j+1] = tempArr0[j+1], tempArr0[j]
			}
		}
	}
	fmt.Println(tempArr0)


	// 选择排序
	// 时间复杂度：o(n^2)
	// 固定一位元素，和其他的元素进行对比，按照升序的方式进行对比和交换
	// 在未排序的序列中找到最大的元素，放在起始位置，然后再从未排序的序列中寻找第二大的元素放在第二个位置
	tempArr1 := make([]int, len(arr))
	copy(tempArr1, arr)
	for i := 0; i < len(tempArr1); i++ {
		minIndex := i
		for j := i; j < len(tempArr1); j++ {
			if tempArr1[minIndex] > tempArr1[j] {
				minIndex = j
			}
		}
		tempArr1[minIndex] , tempArr1[i]  = tempArr1[i] , tempArr1[minIndex]
	}
	fmt.Println(tempArr1)

	// 插入排序
	// 时间复杂度：
	// 将当前右边的第一个元素，插入到左边已排好序中，
	tempArr2 := make([]int, len(arr))
	copy(tempArr2, arr)
	for i:= 1; i< len(tempArr2); i++ {
		for j := i; j > 0 &&  tempArr2[j] < tempArr2[j-1]; j-- {
			tempArr2[j], tempArr2[j-1] = tempArr2[j-1], tempArr2[j]
		}
	}
	fmt.Println(tempArr2)

}

// 快速排序
func quickSort(arr []int, left int, right int)  {
	if left < right {
		flag := arr[left]
		j := left
		for i := left; i< right; i++ {
			if arr[i] < flag {
				j++
				arr[j] , arr[i]  = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		quickSort(arr,j+1, right)
		quickSort(arr,left, j)
	}
}
// 桶排序
func bucketSort(nums []int, bucketSize int) []int {
	// 获取最大值和最小值
	minVal, maxVal := 0,0
	resp := make([]int, 0)
	for _, val := range nums {
		if minVal > val {
			minVal = val
		}
		if maxVal < val {
			maxVal = val
		}
	}

	// 分桶
	buckets := make([][]int, (maxVal-minVal)/bucketSize+1)
	// 将数据分布到桶中
	for _, num := range nums {
		buckets[(num-minVal)/bucketSize] = append(buckets[(num-minVal)/bucketSize], num)
	}

	for _, bucket := range buckets {
		if len(bucket) > 0 {
			quickSort(bucket, 0, len(bucket))
		}
		resp = append(resp, bucket...)
	}
	return  resp
}
// 归并排序
// 将数组分成两部分，分别对两部分进行排序，然后在合并两个有序的部分
func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) / 2
	left := mergeSort(list[:mid])
	right := mergeSort(list[mid:])
	return merge(left, right)
}
func merge(list1, list2 []int) []int {
	len1, len2 := len(list1), len(list2)
	index1, index2 := 0,0
	resp := make([]int, 0)
	for index1 < len1 && index2 < len2 {
		if list1[index1] < list2[index2] {
			resp = append(resp, list1[index1])
			index1++
		} else {
			resp = append(resp, list2[index2])
			index2++
		}
	}
	if index2 == len2 {
		resp = append(resp, list1[index1:]...)
	}
	if index1 == len1 {
		resp = append(resp, list2[index2:]...)
	}
	return resp
}
/*
21. 旋转图像

给定一个n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。


输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

第一行的旋转90度后，出现在最后一列
第二行旋转90度后，出现在倒数第二列
以此类推，最后一行出现在旋转后的第一列
*/
func rotate(matrix [][]int)  {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := 0; i< n ; i++ {
		tmp[i] = make([]int, n)
	}
	for i, ints := range matrix {
		for i2, i3 := range ints {
			tmp[i2][n -1 - i] =  i3
		}
	}
	copy(matrix, tmp)
	return
}

/*
用翻转代替旋转
先通过水平轴上下翻转。在根据右主对角线翻转得到
 */

func rotateNew(matrix [][]int)  {
	n := len(matrix)
	// 水平翻转
	for i:= 0; i < n /2 ; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j] , matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/*
20. 全排列

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

// 回溯法
1. 确认递归方法

2. 确定终止条件
	长度等于nums的长度
3. 确定单层遍历逻辑
	循环每个元素
*/
func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	for i, num := range nums {
		temp := make([]int, len(nums)-1)
		copy(temp[0:], nums[0:i])
		copy(temp[i:],nums[i+1:])
		sub := permute(temp)
		for _, s := range sub {
			res = append(res, append(s, num))
		}

	}
	return res
}


/*
19. 接雨水

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
 */
func trap(height []int) int {
	var ret int
	if len(height) <= 2 {
		return ret
	}
	s := make([]int,0)
	for i, h := range height {
		for len(s) > 0 && height[s[len(s)-1]] < h {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			t := s[len(s)-1]
			w := i-t-1
			H := min(height[t], h) - height[top]
			ret += w * H
		}
		s = append(s, i)
	}
	return ret
}
func  min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
18. 组合总和

给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按任意顺序返回这些组合。
candidates中的同一个数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

寻找所有可行解的题，我们都可以尝试使用搜索回溯的方法来解决
我们定义递归函数dfs(target, combine,idx)表示当前在candidates数组的第idx位，还剩target要组合。已经组合的列表为combine。递归的终止条件为
target <= 0 或者candidates数组被全部使用完。那么在当前的函数中，每次我们可以选择跳过不用第idx个数，即执行dfs(target,combine,idx)。也可以选择
使用第idx个数，即执行dfs(target-candidates[idx],combine,idx)，注意到每个数组可以被无限制重复选取，因此搜索的下标仍为idx。这里的写法是不带剪枝

回溯法三部曲：
递归函数
终止条件
单层搜索逻辑
*/

func combinationSum(candidates []int, target int) [][]int {
	comb := []int{}
	ans := make([][]int, 0)
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans,append([]int(nil),comb...))
			return
		}
		// 直接跳过
		dfs(target,idx+1)
		// 选择当前数
		if target - candidates[idx] >=0 {
			comb = append(comb,candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}

/*
17，在排序数组中查找元素的第一个和最后一个位置

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]


通过收缩边界的方式来实现，，目的是为了寻找出target的左右边界
*/

func searchRange(nums []int, target int) []int {
	// 目标值开始位置：为目标值的左侧边界，无此值则返回比它大的数的左侧边界
	start := findBound(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1,-1}
	}
	// 目标值结束位置：为目标值+1的左侧边界-1，无此值则返回比它大的数的左侧边界-1
	end := findBound(nums, target+1) - 1
	return []int{start, end}
}

func findBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid-1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
/*
16. 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

复杂度为log n的算法

 */
func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l := 0
	r := len(nums)-1
	for l <= r {
		mid := (l+r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}


/*
15. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

出栈入栈的做法
对于每个遇到"("，我们将他的下标放入栈中
对于每个遇到")", 我们弹出栈顶元素进行匹配当前右括号
	如果找当前为空，说明当前的右括号为没有被匹配的右括号，我们将其放入栈低来更新当初我们认为的最后一个没有被匹配的左括号的下标
	如果栈不为空，当前右括号的下标减去栈顶元素的下班即为以该右括号为结尾的最长有效括号的长度，
循环一次，更新最大的子串长度

如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，我们在一开始的时候往栈中放入一个值为 −1的元素

 */
func longestValidParentheses(s string) int {
	lengthMax := 0
	stack := []int{}
	stack = append(stack, -1)
	for i :=0; i<len(s); i++ {
		if string(s[i]) == "(" {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else  {
				lengthMax = max(lengthMax, i - stack[len(stack)-1])
			}
		}
	}
	return lengthMax
}

/*
14. 下一个排列
整数数组的一个排列就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序
从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。
例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。

输入：nums = [1,2,3]
输出：[1,3,2]

1. 首先从后向前查找第一个顺序对(i, i+1)，满足a[i]<a[i+1]。这样较小数即为a[i]。此时[i+1,n]，必然为下降序列
2. 如果找到了顺序对，那么在区间[i+1, n)中从后向前查找第一个元素，满足a[i]<a[j]，这样较大数即为a[j]
3. 交换a[i]与a[j]，此时可以证明区间[i+1, n)必为降序，我们可以直接使用双指针翻转区间[i+1,n)使其变为升序，而无需对区间进行排序

 */
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 &&  nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := n-1
		for j >=0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}
func reverse(a []int)  {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}


/*
13. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表
 */

// 递归解法
func mergeKListsNew(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoListsNew(mergeKListsNew(lists[:len(lists)/2]), mergeKListsNew(lists[len(lists)/2:]))
}

//暴力解法
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	head := lists[0]
	for i:=1; i< len(lists); i++ {
		head = mergeTwoListsNew(head, lists[i])
	}
	return head
}
func mergeTwoListsNew(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val{
			temp.Val = list1.Val
			list1 = list1.Next
		} else {
			temp.Val = list2.Val
			list2 = list2.Next
		}
		temp.Next = &ListNode{}
		temp = temp.Next
	}
	if list1 == nil {
		temp.Val = list2.Val
		temp.Next = list2.Next
	}
	if list2 == nil {
		temp.Val = list1.Val
		temp.Next = list1.Next
	}
	return head
}

/*
12. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/
func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(lrmain int,rrmain int, path string)
	dfs = func(lrmain int,rrmain int, path string) {
		if 2 * n == len(path) {
			res = append(res , path)
			return
		}
		if lrmain > 0  {
			dfs(lrmain -1 , rrmain, path + "(")
		}
		if lrmain < rrmain {
			dfs(lrmain, rrmain-1, path+")")
		}
	}
	dfs(n ,n,"")
	return res
}

/*
11. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	temp := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Val = list1.Val
			temp.Next = &ListNode{}
			temp = temp.Next
			list1 = list1.Next
		} else {
			temp.Val = list2.Val
			temp.Next = &ListNode{}
			temp = temp.Next
			list2 = list2.Next
		}
	}
	if list2 == nil {
		temp.Val = list1.Val
		temp.Next= list1.Next
	}
	if list1 == nil {
		temp.Val = list2.Val
		temp.Next = list2.Next
	}
	return head
}

/*
10. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

输入：s = "()"
输出：true

 */
func isValid(s string) bool {
	if s == ""{
		return true
	}
	valSlice := make([]string, 0)
	for _, str := range []rune(s) {
		switch string(str) {
		case ")":
			if len(valSlice) == 0 {
				return false
			}
			if valSlice[len(valSlice)-1] == "(" {
				valSlice = valSlice[:len(valSlice)-1]
			} else {
				return false
			}
		case "}":
			if len(valSlice) == 0 {
				return false
			}
			if valSlice[len(valSlice)-1] == "{" {
				valSlice = valSlice[:len(valSlice)-1]
			} else {
				return false
			}
		case "]":
			if len(valSlice) == 0 {
				return false
			}
			if valSlice[len(valSlice)-1] == "[" {
				valSlice = valSlice[:len(valSlice)-1]
			} else {
				return false
			}
		default:
			valSlice = append(valSlice, string(str))
		}
	}
	if len(valSlice) == 0 {
		return true
	}
	return false
}
/*
9. 删除链表的倒数第N个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

双指针游走，建立两个指针，其间隔为n+1个节点的宽度，当第一个节点走到最后一个节点时，则第二个指针所指的就是应该删除的节点，操作当前节点指向下下个节点即可
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	left, right := dummy, dummy
	for i:= 0; i < n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}

/*
8 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

根据数字选择对应数字的字母映射,先获取对应的数字字母映射表，其次在循环当前每个字母，和已有的排列组合进行组合。
如果当前的排列组合的字符串数组为空，则需要将当前的字符串，切割，加入当前的字符串数组中，等待下一个字符串的排列组合

*/
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	resp := make([]string, 0)
	for _, v := range []rune(digits) {
		if _, ok := phoneMap[string(v)]; ok {
			resp = match(resp, phoneMap[string(v)])
		}
	}
	return resp
}

func match(str1 []string, str2 string) []string {
	resp := make([]string, 0)
	for _, s := range []rune(str2) {
		if len(str1) != 0 {
			for _, s2 := range str1 {
				resp = append(resp, fmt.Sprintf("%v%v",s2, string(s)))
			}
		} else {
			resp = append(resp,string(s))
		}
	}
	return resp
}

/*
7 三数之和

给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

二数之和的
 */
func threeSum(nums []int) [][]int {
	resp := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++  {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		right := len(nums)-1
		target := -1 * nums[i]
		for second := i+1; second < len(nums); second ++ {
			// 枚举不能重复
			if second > i+ 1 && nums[second] == nums[second - 1] {
				continue
			}
			for second < right && nums[second] + nums[right] > target {
				right--
			}
			if second == right {
				break
			}

			// 判断是否相等
			if nums[second] + nums[right] == target {
				resp = append(resp, []int{nums[i], nums[second], nums[right]})
			}
		}

	}
	return resp
}

/*
6
给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。

输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

两个字符都需要循环，'.'和'*' 需要单独处理，
'.'处理时需要知道上一个字符是什么，并且只匹配当前字符

'*'处理时需要知道上一个字符是什么，并且需要匹配多次

 */
func isMatch(s string, p string) bool {
	if p == "*" {
		return  true
	}
	i,j := 0,0
	for i < len(s) && j < len(p) {
		if p[j] == '*' {
			for s[i] != p[j-1] {

			}

		}

	}
	return true
}


/*
5
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。

找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

希望长，宽都是最大，这样盛水量才是最大

 */
func maxArea(height []int) int {
	maxArea := 0
	start, end := 0, len(height)-1
	for start < end {
		if height[start] < height[end] {
			maxArea = max(height[start] * (end- start), maxArea)
			start++
		} else {
			maxArea = max(height[end] * (end- start), maxArea)
			end--
		}
	}
	return maxArea
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
4
给你一个字符串 s，找到 s 中最长的回文子串。

确定回文子串的开始(begin)和长度(maxNum)，
先默认字符为奇数，按照奇数的方式求出以当前字符为中心的最长回文子串，更新begin和maxNum，确定开始和子串长度
在默认字符串为偶数，更新begin和maxNum，确定开始和子串长度
奇数和偶数在循环时的条件为，right<len(s),left>=0,s[left] == s[right]

最终返回begin和maxNum截取后的字符串
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	begin := 0
	maxNum := 1
	for i := 0; i < len(s); i++ {
		// 默认字符为奇数
		left := i - 1
		right := i + 1
		for right < len(s) && left >= 0 && s[left] == s[right] {
			if right-left+1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++
		}

		// 字符数为偶数
		left = i
		right = i + 1
		for right < len(s) && left >= 0 && s[left] == s[right] {
			if right-left+1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++
		}
	}
	return s[begin : begin+maxNum]
}

/*
3
两数之和等于固定的数值，借用map存已有的数值，目标值减去当前值后，查看差值是否在map中
 */
func twoSum(nums []int, target int) []int {
	keyValue := make(map[int]int, 0)
	resp := make([]int, 0)
	for i, num := range nums {
		subValue := target - num
		if _, ok := keyValue[subValue]; ok {
			return append(resp, i, keyValue[subValue])
		}
		keyValue[num] = i
	}
	return resp
}


/*
2
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。


判断当前两个链表node的val是否为0&next是否为nil

循环获取当前两个node的val，相加，

如果当前l1节点为nil，并且l2节点为nil， 并且相加进位为0 ，则退出循环

否则继续循环
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 前置判断
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	var (
		sum int
		carryBit int
	)
	resp := &ListNode{}
	node := resp

	for {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carryBit
		carryBit = sum / 10
		node.Val = sum % 10
		if l1 == nil && l2 == nil && carryBit == 0 {
			break
		}
		tempNode := &ListNode{}
		node.Next = tempNode
		node = node.Next
		sum = 0
	}
	return resp
}

/*
1
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

循环整个字符串

设置不重复字符开始位置和最大长度以及长度字段

map存储当前字符和字符下标，如果当前字符在map中存在，并且已存在的字符的下标大于当前不重复字符的开始位置，更新当前不重复字符的开始位置为已存在字符的下标
并且更新不重复字符的长度为当前字符的下标减去已存在字符的下标
如果不存在则，不重复字符长度++，
和最大长度进行比较
*/
func lengthOfLongestSubstring(s string) int {
	checkMap := make(map[rune]int, 0)
	start, maxLength, length := 0,0,0
	for i, s1 := range []rune(s) {
		if lastIndex, ok := checkMap[s1]; ok && start < lastIndex  {
			start = lastIndex + 1
			length = i - start + 1
		} else {
			length ++
		}
		checkMap[s1] = i
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
