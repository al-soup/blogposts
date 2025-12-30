package renderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/al-soup/blogposts/internal/blogposts"
)

var (
	// Embed gives access to files embedded in the running Go program
	// Can init vars of type string, []byte or FS
	// Using embed instead of mounting it through the FS brings simplicity
	//go:embed "templates/*"
	postTemplate embed.FS
)

// Accept `io.Writer` for flexibility: write to `os.File`, `http.ResponseWriter` or anything else
func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
