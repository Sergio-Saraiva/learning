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

	type Car struct {
		color string;
		fabrication int32;
	}

	fmt.Println("Queue:")
	q := NewQueue[Car]();

	fmt.Println(q.length);
	fmt.Println(q.Peek());
	fmt.Println(q.Peek());

	q.enqueue(Car{color: "red", fabrication: 2020});
	q.enqueue(Car{color: "green", fabrication: 2019});
	q.enqueue(Car{color: "yellow", fabrication: 2023});
	
	fmt.Println(q.length);
	fmt.Println(q.Peek());
	fmt.Println(q.Peek());
	
	fmt.Println(q.deque());
	fmt.Println(q.deque());
	fmt.Println(q.deque());
	fmt.Println(q.deque());
	fmt.Println(q.length);
	fmt.Println(q.Peek());

	fmt.Println("Stack:")

	s := NewStack[Car]();

	s.Push(Car{color: "red", fabrication: 2020});
	s.Push(Car{color: "green", fabrication: 2019});
	s.Push(Car{color: "yellow", fabrication: 2023});

	fmt.Println(s.Peek());
	fmt.Println(s.Pop())
	fmt.Println(s.Peek());
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
}