package main

import (
	"fmt"
	fp_golang "fp-golang"
)

func main() {
	v := fp_golang.NewList[fp_golang.List[int]]()
	for i := 0; i < 10; i++ {
		v.Add(fp_golang.ArrayToList(i))
	}
	v.ToStream().ForEach(func(j fp_golang.List[int]) {
		fmt.Println(j)
	})

	fp_golang.FlatMap[fp_golang.List[int], int](v.ToStream(), func(ints fp_golang.List[int]) fp_golang.Stream[int] {
		return ints.ToStream()
	}).ForEach(func(j int) {
		fmt.Println(j)
	})

}
