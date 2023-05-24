package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}

	fmt.Printf("\n%T \n", sum)
	fmt.Println(sum)

	arr := []int{1, 2, 3}

	for _, value := range arr {
		fmt.Printf("%d ", value)
	}

}
