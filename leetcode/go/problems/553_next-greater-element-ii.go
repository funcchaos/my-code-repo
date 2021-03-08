package problems

// 给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），
// 输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
// 这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
// 如果不存在，则输出 -1。
// 输入: [1,2,1]
// 输出: [2,-1,2]

// nt做法尝试
// 先定义一个等长数组，全是-1
// 定义一个当前下一个较大数，记录数v和索引位置i
// 直接遍历索引小于i的数，如果小于v，就遍历当前索引到i之间的数，查找是否有符合条件的
// 原本的数组 			so
// 实际需要遍历的数组	 sp=so+so[:len(so)-1]
// 结果数组				sr len(sr)==len(so)

type tmps struct {
	v  int
	i  int
	ri int
}

//NextGreaterElements 直接无脑循环也能过
func NextGreaterElements(nums []int) []int {
	lens := len(nums)
	sr := make([]int, lens)
	for i := range sr {
		sr[i] = -1
	}
	if lens > 0 {
		sp := append(nums, nums[:lens-1]...)
		// 找到第一个数的下一个较大数
		// a1:= nums[0]
		for i := 0; i < lens; i++ {
			// 确定当前数查找范围
			spi := sp[i+1 : lens+i]
			for j := 0; j < lens-1; j++ {
				if spi[j] > nums[i] {
					sr[i] = spi[j]
					break
				}
			}
		}
	}
	return sr
}

//NextGreaterElements1 用栈来处理
func NextGreaterElements1(nums []int) []int {
	lens := len(nums)
	sr := make([]int, lens)
	for i := range sr {
		sr[i] = -1
	}

	// 使用栈来存放还没有找到大于的数的元素索引
	stack := make([]int, 0, lens)
	for i := 0; i < lens*2-1; i++ {
		ri := i % lens
		// 循环判断下一位值是否大于栈顶元素,因为循环数组，所以索引取模
		// for len(stack) > 0 &&   nums[stack[len(stack)-1]] < nums[ri] {
		// 	last := stack[len(stack)-1]
		// 	sr[last] = nums[ri]
		// 	stack = stack[:len(stack)-1]
		// }
		for len(stack) > 0 {
			// 符合条件的就给sr对应位置(栈顶对应的值为索引)赋值，并且从栈顶删除元素
			// 比如 3，2， 1， 4 ，栈里面应该有 3,2,1 直到4的时候
			last := stack[len(stack)-1]
			if nums[last] < nums[ri] {
				sr[last] = nums[ri]
				stack = stack[:len(stack)-1]
			} else{
				break
			}
		}
		// 往栈里面塞索引
		if i < lens {
			stack = append(stack, ri)
		}
	}

	return sr
}
