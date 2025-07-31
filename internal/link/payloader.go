package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkCreateResponse struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}
