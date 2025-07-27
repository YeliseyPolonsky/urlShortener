package link

type CreateRequest struct {
	Url string `json:"url" validate:"required"`
}

type CreateResponse struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}
