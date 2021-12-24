package blogposts

import "io/fs"

type Post struct {
}

func NewPostsFromFS(dir fs.FS) []Post {
	return []Post{{}, {}}
}
