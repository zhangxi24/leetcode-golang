/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (32.66%)
 * Likes:    856
 * Dislikes: 362
 * Total Accepted:    159.6K
 * Total Submissions: 484.6K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 *
 * Example:
 *
 *
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 *
 *
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := []string{}
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				for d := 1; d <= 3; d++ {
					if a+b+c+d == len(s) {
						if check(s[:a]) && check(s[a:a+b]) && check(s[a+b:a+b+c]) && check(s[a+b+c:]) {
							str := s[:a] + "." + s[a:a+b] + "." + s[a+b:a+b+c] + "." + s[a+b+c:]
							res = append(res, str)
						}
					}
				}
			}
		}
	}

	return res
}

func check(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	if len(s) == 3 && s > "255" {
		return false
	}
	return true
}

// @lc code=end

