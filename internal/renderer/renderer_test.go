package renderer_test

import (
	"bytes"
	"flag"
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

		if err := renderer.Render(&buf, post); err != nil {
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
