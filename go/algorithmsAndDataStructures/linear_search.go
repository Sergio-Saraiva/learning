package main

import "fmt"

func linear_search(haystack[]int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if(haystack[i] == needle) {
			fmt.Println("Found");
			return true;
		}
	}
	fmt.Println("Not Found");
	return false;
}