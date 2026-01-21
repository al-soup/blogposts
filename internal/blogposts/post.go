package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// TODO check for correct formatting
// TODO check for correct order
// TODO render MarkDown to HTML in the body

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line

	buf := bytes.Buffer{}
	// return bool whether there is more to scan
	for scanner.Scan() {
		// scanner removes newline
		fmt.Fprintln(&buf, scanner.Text())
	}
	// remove trailing newline
	return strings.TrimSuffix(buf.String(), "\n")
}
