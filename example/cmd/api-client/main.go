package main

import (
	"context"
	"fmt"
	"time"

	"github.com/comigor/pot/example/handler/pb/blog"
	"github.com/comigor/pot/option"
	"github.com/google/uuid"
)

func main() {
	blogClient := blog.NewBlogServiceHTTPClient(
		option.WithBaseURL("http://localhost:8080"),
		option.WithTimeout(time.Second*10),
	)

	ctx := context.Background()
	createResp, err := blogClient.CreateBlog(ctx, &blog.CreateBlogRequest{
		AuthorId: "afikrim",
		Title:    "My first blog",
		Content:  "Hello World",
	}, option.WithContentType(option.ContentTypeApplicationJson))
	if err != nil {
		panic(err)
	}
	fmt.Printf("CreateBlogResp: %v\n", createResp)

	ctx = context.Background()
	readResp, err := blogClient.ReadBlog(ctx, &blog.ReadBlogRequest{
		Id: uuid.New().String(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ReadBlogResp: %v\n", readResp)

	ctx = context.Background()
	updateResp, err := blogClient.UpdateBlog(ctx, &blog.UpdateBlogRequest{
		Id:       uuid.New().String(),
		AuthorId: "afikrim",
		Title:    "My Updated blog",
		Content:  "Hello World",
	}, option.WithContentType(option.ContentTypeApplicationJson))
	if err != nil {
		panic(err)
	}
	fmt.Printf("UpdateBlogResp: %v\n", updateResp)

	ctx = context.Background()
	deleteResp, err := blogClient.DeleteBlog(ctx, &blog.DeleteBlogRequest{
		Id: uuid.New().String(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("DeleteBlogResp: %v\n", deleteResp)
}
