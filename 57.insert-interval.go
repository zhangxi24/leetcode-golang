import "sort"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (31.93%)
 * Likes:    1089
 * Dislikes: 136
 * Total Accepted:    203.5K
 * Total Submissions: 635.1K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * Given a set of non-overlapping intervals, insert a new interval into the
 * intervals (merge if necessary).
 *
 * You may assume that the intervals were initially sorted according to their
 * start times.
 *
 * Example 1:
 *
 *
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 *
 *
 * Example 2:
 *
 *
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */

// @lc code=start
type Interval [][]int

func (s Interval) Less(i, j int) bool {
	return s[i][0] <= s[j][0]
}

func (s Interval) Swap(i, j int) {
	for m := range s[i] {
		s[i][m], s[j][m] = s[j][m], s[i][m]
	}
}

func (s Interval) Len() int {
	return len(s)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	sort.Sort(Interval(intervals))
	// 找到受影响的位置
	start := 0
	for start < len(intervals) {
		if intervals[start][1] < newInterval[0] {
			start++
		} else {
			break
		}
	}
	end := start
	for end < len(intervals) {
		if intervals[end][0] <= newInterval[1] {
			end++
		} else {
			break
		}
	}

	if start == end && start == 0 { // 插入最前
		res = append(res, add2Interval(newInterval, intervals[0])...)
		for i := 1; i < len(intervals); i++ {
			res = append(res, intervals[i])
		}
	} else if start == end && start == len(intervals) { // 最后
		for i := 0; i < start; i++ {
			res = append(res, intervals[i])
		}
		res = append(res, newInterval)
	} else {
		if start == end {
			end++
		}
		for i := 0; i < start; i++ {
			res = append(res, intervals[i])
		}
		res = append(res, add2Interval(newInterval, intervals[start])...)
		for i := start + 1; i < end; i++ {
			addInterval(&res, intervals[i])
		}
		for i := end; i < len(intervals); i++ {
			res = append(res, intervals[i])
		}
	}

	return res
}

// 必须将new放在前边
func add2Interval(a, b []int) [][]int {
	r := [][]int{}
	if a[1] < b[0] {
		r = append(r, a)
		r = append(r, b)
	} else {
		min, max := a[0], a[1]
		if b[0] < min {
			min = b[0]
		}
		if b[1] > max {
			max = b[1]
		}
		r = append(r, []int{min, max})
	}
	return r
}

func addInterval(r *[][]int, add []int) {
	l := len(*r)
	if l == 0 {
		*r = append(*r, add)
	} else {
		if add[0] > (*r)[l-1][1] {
			*r = append(*r, add)
		}
		if add[1] > (*r)[l-1][1] {
			(*r)[l-1][1] = add[1]
		}
	}
}

/*
// 利用上一题代码
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	if len(intervals) == 0 {
		return res
	}

	sort.Sort(Interval(intervals))
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else if res[len(res)-1][1] >= intervals[i][0] && res[len(res)-1][1] < intervals[i][1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}
	return res
}
*/

// @lc code=end

