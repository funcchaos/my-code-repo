package problems

// import "fmt"

// 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

// 返回 s 所有可能的分割方案。

// 示例:

// 输入: "aab"
// 输出:
// [
//   ["aa","b"],
//   ["a","a","b"]
// ]
// aabbaccddcca
// [a,a,b,b,a,c,c,d,d,c,c,a]
// [                     ,a]
// [                   ,c,a]
// 1-[                 ,c,c,a]
// 1-[                  ,cc,a]
// [                ,d,1]
// 2-[             ,dd,1]
// 2-[             ,d,d,1]
// 3-[              ,c,2]
// 3-[           ,cddc,c,a]
// 4-[,c,c,d,d,c,c,a]
// 4-[,cc,2]
// 4-[,ccddcc,a]

// 判断：
// 从左开始
// i=0,i<len(s),i++
// a0 从包含这个元素作为一次大循环            0 1 2 3 4 5 6 7 8    3~8 8~13 或者 9~14  2~5 5~8 6~9
// 判断是否回文， 即 索引 n->m的 元素 倒序是 m->2*m-n 或者 m+1->m+1+m-n
// a0,[]然后对后半段重复这个判断
// a0a1,[]然后也对后半段进行这个操作
// 最上层循环 a aa aab aabb这样下去
// a的时候 再去切abb
// aa           bb
// aab 不是回文不用管他
// aabb         无了

// 所以从a开始，循环就可以找到 a开头的所有回文串
// 比如a,a作为前面的base数组，用来给后续拼接完整答案 所以入参应该就是 a ，abb， func(base []string,tohandle []string)
// abb 同样像之前的处理，找到 abb中以a开头的所有回文串 base参数，就应该是 a+a了
// 每条子线路的终止条件就是 tohandle 为空的时候

//Partition 分割回文串
func Partition(s string) [][]string {
	res := [][]string{}
	// 先获取基础数组
	base := []byte(s)
	// base := make([]byte, len(s))
	// for _, v := range s {
	// 	base = append(base, v)
	// }
	// 从左开始循环
	partRecur(base, []string{}, &res)

	return res
}

// 查找回文串
func partRecur(s []byte, t []string, res *[][]string) {
	// 从s[0]开始
	pmap := getPalindrome(s)
	//fmt.Println("s--", string(s))
	//fmt.Println("pm--", pmap)
	// 遍历kyes
	// for k, v := range pmap {
	// 	// tmps := make([]string,0)
	// 	tmps := append(t, k)
	// 	//fmt.Println("tmps---", tmps)
	// 	if len(v) == 0 {
	// 		*res = append(*res, tmps)
	// 	} else {
	// 		partRecur(v, tmps, res)
	// 	}
	// }
	// 答案对顺序有要求的，回文短的排前面，这里对key进行排序,其实前面可以直接使用两个数组，能快几十ms
	karr := make([]string, 0,len(pmap))
	for k := range pmap {
		if len(karr) == 0 {
			karr = append(karr, k)
		} else {
			if k > karr[len(karr)-1] {
				karr = append(karr, k)
			} else if k < karr[0] {
				karr = append([]string{k}, karr...)
			} else {
				for i := 0; i < len(karr); i++ {
					if k > karr[i] && k < karr[i+1] {
						end:= append([]string{},karr[i+1:]...)
						karr = append(append(karr[:i+1], k),end...)
						break
					}
				}
			}
		}
	}
	for _, v := range karr {
		tmps := append(append([]string{},t...), v)
		if len(pmap[v]) == 0 {
			*res = append(*res, tmps)
		} else {
			partRecur(pmap[v], tmps, res)
		}
	}

}

// 获取切片中第一个元素为首的所有回文串
func getPalindrome(s []byte) map[string][]byte {
	pmap := make(map[string][]byte, 0)
	if len(s) > 0 {
		// pmap[string(s[0])] = s[1:]
		for i := 0; i < len(s)/2+1; i++ {
			if isMirror(s[0:i+1], s[i:i*2+1]) {
				pmap[string(s[0:i*2+1])] = s[i*2+1:]
			}
			if isMirror(s[0:i+1], s[i+1:i*2+2]) {
				pmap[string(s[0:i*2+2])] = s[i*2+2:]
			}
		}
	}
	return pmap
}

// 判断两个切片是否是镜像
func isMirror(a, b []byte) bool {
	isM := true
	for i, j := 0, len(a)-1; j >= 0; i, j = i+1, j-1 {
		if a[i] != b[j] {
			isM = false
			break
		}
	}
	return isM
}
