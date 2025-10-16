package handler

import (
	"context"

	"github.com/comigor/pot/example/handler/pb/blog"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

var _ blog.BlogServiceHTTPServer = &BlogHandler{}

type BlogHandler struct{}

// CreateBlog implements blog.BlogServiceHTTPServer.
func (b *BlogHandler) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.CreateBlogResponse, error) {
	return &blog.CreateBlogResponse{
		Id:       uuid.New().String(),
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}, nil
}

// DeleteBlog implements blog.BlogServiceHTTPServer.
func (b *BlogHandler) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// ReadBlog implements blog.BlogServiceHTTPServer.
func (b *BlogHandler) ReadBlog(ctx context.Context, in *blog.ReadBlogRequest) (*blog.ReadBlogResponse, error) {
	return &blog.ReadBlogResponse{
		Id:       in.Id,
		AuthorId: "Aziz",
		Title:    "Title of the current blogs",
		Content:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}, nil
}

// UpdateBlog implements blog.BlogServiceHTTPServer.
func (b *BlogHandler) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.UpdateBlogResponse, error) {
	return &blog.UpdateBlogResponse{
		Id:       in.Id,
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}, nil
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{}
}
