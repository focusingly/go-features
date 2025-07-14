package main

import (
	"bindgen"
	"fmt"
	"sync"
)

func main() {
	bindgen.PrintGOFuncToCExtern()
	bindgen.PrintCStructInGO()
	v1 := make([]float32, 16)
	v2 := make([]float32, 16)
	for i := range v1 {
		v1[i] = 1.2
		v2[i] = 1.2
	}

	fmt.Println()
	fmt.Println("============= avx 256 matrix calc =================")
	fmt.Println(bindgen.AVX2MatAdd(v1, v2))
	fmt.Println(bindgen.AVX2MatSub(v1, v2))

	fmt.Println()
	fmt.Println("CGO Atomic ID Gen")
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for d := range 10 {
				fmt.Printf("[coroutine:%d - loop: [%d]: ---> cgo id: [%d]\n", i, d, bindgen.GetCGONextID())
			}
		}(i)
	}

	wg.Wait()
}
