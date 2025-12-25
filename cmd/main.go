package main

import (
	"log"
	"os"

	post "github.com/al-soup/blogposts/internal/blogposts"
)

func main() {
	posts, err := post.NewBlogPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
