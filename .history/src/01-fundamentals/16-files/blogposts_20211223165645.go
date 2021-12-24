package blogposts

import (
	"bytes"
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

		b := bytes.Buffer{}

		postFile.Read()
		posts = append(posts, Post{})
	}
	return posts, nil
}
