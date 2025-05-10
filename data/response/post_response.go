package response

import (
	"context"
	"github.com/C0deNe0/go-prisma/models"
)

// when a client request about a information about a post the server can use the postResponse struct to construct a response payload
// containing the necessary detail of the post
type PostResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Published   bool   `json:"published"`
	Description string `json:"description"`
}

// Delete implements repository.PostRepository.
func (p PostResponse) Delete(ctx context.Context, postId string) {
	panic("unimplemented")
}

// FindAll implements repository.PostRepository.
func (p PostResponse) FindAll(ctx context.Context) []models.Post {
	panic("unimplemented")
}

// FindById implements repository.PostRepository.
func (p PostResponse) FindById(ctx context.Context, postId string) (models.Post, error) {
	panic("unimplemented")
}

// Save implements repository.PostRepository.
func (p PostResponse) Save(ctx context.Context, post models.Post) {
	panic("unimplemented")
}

// Update implements repository.PostRepository.
func (p PostResponse) Update(ctx context.Context, post models.Post) {
	panic("unimplemented")
}
