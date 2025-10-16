package main

import (
	"github.com/comigor/pot/example/handler"
	"github.com/comigor/pot/example/handler/pb/blog"
	"github.com/labstack/echo"
)

func main() {
	blogHandler := handler.NewBlogHandler()
	httpHandler := blog.RegisterBlogServiceHTTPServer(blogHandler)

	// native http server
	// mux := http.NewServeMux()
	// mux.Handle("/", httpHandler)
	// http.ListenAndServe(":8080", mux)

	// using with echo
	e := echo.New()
	e.Any("v1/*{blog_handler}", echo.WrapHandler(httpHandler))
	e.Start(":8080")
}
