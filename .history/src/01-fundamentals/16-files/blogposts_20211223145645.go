package blogposts

import "io/fs"

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(dir fs.FS) []Post {
	return []Post{}
}
