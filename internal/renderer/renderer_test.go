package renderer_test

import (
	"bytes"
	"testing"

	"github.com/al-soup/blogposts/internal/blogposts"
	"github.com/al-soup/blogposts/internal/renderer"
)

func TestRender(t *testing.T) {
	post := blogposts.Post{
		Title:       "Title",
		Description: "Description",
		Tags:        []string{"tag1", "tag2"},
		Body:        "Body",
	}

	buf := bytes.Buffer{}
	err := renderer.Render(&buf, post)

	if err != nil {
		t.Fatal(err)
	}

	got := buf.String()
	want := "<h1>Title</h1>"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
