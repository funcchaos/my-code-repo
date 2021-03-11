package problems

// 输入：s = "(1+(4+5+2)-3)+(6+8)"
// 输出：23

// 将正常的中缀表达式转换为后缀表达式，也就是逆波兰表达式
// 一个栈用来存放数据，一个栈用来存放操作符

// []byte("0123456789+-*/() ")
// []uint8{
//   0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x2b, 0x2d, 0x2a, 0x2f, 0x28, 0x29, 0x20,
// }

// operators := make([]rune,0,len(s))
// nummap := make(map[byte]int, 10)
// nummap[0x30] = 0
// nummap[0x31] = 1
// nummap[0x32] = 2
// nummap[0x33] = 3
// nummap[0x34] = 4
// nummap[0x35] = 5
// nummap[0x36] = 6
// nummap[0x37] = 7
// nummap[0x38] = 8
// nummap[0x39] = 9
// '-' 0x2d
// oplv["+"] = 1
// oplv["-"] = 1
// oplv["*"] = 2
// oplv["/"] = 2
// oplv["("] = 0
// oplv[")"] = 0
// oplv["%"] = 2
// 减法可以按 + 负数处理，或者0-处理


func calculate(s string) int {
	stack := make([]int, 0)
	res := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			num := 0
			for ; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					break
				}
				num = num*10 + int(s[i]-'0')
			}
			res = res + sign*num
			i--
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, res, sign)
			res, sign = 0, 1
		case ')':
			sign := stack[len(stack)-1]
			num := stack[len(stack)-2]
			res = num + sign*res
			stack = stack[:len(stack)-2]
		}
	}

	return res
}
