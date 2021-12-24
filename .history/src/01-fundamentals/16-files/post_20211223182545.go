package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title, Description string
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
