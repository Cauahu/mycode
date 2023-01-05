package mycode

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*45
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
*/
func jump(nums []int) int {
	end, maxstep, step := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		maxstep = max(maxstep, i+nums[i])
		if i == end {
			end = maxstep
			step++
		}
	}
	return step
}

/*53
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}

	res := nums[0]
	for i := 1; i < len(dp); i++ {
		res = max(res, dp[i])
	}

	return res
}

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(nums))

	nums = []int{5, 4, -1, 7, 8}
	t.Log(maxSubArray(nums))
}

/*96
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/
func numTrees(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i] += f[j-1] * f[i-j]
		}
	}

	return f[n]
}

/*95
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var traverse func(beg, end int) []*TreeNode
	traverse = func(beg, end int) []*TreeNode {
		if beg > end {
			return []*TreeNode{nil}
		}

		res := make([]*TreeNode, 0)
		for i := beg; i <= end; i++ {
			lefts := traverse(beg, i-1)
			rights := traverse(i+1, end)
			for _, left := range lefts {
				for _, right := range rights {
					cur := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}
					res = append(res, cur)
				}
			}

		}
		return res
	}

	return traverse(1, n)
}

/*120
给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}

	return res
}

func Test_minimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	t.Log(minimumTotal(triangle))
}

/*131
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。
s = "aab", [["a","a","b"],["aa","b"]]
*/
func partition(s string) [][]string {
	return nil
}

func Test_timercost(t *testing.T) {
	fi, err := os.Open("./triggerTime")
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close()

	timemap := make(map[string][]string, 1300)
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		sa := strings.Split(string(a), " ")
		if len(sa) != 2 {
			t.Fatal("2")
		}

		timemap[sa[1]] = []string{sa[0]}
	}

	fi1, err1 := os.Open("./completionTime")
	if err1 != nil {
		t.Fatal(err1)
	}
	defer fi1.Close()

	count := 0
	br1 := bufio.NewReader(fi1)
	for {
		a, _, c := br1.ReadLine()
		if c == io.EOF {
			break
		}
		sa := strings.Split(string(a), " ")
		if len(sa) != 2 {
			t.Fatal("2")
		}

		if ti, ok := timemap[sa[1]]; ok {
			ti = append(ti, sa[0])
			timemap[sa[1]] = ti
		} else {
			count++
		}
	}

	costsum, i := 0, 0
	costs := make([]time.Duration, 0, 1300)
	max, min := math.MinInt, math.MaxInt
	for _, sa := range timemap {
		if len(sa) != 2 {
			continue
		}
		beg, _ := strconv.Atoi(sa[0])
		end, _ := strconv.Atoi(sa[1])
		cost := end - beg
		costsum += cost
		costs = append(costs, time.Duration(cost)*time.Millisecond)
		if max < cost {
			max = cost
		}

		if min > cost {
			min = cost
		}
		i++
	}

	t.Logf("avg=%v", time.Duration(costsum/i)*time.Millisecond)
	t.Logf("mean=%v", costs[len(costs)/2])
	t.Logf("max=%v", time.Duration(max)*time.Millisecond)
	t.Logf("min=%v", time.Duration(min)*time.Millisecond)
}

func Test1(t *testing.T) {
	t.Run("周三", func(t *testing.T) {
		for i := 1; i < 31; i++ {
			t.Log(getTitleDesc(i, 3))
		}
	})

	t.Run("周日", func(t *testing.T) {
		for i := 1; i < 31; i++ {
			t.Log(getTitleDesc(i, 0))
		}
	})

}

var WeekDayMap = map[int]string{
	1: "周一",
	2: "周二",
	3: "周三",
	4: "周四",
	5: "周五",
	6: "周六",
	0: "周日",
}

func getTitleDesc(idx, weekday int) string {
	if idx == 0 {
		return "今日 "
	}

	if idx == 1 {
		return "明日 "
	}

	// 重置周日号码
	if weekday == 0 {
		weekday = 7
	}

	// 超过2周
	if idx+weekday > 14 {
		return ""
	}

	ret := ""
	if idx+weekday < 8 {
		ret = "本"
	} else {
		ret = "下"
	}

	return ret + WeekDayMap[(idx+weekday)%7] + " "
}
