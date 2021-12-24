package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/16-files"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("failed to open file")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("failed to read files", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFs{})

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("read files", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
}
