package request

//It is the format of the payload to be expected when creating a new post, it consists of three fields
//Title
//Publish
//Description
type PostCreateRequest struct {
	Title       string `validate:"required min=1,max=100" json:"title"`
	Published   bool   `json:"published"` // it is a boolean which indicates the publication status of the code.
	Description string `json:"description"`  // string providing additional details about the post
}
