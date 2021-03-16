package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	// "github.com/pkg/profile"

	mylib "leetcode/problems"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// f, _ := os.Create("./cpu.prof")
	// fm, _ := os.Create("./mem.prof")
	// pprof.StartCPUProfile(f)
	// defer profile.Start().Stop()

	fmt.Println(time.Now())

	// arr := []int{1, 2, 1}
	// sp := mylib.NextGreaterElements(arr)
	// fmt.Println(sp)

	// s := "aabbaccddcca"
	// s := "ababbbabbabaaasbbabskbbasaaabbababbbsndbandjaaaabb"
	// mylib.MinCut(s)
	// s := "(1+(4+5+2)-3)+(6+8)"
	// fmt.Println(mylib.Calculate2(s))

	// s331:="9,3,4,#,#,1,#,#,2,#,6,#,#"
	// s331 := "#"
	// fmt.Println(mylib.IsValidSerialization(s331))

	// m1:=[]int{1, 2, 3, 4}
	// m2:=[]int{5, 6, 7, 8}
	// m3:=[]int{9, 10, 11, 12}
	// matrix54 := [][]int{m1,m2 ,m3 }
	// fmt.Println(mylib.SspiralOrder(matrix54))

	matrix59:=mylib.GgenerateMatrix(5)
	fmt.Println(matrix59)

}
