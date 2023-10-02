package main

import (
	"fmt"
)

func binary_search(haystack[]int, needle int) bool {
	lo := 0;
	hi := len(haystack);
	
	
	for {
		
		m := int(lo + (hi - lo) / 2);
		
		v := haystack[m];
		
		if v == needle {
			fmt.Println("Found")
			return true
		} else if v > needle {
			hi = m;
		} else {
			lo = m + 1;
		}

		if !(lo < hi){
			break;
		}
		
	}

	fmt.Println("Not found");
	return false
}