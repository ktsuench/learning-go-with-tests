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

	return Post{
		Title:       readline(titleSeparator),
		Description: readline(descriptionSeparator)}, nil
}
