package problems

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
// 输入：n = 3
// 输出：[[1,2,3],[8,9,4],[7,6,5]]


func GgenerateMatrix(n int) [][]int {
	// 创建n x n 的切片
	ret := make([][]int,n)
	for i := range ret {
		ret[i]=make([]int,n)
	}

	// 正方形 外围一圈 四个点 加四条边,走一圈，四个点都会向内缩小，xy都-1
	// (0,0)     (0,n-1)	
	// (n-1,0)   (n-1,n-1)

	fillMatrix(0,0,n-1,n-1,1,ret)
	return ret
}

func fillMatrix(xs,ys,xe,ye,start int,sl [][]int) {
	// xy 起始位置大于终止位置时跳出
	if( xe<xs|| ye<ys){
		return 
	}
	if(xs==xe){
		sl[xe][ye]=start
		return 
	}
	v:=start
	// 填充一圈
	for i:=ys;i<ye;i++{
		sl[xs][i]=v
		v++
	}
	for i:=xs;i<xe;i++{
		sl[i][ye]=v
		v++
	}
	for i:=ye;i>ys;i--{
		sl[xe][i]=v
		v++
	}
	for i:=xe;i>xs;i--{
		sl[i][ys]=v
		v++
	}

	// 填充下一圈
	fillMatrix(xs+1,ys+1,xe-1,ye-1,v,sl)

}

