package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/ktsuench/learning-go-with-tests/src/01-fundamentals/16-files"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("failed to read files", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFs{})

		if err == nil {
			t.Error("expected an error")
		}
	})
	t.Run("read files", func(t *testing.T) {
		fst := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blogposts.NewPostsFromFS(fst)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fst) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fst))
		}

		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("failed to open file")
}