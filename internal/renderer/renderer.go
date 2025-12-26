package renderer

import (
	"fmt"
	"io"

	"github.com/al-soup/blogposts/internal/blogposts"
)

// Accept `io.Writer` for flexibility: write to `os.File`, `http.ResponseWriter` or anything else
func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
