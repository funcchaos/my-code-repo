package problems

// import "fmt"

// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

func SspiralOrder(matrix [][]int) []int {
	lh:=len(matrix)
	lw:=len(matrix[0])
	// x,y:=0,lw
	ret:=make([]int,0,lw*lh)
	xs,xe:=0,lh-1
	ys,ye:=0,lw-1

	for x,y:=0,0;x<=xe &&y <=ye;{

		// ret = append(ret, matrix[xs][ys])

		for i := ys; i <= ye; i++ {
			ret = append(ret, matrix[xs][i])
		}
		if(xs>=xe){ 
			break
		}
		xs++
		// ret = append(ret, matrix[xs][ye])
		for i := xs; i <= xe; i++ {
			ret = append(ret, matrix[i][ye])
		}
		if(ys>=ye){   
			break
		}
		ye--
		// ret = append(ret, matrix[xe][ye])
		for i := ye; i >=ys; i-- {
			ret = append(ret,matrix[xe][i])
		}
		if(xe==xs){
			break
		}
		xe--
		// ret = append(ret, matrix[xe][ys])
		for i := xe; i >=xs; i-- {
			ret = append(ret,matrix[i][ys])
		}
		if(ys==ye){
			break
		}
		ys++

	}

	// for x,y :=0, 0; x <lh-x && y<lw-y ;  {
	// 	for i:=y;i<lw-y;i++{
	// 		fmt.Println("x,y---",x,"--",i)
	// 		ret = append(ret, matrix[x][i])
	// 	}
	// 	for i:=x+1;i<lh-x;i++{
	// 		fmt.Println("x,y---",i,"--",lw-y-1)
	// 		ret=append(ret, matrix[i][lw-y-1])
	// 	}
	// 	for i:=lw-y-2;i>y;i--{
	// 		fmt.Println("x,y---",lh-x-1,"--",i)
	// 		ret = append(ret, matrix[lh-x-1][i])
	// 	}
	// 	for i:=lh-x-2;i>x;i--{
	// 		fmt.Println("x,y---",i,"--",y)
	// 		ret=append(ret,matrix[i][y])
	// 	}
	// 	x++
	// 	y++
	// }
	return ret
}


	// 第一圈
	// x  0->0->(lh-1)->(lh-1)->1
	// y  0->(lw-1)->(lw-1)->0->0

	// 第二圈
	// x 范围 1->lh-2
	// y 范围 1->lw-2
	