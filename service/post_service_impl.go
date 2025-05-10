package service

import (
	"context"

	"github.com/C0deNe0/go-prisma/data/request"
	"github.com/C0deNe0/go-prisma/data/response"
	"github.com/C0deNe0/go-prisma/models"
	"github.com/C0deNe0/go-prisma/repository"
	"github.com/C0deNe0/go-prisma/utils"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

// Create implements PostService.
// This func in the post service is responsible for persisting a new post, prior to executing the save operation
// we perform a mapping of the request data to post model
func (p *PostServiceImpl) Create(ctx context.Context, request request.PostCreateRequest) {
	postData := models.Post{
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.FindById(ctx, postId)
	utils.ErrorPanic(err)

	p.PostRepository.Delete(ctx, post.Id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postResp []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Published:   value.Published,
			Description: value.Description,
		}
		postResp = append(postResp, post)
	}

	return postResp
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(ctx context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	utils.ErrorPanic(err)

	postResponse := response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}
	return postResponse
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, request request.PostUpdateRequest) {
	postData := models.Post{
		Id:          request.Id,
		Title:       request.Title,
		Published:   request.Published,
		Description: request.Description,
	}

	p.PostRepository.Update(ctx, postData)
}

func NewPostServiceImpl(PostR repository.PostRepository) PostService {
	return &PostServiceImpl{PostRepository: PostR}
}
