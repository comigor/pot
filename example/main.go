package main

import (
	"net/http"

	"github.com/afikrim/pot/example/handler"
	"github.com/afikrim/pot/example/handler/pb/blog"
)

func main() {
	blogHandler := handler.NewBlogHandler()
	httpHandler := blog.RegisterBlogServiceHTTPServer(blogHandler)

	http.ListenAndServe(":8080", httpHandler)
}
