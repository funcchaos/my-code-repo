package problems

// import "fmt"

//MinCut 找到最少回文分割次数
func MinCut(s string) int {
	l := len(s)
	list := make([]int, l+1)   // list[i]代表前i个字符需要划几次，特别地，list[0]=-1 (这个解法里面的这个数组就是精髓了)
	for i := 0; i < l+1; i++ { // 初始化[-1, 0, 1, 2, 3...]，最差情况就是每个字母独立分割
		list[i] = i - 1
	}
	for i := 0; i < l; i++ { // 以每个字符为中心找最长回文子串
		list[i+1] = min(list[i+1], list[i]+1) // 初始化，最坏情况下就比左边的多划一次
		if i == l-1 {                         // 最后一个了没必要找了
			break
		}
		// 先找偶数个的
		start := i
		end := i + 1

		findPalindrome(s, list, &start, &end)
		// 再找奇数个的
		start = i - 1
		end = i + 1
		if start < 0 {
			continue
		}
		findPalindrome(s, list, &start, &end)

		// 如果整个串都是回文串，那么就中断
		if list[l] == 0 {
			return 0
		}
	}
	return list[l]
}

// 搜索回文
func findPalindrome(s string, list []int, start, end *int) {
	l := len(s)
	for s[*start] == s[*end] {
		list[*end+1] = min(list[*end+1], list[*start]+1)
		// fmt.Println("---", list)
		if *end == l-1 || *start == 0 {
			break
		}
		*start--
		*end++
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
