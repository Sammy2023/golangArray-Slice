package main 
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Enter your number: ")
	var x int

	fmt.Scanln(&x)
	fmt.Println(x)
	var array = make([]int, x)
	fmt.Println(array)

	m := make(map[int]int)
	start := time.Now()

	for i:=0; i<x; i++ {
		array[i] = i+1
	}
	elapsed := time.Since(start)
	fmt.Println("Time took to intialize array %s", elapsed)

	start = time.Now()
	for i:=0; i<x+1; i++ {
		m[i] = i+1
	}

	elapsed = time.Since(start)
	fmt.Println("Time took to intialize map %s", elapsed)

	start = time.Now()
	var slice []int = array[:]
	elapsed = time.Since(start)
	fmt.Println("Time took to intialize slice %s", elapsed)

	fmt.Println()

	start = time.Now()
	for i:=0; i<x+1; i++ {
		array[i] +=1; 
	}
	elapsed = time.Since(start)
	fmt.Println("Time took to increment array %s", elapsed)

	start = time.Now()
	for i:=0; i<x+1; i++ {
		m[i] +=1; 
	}
	elapsed = time.Since(start)
	fmt.Println("Time took to increment array %s", elapsed)

	start = time.Now()
	for i:=0; i<x+1; i++ {
		slice[i] +=1; 
	}
	elapsed = time.Since(start)
	fmt.Println("Time took to increment slice %s", elapsed)

	fmt.Println(m)
	fmt.Println(array)
	fmt.Println(slice)
}