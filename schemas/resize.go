package schemas

import "mime/multipart"

type ResizeRequest struct {
	File *multipart.FileHeader `json:"file" binding:"required"`
}

type ResizeRequestQuery struct {
	W    string `form:"w"`
	H    string `form:"h"`
	Size string `form:"size"`
	Q    string `form:"q"`
	F    string `form:"f"`
}

type GetResizeRequestQuery struct {
	Url  string `form:"url"`
	W    string `form:"w"`
	H    string `form:"h"`
	Size string `form:"size"`
	Q    string `form:"q"`
	F    string `form:"f"`
}

type ResizeResponse struct {
	Message string `json:"message"`
}
