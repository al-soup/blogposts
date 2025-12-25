package main

import (
	"log"
	"os"

	blogposts "github.com/al-soup/blogposts"
)

func main() {
	posts, err := blogposts.NewBlogPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
