package main

// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75
func main() {
	res := mergeAlternately("abcd", "")
	println(res)
}

func mergeAlternately(word1 string, word2 string) string {
	finalStr := ""
	j := 0
	for i := range word1 {
		finalStr += string(word1[i])
		if len(word2) > i {
			finalStr += string(word2[i])
		}
		j++
	}

	for ; j < len(word2); j++ {
		finalStr += string(word2[j])
	}

	return finalStr
}
