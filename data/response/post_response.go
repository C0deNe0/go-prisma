package response

//when a client request about a information about a post the server can use the postResponse struct to construct a response payload 
//containing the necessary detail of the post
type PostResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Published   bool   `json:"published"`
	Description string `json:"description"`
}
