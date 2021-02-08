package main

import (
	loaders "GoBlog/loaders"
	"fmt"
)

func main() {
	blog := loaders.LoadAllBlogs()
	fmt.Println(blog)
}
