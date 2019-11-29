/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 *
 * https://leetcode.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (32.46%)
 * Likes:    2495
 * Dislikes: 64
 * Total Accepted:    206.7K
 * Total Submissions: 631.6K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * Given n non-negative integers representing the histogram's bar height where
 * the width of each bar is 1, find the area of largest rectangle in the
 * histogram.
 *
 *
 *
 *
 * Above is a histogram where width of each bar is 1, given height =
 * [2,1,5,6,2,3].
 *
 *
 *
 *
 * The largest rectangle is shown in the shaded area, which has area = 10
 * unit.
 *
 *
 *
 * Example:
 *
 *
 * Input: [2,1,5,6,2,3]
 * Output: 10
 *
 *
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	heights = append(heights, 0)
	s := &Stack{}
	max := 0
	for i := 0; i < len(heights); i++ {
		for !s.Empty() && heights[i] < heights[s.Top()] {
			top := s.Top()
			s.Pop()
			if s.Empty() {
				area := heights[top] * i
				if max < area {
					max = area
				}
			} else {
				area := heights[top] * (i - s.Top() - 1)
				if max < area {
					max = area
				}
			}
		}
		s.Push(i)
	}

	return max
}

type Stack struct {
	num []int
}

func (s *Stack) Pop() {
	if len(s.num) == 0 {
		return
	}
	s.num = s.num[:len(s.num)-1]
	return
}

func (s *Stack) Push(d int) {
	s.num = append(s.num, d)
	return
}

func (s *Stack) Top() int {
	if len(s.num) == 0 {
		return -1
	}
	return s.num[len(s.num)-1]
}

func (s *Stack) Empty() bool {
	return len(s.num) == 0
}

// @lc code=end

