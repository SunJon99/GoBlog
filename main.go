package main

import (
	loaders "GoBlog/loaders"
	"fmt"
)

func main() {

	isDele := loaders.DelByID(12)
	fmt.Println(isDele)
}
