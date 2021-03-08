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

	s := "aabbaccddcca"
	// s := "ababbbabbabaaasbbabskbbasaaabbababbbsndbandjaaaabb"
	mylib.MinCut(s)

}
