package main;

import (
	"fmt"
	"os"
	"strconv"
	wp "github.com/colinsmith/waterpouring"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: capacity... goal")
		return
	}
	caps := make([]int, len(os.Args)-2)
	var err error
	for i, a := range os.Args[1:len(os.Args)-1] {
		if caps[i], err = strconv.Atoi(a); err != nil {
			fmt.Fprintln(os.Stderr, "bad argument", a)
			return
		}
	}
	var goal int
	if goal, err = strconv.Atoi(os.Args[len(os.Args)-1]); err != nil {
		fmt.Fprintln(os.Stderr, "bad argument", os.Args[len(os.Args)-1])
		return
	}

	steps := wp.Problem{caps, goal}.Solve()
	for _, step := range steps {
		fmt.Println(step)
	}
}
