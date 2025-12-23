package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/al-soup/blogposts"
)

func TestNewBlogPost(t *testing.T) {
	// in-memory mock of file-system
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hallo")},
	}

	posts := blogposts.NewBlogPostFromFs(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
