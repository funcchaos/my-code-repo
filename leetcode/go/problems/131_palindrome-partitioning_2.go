package problems

import "fmt"

//Partition1 回文分割
func Partition1(s string) [][]string {
    if len(s) == 0 {
        return [][]string{}
    }
    res := [][]string{}
    backtracking(s,0,[]string{},&res)
    return res
}

func backtracking(s string,index int,list []string,res *[][]string){
    if index == len(s) {
        tmp := make([]string,len(list))
        copy(tmp,list)
        *res = append(*res,tmp)
        return
    }

    for i := index;i < len(s);i++ {
        if helper(s[index:i+1]) {
            list = append(list,s[index:i+1])
        }else{
            continue
        }
		fmt.Println("he--l--",list)
		fmt.Println("i1--,index1---",i,index)
        backtracking(s,i+1,list,res)
        list = list[:len(list)-1]
		fmt.Println("list-:len-1---",list)
		fmt.Println("i2--,index2---",i,index)
    }
}

// 判断是否回文
func helper(s string)bool{
    if len(s) == 1 {
        return true
    }
    i,j := 0,len(s) - 1
    for i < j {
        if s[i] != s[j]{
            return false
        }
        i++
        j--
    }
    return true
}