package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		postFile, _ := fileSystem.Open(f.Name())

		b := []byte{}
		postFile.Read(b)

		posts = append(posts, Post{})
	}
	return posts, nil
}
