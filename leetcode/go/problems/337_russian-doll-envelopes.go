package problems

import "sort"

//MaxEnvelopes 找到最多的俄罗斯套娃
func MaxEnvelopes(envelopes [][]int) int {
	// 先按w升序,相同w的按h降序
	sort.Slice(envelopes, func(a, b int) bool {
		return envelopes[a][0] < envelopes[b][0] ||
			(envelopes[a][0] == envelopes[b][0] &&
				envelopes[a][1] > envelopes[b][1])
	})

	max := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		if len(max) == 0 || envelopes[i][1] > max[len(max)-1] {
      // 根据上面的排序，后面h大的就是符合套娃
			max = append(max, envelopes[i][1])
		} else {
      // 如果后面h小于当前最大的h,从前面记录中找到当前h合适位置
			for j := 0; j < len(max); j++ {
				if envelopes[i][1] <= max[j] {
					max[j] = envelopes[i][1]
					break
				}
			}
		}
	}
	return len(max)
}
