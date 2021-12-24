package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title, Description string
}

func NewPostsFromFS(fsys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fsys, f.Name())
		// TODO: determine if operation should fail
		// after failure reading file
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fsys fs.FS, filename string) (Post, error) {
	postFile, err := fsys.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readline := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	post := Post{Title: readline()[7:], Description: readline()[13:]}
	return post, nil
}
