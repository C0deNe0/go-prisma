package service

import (
	"context"
	"github.com/C0deNe0/go-prisma/data/request"
	"github.com/C0deNe0/go-prisma/data/response"
	"github.com/C0deNe0/go-prisma/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

// Create implements PostService.
func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	panic("unimplemented")
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	panic("unimplemented")
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	panic("unimplemented")
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(ctx context.Context, postId string) response.PostResponse {
	panic("unimplemented")
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	panic("unimplemented")
}

func NewPostServiceImpl(PostR repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: PostR}
}
