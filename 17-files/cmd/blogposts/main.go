package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/evgenymarkov/learn-golang/17-files/pkg/blogposts"
)

func main() {
	postsDir := os.DirFS("tests/data").(fs.ReadDirFS)
	posts, createErr := blogposts.NewPostsFromFS(postsDir)

	if createErr != nil {
		fmt.Println(createErr)
		os.Exit(1)
	}

	fmt.Println(posts)
}
