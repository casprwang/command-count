package main

import (
	"fmt"

	"github.com/wangsongiam/command-count"
)

func main() {

	ss := cmc.GetSlice()

	for _, kv := range ss {
		fmt.Printf("%s %d %v%% \n", kv.Key, kv.Val, GetPercentage(kv.Val, counter))
	}

	fmt.Printf("Total executed commands: %d", counter)
}
