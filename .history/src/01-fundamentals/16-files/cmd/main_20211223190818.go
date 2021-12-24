package main

import (
	"log"
	"os"

	blogposts "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/16-files"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
