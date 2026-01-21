package renderer_test

import (
	"bytes"
	"flag"
	"io"
	"os"
	"testing"

	"github.com/al-soup/blogposts/internal/blogposts"
	"github.com/al-soup/blogposts/internal/renderer"
)

// update test data with `go test -update`
var update = flag.Bool("update", false, "update golden files")

func TestRender(t *testing.T) {
	post := blogposts.Post{
		Title:       "Title",
		Description: "Description",
		Tags:        []string{"tag1", "tag2"},
		Body:        "Body",
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		postRenderer, err := renderer.NewPostRenderer()

		if err != nil {
			t.Fatal(err)
		}

		if err := postRenderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		goldenFile := "testdata/post.golden"

		if *update {
			os.WriteFile(goldenFile, []byte(got), 0644)
		}

		want, err := os.ReadFile(goldenFile)
		if err != nil {
			t.Fatal(err)
		}

		if got != string(want) {
			t.Errorf("mismatch in generated template:\n%s", got)
		}
	})
}

// run with `go test -bench=.`
func BenchmarkRenderer(b *testing.B) {
	var (
		post = blogposts.Post{
			Title:       "Title",
			Description: "Description",
			Tags:        []string{"tag1", "tag2"},
			Body:        "Body",
		}
	)
	for b.Loop() {
		postRenderer, err := renderer.NewPostRenderer()

		if err != nil {
			b.Fatal(err)
		}
		postRenderer.Render(io.Discard, post)
	}

}
