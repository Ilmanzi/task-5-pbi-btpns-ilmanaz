package app

type PictureRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Caption    string `json:"caption" binding:"required"`
	PictureUrl string `json:"picture_url" binding:"required"`
}
