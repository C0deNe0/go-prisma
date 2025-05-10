package controllers

import (
	"net/http"

	"github.com/C0deNe0/go-prisma/data/request"
	"github.com/C0deNe0/go-prisma/data/response"
	"github.com/C0deNe0/go-prisma/service"
	"github.com/C0deNe0/go-prisma/utils"
	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (c *PostController) Create(w http.ResponseWriter, r *http.Request) {
	postCreateRequest := request.PostCreateRequest{}
	utils.ReadRequest(r, &postCreateRequest)

	c.PostService.Create(r.Context(), postCreateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	utils.WriteResponse(w, webResponse)
}

func (c *PostController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postUpdateRequest := request.PostUpdateRequest{}
	utils.ReadRequest(r, &postUpdateRequest)

	c.PostService.Update(r.Context(), postUpdateRequest)

	postId := params.ByName("postId")
	postUpdateRequest.Id = postId

	webResponse := response.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   nil,
	}
	utils.WriteResponse(w, webResponse)
}
