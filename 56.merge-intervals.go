import (
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (36.92%)
 * Likes:    2746
 * Dislikes: 218
 * Total Accepted:    438.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 *
 * Example 1:
 *
 *
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 *
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 *
 */

// @lc code=start
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

type Interval [][]int

func (s Interval) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s Interval) Swap(i, j int) {
	for n := 0; n < len(s[i]); n++ {
		s[i][n], s[j][n] = s[j][n], s[i][n]
	}
}

func (s Interval) Len() int {
	return len(s)
}

// @lc code=end

