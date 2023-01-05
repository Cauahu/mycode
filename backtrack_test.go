package mycode

import (
	"sort"
	"strings"
	"testing"
)

/*140
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
说明：
分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：
输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
*/
func wordBreak(s string, wordDict []string) []string {
	wordmap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordmap[word] = struct{}{}
	}

	res := make([]string, 0)
	n := len(s)
	var backtrack func(i int, sentence []string)
	backtrack = func(i int, sentence []string) {
		if i == n {
			res = append(res, strings.Join(sentence, " "))
			return
		}

		for j := i + 1; j <= n; j++ {
			if _, ok := wordmap[s[i:j]]; ok {
				sentence = append(sentence, s[i:j])
				backtrack(j, sentence)
				sentence = sentence[:len(sentence)-1]
			}
		}
	}
	backtrack(0, []string{})

	return res
}

func Test_wordBreak(t *testing.T) {
	res := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	for _, s := range res {
		t.Log(s)
	}
}

/*17
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
func letterCombinations(digits string) []string {
	digitmap := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	n := len(digits)
	res := make([]string, 0)

	var backtrack func(i int, s []byte)
	backtrack = func(i int, s []byte) {
		if i == n {
			if len(s) != 0 {
				res = append(res, string(s))
			}
			return
		}

		ds := digitmap[digits[i]]
		for _, d := range ds {
			s = append(s, d)
			backtrack(i+1, s)
			s = s[:len(s)-1]
		}
	}

	backtrack(0, []byte{})

	return res
}

func Test_letterCombinations(t *testing.T) {
	t.Log(letterCombinations("23"))
	t.Log(letterCombinations(""))
}

/*22
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
*/
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var backtrack func(l, r int, s []byte)
	backtrack = func(l, r int, s []byte) {
		if l == n && r == n {
			if len(s) != 0 {
				res = append(res, string(s))
			}
			return
		}

		if l > n || r > n {
			return
		}

		if l < n {
			s = append(s, '(')
			backtrack(l+1, r, s)
			s = s[:len(s)-1]
		}

		if l > r {
			s = append(s, ')')
			backtrack(l, r+1, s)
			s = s[:len(s)-1]
		}
	}
	backtrack(0, 0, []byte{})

	return res
}

func Test_generateParenthesis(t *testing.T) {
	t.Log(generateParenthesis(3))
}

/*698
给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
*/
func canPartitionKSubsets(nums []int, k int) bool {
	target := 0
	used := make(map[int]int, 0)
	for _, n := range nums {
		target += n
		used[n]++
	}
	if target%k != 0 {
		return false
	}

	target = target / k
	n := len(nums)
	var backtrack func(idx, sum int, s []int) bool
	backtrack = func(idx, sum int, s []int) bool {
		if idx == k {
			return true
		}

		for i := 0; i < n; i++ {
			num := nums[i]
			if used[num] != 0 {
				used[num]--
				s = append(s, num)
				if sum+num == target {
					return backtrack(idx+1, sum+num, s)
				}
				s = s[:len(s)-1]
				if backtrack(idx+1, sum, s) {
					return true
				}

			}
		}
		return false
	}
	return backtrack(0, 0, []int{})
}

/*77
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
*/
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(idx int, s []int)
	backtrack = func(idx int, s []int) {
		if len(s) == k {
			tmp := make([]int, k)
			copy(tmp, s)
			res = append(res, tmp)
			return
		}

		for i := idx; i <= n; i++ {
			s = append(s, i)
			backtrack(i+1, s)
			s = s[:len(s)-1]
		}
	}

	backtrack(1, []int{})
	return res
}

func Test_combine(t *testing.T) {
	t.Log(combine(4, 2))
}

/*78
给定一组 不含重复元素 的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
*/
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	var backtrack func(idx int, s []int)
	backtrack = func(idx int, s []int) {
		if idx <= n {
			tmp := make([]int, len(s))
			copy(tmp, s)
			res = append(res, tmp)
		}

		for i := idx; i < n; i++ {
			s = append(s, nums[i])
			backtrack(i+1, s)
			s = s[:len(s)-1]
		}
	}

	backtrack(0, []int{})
	return res
}

func Test_subsets(t *testing.T) {
	t.Log(subsets([]int{1, 2, 3}))
}

/*859
给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。
*/
func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		dup := make(map[byte]struct{}, len(s))
		for _, b := range []byte(s) {
			if _, ok := dup[b]; ok {
				return true
			}
		}
	}

	loc1, loc2, count, n := 0, 0, 0, len(s)
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			count++
			if loc1 == 0 {
				loc1 = i
			} else {
				loc2 = i
			}
		}

		if count > 2 {
			return false
		}
	}

	if s[loc1] == goal[loc2] && s[loc2] == goal[loc1] {
		return true
	}

	return false
}

func Test_buddyStrings(t *testing.T) {
	t.Log(buddyStrings("aaaaaaabc", "aaaaaaacb"))
}

/*368
给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i], answer[j]) 都应当满足：
answer[i] % answer[j] == 0 ，或
answer[j] % answer[i] == 0
如果存在多个有效解子集，返回其中任何一个均可。
示例 1：
输入：nums = [1,2,3]
输出：[1,2]
解释：[1,3] 也会被视为正确答案。
输入：nums = [1,2,4,8]
输出：[1,2,4,8]
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	isvalid := func(a, b int) bool {
		if b%a == 0 {
			return true
		}
		return false
	}

	res := make([]int, 0, len(nums))
	n := len(nums)
	var backtrack func(idx int, s []int)
	backtrack = func(idx int, s []int) {
		if len(s) > len(res) {
			res = res[:0]
			res = append(res, s...)
		}

		for i := idx; i < n; i++ {
			if len(s) == 0 || isvalid(s[len(s)-1], nums[i]) {
				s = append(s, nums[i])
				backtrack(i+1, s)
				s = s[:len(s)-1]
			}
		}
	}

	backtrack(0, []int{})
	return res
}

func largestDivisibleSubset_dp(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	trans := make([]int, len(nums))
	n := len(nums)
	for i := 0; i < n; i++ {
		dp[i] = 1
		trans[i] = -1
	}

	maxidx := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				trans[i] = j
			}

			if dp[i] > dp[maxidx] {
				maxidx = i
			}
		}
	}

	res := make([]int, 0)
	for i := maxidx; i >= 0; {
		res = append(res, nums[i])
		i = trans[i]
	}

	return res
}

func Test_largestDivisibleSubset(t *testing.T) {
	//t.Log(largestDivisibleSubset([]int{1, 2, 3}))
	//t.Log(largestDivisibleSubset([]int{1, 2, 4, 8}))
	//t.Log(largestDivisibleSubset([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))

	t.Log(largestDivisibleSubset_dp([]int{1, 2, 3}))
	t.Log(largestDivisibleSubset_dp([]int{1, 2, 4, 8}))
	t.Log(largestDivisibleSubset_dp([]int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}))
}
