package main

import (
	"fmt"

	"github.com/wangsongiam/command-count"
)

func main() {

	ss, counter := cmc.GetSlice()

	for _, kv := range ss {
		fmt.Printf("%s-> %.2f%%\nCount: %d \n", kv.Key, cmc.GetPercentage(kv.Val, counter), kv.Val)
	}

	fmt.Printf("Total executed commands: %d", counter)
}
