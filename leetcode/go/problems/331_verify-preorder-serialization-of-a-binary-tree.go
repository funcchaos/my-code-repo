package problems

import "fmt"

// "9,3,4,#,#,1,#,#,2,#,6,#,#"

func IsValidSerialization(preorder string) bool {
	l:=len(preorder)
	// 可用的挂载数量
	mountable :=0
	// 已挂载数量
	mounted:=0
	// 用来跳过多位数字
	flag:=false
	for i := 0; i < l; i++ {
		b:=preorder[i]
		if(i==0){
			if(b!='#'){
				mountable+=2
			}
			continue
		}

		fmt.Println("b---------------------------------",string(b))
		fmt.Println("ma--",mountable)
		fmt.Println("md--",mounted)

		if i<l-1 && mounted>=mountable{
			return false
		}
		if(b==','){
			flag=false
			continue
		}else if(b=='#'){
			flag=false
			mounted++
		}else{
			if(!flag){
				mountable+=2
				mounted++
				flag=true
			}else{
				continue
			}
		}
		
	}
	return mountable==mounted
}