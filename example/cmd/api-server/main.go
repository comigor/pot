package main

import (
	"github.com/afikrim/pot/example/handler"
	"github.com/afikrim/pot/example/handler/pb/blog"
	"github.com/labstack/echo/v4"
)

func main() {
	blogHandler := handler.NewBlogHandler()
	httpHandler := blog.RegisterBlogServiceHTTPServer(blogHandler)

	// native http server
	// mux := http.NewServeMux()
	// mux.Handle("v1/*{blog_handler}", httpHandler)
	// http.ListenAndServe(":8080", mux)

	// using with echo
	e := echo.New()
	e.Any("v1/*{blog_handler}", echo.WrapHandler(httpHandler))
	e.Start(":8080")
}
