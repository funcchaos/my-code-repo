package problems

import (
	"math/rand"
	"time"
)

type Solution struct {
	r    *rand.Rand
	nums []int
}

func Constructor398(nums []int) Solution {
	return Solution{r: rand.New(rand.NewSource(time.Now().UnixNano())), nums: nums}
}

func (this *Solution) Pick(target int) int {
	count ,index:= 0,0
	for i, v := range this.nums {
		if v == target {
			count++
			if this.r.Intn(count) == 0 {
				index= i
			}
		}
	}
	return index
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
