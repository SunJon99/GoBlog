package main

import (
	loaders "GoBlog/loaders"
	"fmt"
)

func main() {
	blog := loaders.LoadOneBlogByID(1)
	fmt.Println(blog)
}
