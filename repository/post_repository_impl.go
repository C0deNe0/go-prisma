package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/C0deNe0/go-prisma/models"
	"github.com/C0deNe0/go-prisma/prisma/db"
	"github.com/C0deNe0/go-prisma/utils"
)

type PostRepositoryImpl struct {
	Db *db.PrismaClient
}

// Delete implements PostRepository.
func (p *PostRepositoryImpl) Delete(ctx context.Context, postId string) {
	result, err := p.Db.Post.FindUnique(db.Post.ID.Equals(postId)).Delete().Exec(ctx)
	utils.ErrorPanic(err)
	fmt.Println("Rows affected ", result)
}

// FindAll implements PostRepository.
func (p *PostRepositoryImpl) FindAll(ctx context.Context) []models.Post {
	allPost, err := p.Db.Post.FindMany().Exec(ctx)
	utils.ErrorPanic(err)

	var posts []models.Post

	for _, post := range allPost {
		published := post.Published
		description, _ := post.Description()

		postData := models.Post{
			Id:          post.ID,
			Title:       post.Title,
			Published:   published,
			Description: description,
		}

		posts = append(posts, postData)
	}

	return posts

}

// FindById implements PostRepository.
func (p *PostRepositoryImpl) FindById(ctx context.Context, postId string) (models.Post, error) {
	post, err := p.Db.Post.FindFirst(db.Post.ID.Equals(postId)).Exec(ctx)
	utils.ErrorPanic(err)

	if post == nil {
		return models.Post{}, errors.New("post id not found")
	}

	published := post.Published
	description, _ := post.Description()
	postData := models.Post{
		Id:          post.ID,
		Title:       post.Title,
		Published:   published,
		Description: description,
	}

	return postData, errors.New("post id not found")

}

// Save implements PostRepository.
func (p *PostRepositoryImpl) Save(ctx context.Context, post models.Post) {
	result, err := p.Db.Post.CreateOne(
		db.Post.Title.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	utils.ErrorPanic(err)

	fmt.Println("Rows updated", result)
}

// Update implements PostRepository.
func (p *PostRepositoryImpl) Update(ctx context.Context, post models.Post) {
	result, err := p.Db.Post.FindMany(db.Post.ID.Equals(post.Id)).Update(
		db.Post.ID.Set(post.Title),
		db.Post.Published.Set(post.Published),
		db.Post.Description.Set(post.Description),
	).Exec(ctx)
	utils.ErrorPanic(err)

	fmt.Println("Rows affected: ", result)
}

func NewPostRepositoryImpl(Db *db.PrismaClient) PostRepository {
	return &PostRepositoryImpl{
		Db: Db,
	}
}
