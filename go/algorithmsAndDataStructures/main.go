package main

import "fmt"

func main() {
	fmt.Println("Linear search:")
	linear_search([]int{1, 2, 3, 4, 5}, 3)

	fmt.Println("Binary search:")
	binary_search([]int{1, 2, 3, 4, 5, 6, 7, 8 , 9}, 9);

	fmt.Println("Two Crystal Balls: ")
	fmt.Println(two_crystal_balls([]bool {false, false, false, false, false, false, false, false, true, true, true, true, true, true , true}));

	fmt.Println("Bubble sort: ")
	fmt.Println(bubble_sort([]int {2,5,4,6,8,9,10,3}))

	fmt.Println("teste")
}
