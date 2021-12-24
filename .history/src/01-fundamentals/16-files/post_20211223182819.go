package blogposts

import (
	"bufio"
	"io"
	"strings"
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

	readline := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readline(titleSeparator)
	description := readline(descriptionSeparator)

	post := Post{title, description}
	return post, nil
}
