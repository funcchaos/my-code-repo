package problems

import (
	"math"
	"strings"
)

// 输入：s = " 3+5 / 2 "
// 输出：5

func Calculate2(s string) int {
	// 创建map来存放运算符优先级
	oplv := make(map[byte]int, 8)
	oplv['+'] = 1
	oplv['-'] = 1
	oplv['*'] = 2
	oplv['/'] = 2
	oplv['%'] = 2
	oplv['^'] = 3
	oplv['('] = 0
	oplv[')'] = 0
	// 去空格
	s = strings.ReplaceAll(s, " ", "")
	// 负号换成 0- 的形式
	s = strings.ReplaceAll(s, "(-", "(0-")

	lens := len(s)
	// 创建存放数字的栈
	nums := make([]int, 0, lens)
	// 创建存放运算符的栈
	ops := make([]byte, 0, lens)
	for i := 0; i < lens; i++ {
		c := s[i]
		if c == '(' {
			ops = append(ops, c)
		} else if c == ')' {
			// 计算到最近一个左括号为止
			for len(ops) > 0 {
				if ops[len(ops)-1] != '(' {
					calc(&nums, &ops)
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		} else {
			if isNumber(c) {
				u, j := 0, i
				// 将从 i 位置开始后面的连续数字整体取出，加入 nums
				for j < lens && isNumber(s[j]) {
					u = u*10 + int(s[j]-'0')
					j++
				}
				nums = append(nums, u)
				i = j - 1
			} else {
				// 有一个新操作要入栈时，先把栈内可以算的都算了
				// 只有满足「栈内运算符」比「当前运算符」优先级高/同等，才进行运算
				for len(ops) > 0 && ops[len(ops)-1] != '(' {
					prev := ops[len(ops)-1]
					if oplv[prev] >= oplv[c] {
						calc(&nums, &ops)
					} else {
						break
					}
				}
				ops = append(ops, c)
			}
		}
	}
	// 将剩余的计算完
	for len(ops) > 0 && ops[len(ops)-1] != '(' {
		calc(&nums, &ops)
	}

	return nums[len(nums)-1]
}

func isNumber(c byte) bool {
	return (c > 0x2f && c < 0x3a)
}

func calc(nums *[]int, ops *[]byte) {
	if len(*nums) < 2 {
		return
	}
	if len(*ops) <= 0 {
		return
	}

	a := (*nums)[len(*nums)-2]
	b := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-2]
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]
	ans := 0
	if op == '+' {
		ans = a + b
	} else if op == '-' {
		ans = a - b
	} else if op == '*' {
		ans = a * b
	} else if op == '/' {
		ans = a / b
	} else if op == '^' {
		ans = int(math.Pow(float64(a), float64(b)))
	} else if op == '%' {
		ans = a % b
	}
	*nums = append(*nums, ans)
}
