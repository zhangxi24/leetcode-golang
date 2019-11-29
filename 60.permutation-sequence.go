import "fmt"

/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (34.34%)
 * Likes:    1046
 * Dislikes: 269
 * Total Accepted:    155.1K
 * Total Submissions: 449.6K
 * Testcase Example:  '3\n3'
 *
 * The set [1,2,3,...,n] contains a total of n! unique permutations.
 *
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * Given n and k, return the k^th permutation sequence.
 *
 * Note:
 *
 *
 * Given n will be between 1 and 9 inclusive.
 * Given k will be between 1 and n! inclusive.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3, k = 3
 * Output: "213"
 *
 *
 * Example 2:
 *
 *
 * Input: n = 4, k = 9
 * Output: "2314"
 *
 *
 */

// @lc code=start
// 获取递归数列
func fac(n int) map[int]int {
	m := map[int]int{}
	m[0] = 1
	for i := 1; i < n; i++ {
		m[i] = i * m[i-1]
	}
	return m
}

// 获取下一个位置pos的数字
func getNext(pos int, used *[]int) string {
	fmt.Println(pos, *used)
	i := 0
	for i <= pos {
		if (*used)[i] == 1 {
			pos++
		}
		i++
	}
	(*used)[pos] = 1
	return string('1' + pos)
}

// 10-2341
// k:10, 4, 2, 1
// pos:1,1,1,0
// (10-1)/6, (4-1)/2, (2-1)/1, (1-1)/1
func getPermutation(n int, k int) string {
	str := ""
	k--
	dict := fac(n)
	used := make([]int, n)
	for i := 0; i < n; i++ {
		pos := k / dict[n-1-i] // 第i位放剩余数字集合中的第pos个数字
		k = k % dict[n-1-i]    // 下一次第k个排列
		str += getNext(pos, &used)
	}
	return str
}

// @lc code=end

