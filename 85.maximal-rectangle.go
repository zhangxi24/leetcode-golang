/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 *
 * https://leetcode.com/problems/maximal-rectangle/description/
 *
 * algorithms
 * Hard (34.84%)
 * Likes:    1880
 * Dislikes: 55
 * Total Accepted:    142.4K
 * Total Submissions: 405.6K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle
 * containing only 1's and return its area.
 *
 * Example:
 *
 *
 * Input:
 * [
 * ⁠ ["1","0","1","0","0"],
 * ⁠ ["1","0","1","1","1"],
 * ⁠ ["1","1","1","1","1"],
 * ⁠ ["1","0","0","1","0"]
 * ]
 * Output: 6
 *
 *
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	area := 0
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		tmp := largestRectangleArea(heights)
		if area < tmp {
			area = tmp
		}
	}
	return area
}

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

