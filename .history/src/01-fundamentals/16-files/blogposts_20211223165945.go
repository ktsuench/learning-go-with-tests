package blogposts

import (
	"io"
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
		postFile, err := fileSystem.Open(f.Name())
		if err != nil {
			return nil, err
		}
		defer postFile.Close()

		postData, err := io.ReadAll(postFile)

		posts = append(posts, Post{})
	}
	return posts, nil
}
