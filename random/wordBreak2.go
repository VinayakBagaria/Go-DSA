// https://leetcode.com/problems/word-break-ii/description/
package random

import "strings"

func wordBreak2(s string, wordDict []string) []string {
	answers := []string{}

	var dfs func(string, string)
	dfs = func(shortened, tillNow string) {
		if len(shortened) == 0 {
			answers = append(answers, strings.Trim(tillNow, " "))
			return
		}

		for _, w := range wordDict {
			end := len(w)

			if end <= len(shortened) && shortened[0:end] == w {
				dfs(shortened[end:], tillNow + " " + w)
			}
		}
	}

	dfs(s, "")
	return answers
}