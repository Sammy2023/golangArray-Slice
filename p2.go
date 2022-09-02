package main 

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
func main() {
	fmt.Println("Enter your number: ")
	var x int
	fmt.Scanln(&x)
	var array = make([]int, x)
	for i:=0; i<x; i++ {
		array[i] = rand.Intn(10)
	}
	var slice = array[:]

	start := time.Now()
	sort.Ints(array)
	elapsed := time.Since(start)
	fmt.Println("Time to sort array", elapsed)
	fmt.Println(array)


	start = time.Now()
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	   
	elapsed = time.Since(start)
	fmt.Println("Time took to sort slice", elapsed)

	for i:=0; i<x; i++ {
		array[i] = rand.Intn(10)
	}
	slice = array[:]

	start = time.Now()
	sort.SliceStable(array, func(a, b int) bool { 
		return array[a] < array[b] })
	elapsed = time.Since(start)
	fmt.Println("Time took to stable array", elapsed)
	fmt.Println(array)

	start = time.Now()
	sort.SliceStable(slice, func(a, b int) bool { 
		return slice[a] < slice[b] })
	elapsed = time.Since(start)
	fmt.Println("Time took for stable slice", elapsed)
	fmt.Println(slice)
}