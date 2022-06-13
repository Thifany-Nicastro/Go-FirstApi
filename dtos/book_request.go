package dtos

type BookRequestBody struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type BookRequestParams struct {
	ID string `uri:"id" binding:"uuid" json:"id"`
}
