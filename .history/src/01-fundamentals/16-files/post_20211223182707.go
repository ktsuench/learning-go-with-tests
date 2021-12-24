package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title, Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readline := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readline()[len(titleSeparator):]
	description := readline()[len(descriptionSeparator):]

	post := Post{title, description}
	return post, nil
}
