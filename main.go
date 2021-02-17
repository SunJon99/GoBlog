package main

import (
	loaders "GoBlog/loaders"
	"fmt"
)

func main() {

	var bType *int
	var title *string
	var recommend *int
	bType = new(int)
	*bType = 5

	blogs := loaders.LoadBlogsByConditions(bType, title, recommend)
	fmt.Println(blogs)
}
