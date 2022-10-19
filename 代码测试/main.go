package main

import (
	"fmt"
	"sort"
)

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

两数之和
 */
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
	fmt.Println(generateParenthesis(3))

}
/*
17，在排序数组中查找元素的第一个和最后一个位置

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1,-1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0,0}
		}
		return []int{-1,-1}
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			for {

			}
		}
	}

	return
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
