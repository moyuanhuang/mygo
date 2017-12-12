package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
}
func onces() {
	fmt.Println("onces")
}

// output:
// onces
// count:  --- 0
// count:  --- 1
// count:  --- 2
// count:  --- 3
// count:  --- 4
// count:  --- 5
// count:  --- 6
// count:  --- 7
// count:  --- 8
// count:  --- 9
