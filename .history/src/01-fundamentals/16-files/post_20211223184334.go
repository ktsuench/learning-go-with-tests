package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readline := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readline(titleSeparator)
	description := readline(descriptionSeparator)
	tags := strings.Split(readline(tagsSeparator), ",")
	for i, tag := range tags {
		tags[i] = strings.TrimSpace(tag)
	}

	return Post{
		title,
		description,
		tags,
	}, nil
}
