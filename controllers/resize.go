package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Kyutech-C3/img-resize-server/schemas"
	"github.com/Kyutech-C3/img-resize-server/services"

	"github.com/gin-gonic/gin"
)

// @Summary 画像のリサイズ
// @Tags Resize
// @Produce  json
// @Accept multipart/form-data
// @Param file formData file true "this is a file"
// @Param w query string false "width" format(int64)
// @Param h query string false "height" format(int64)
// @Param size query string false "size" format(int64)
// @Param q query string false "quality" format(int64)
// @Param f query string false "format" format(string)
// @Success 200 {object} schemas.ResizeResponse
// @Failure 400 {object} error
// @Router /api/v1/resize [post]
func PostResize(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get file",
		})
		return
	}

	var query schemas.ResizeRequestQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bad query",
		})
	}

	log.Println(query)

	filenames := strings.Split(file.Filename, ".")

	filename := filenames[0]
	extension := filenames[1]
	log.Println(filename)
	log.Println(extension)

	encodeFormat := "webp"

	value, err := file.Open()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to open file",
		})
		return
	}
	defer value.Close()

	resizedImg, err := services.Resize(ctx, value)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to resize",
		})
		return
	}

	formattedFile, err := services.Format(ctx, resizedImg, extension, encodeFormat)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to format",
		})
		return
	}

	// ↓debug
	imgfile, err := os.Create("images/" + filename + "_resized." + encodeFormat)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to create file",
		})
		return
	}
	defer imgfile.Close()

	imgfile.Write(formattedFile.Bytes())

	print(formattedFile)
	// ↑debug

	ctx.Data(http.StatusOK, "image/"+encodeFormat, formattedFile.Bytes())
}

// @Summary 画像のリサイズ
// @Tags Resize
// @Produce  json
// @Param url query string false "url" format(string)
// @Param w query string false "width" format(int64)
// @Param h query string false "height" format(int64)
// @Param size query string false "size" format(int64)
// @Param q query string false "quality" format(int64)
// @Param f query string false "format" format(string)
// @Success 200 {object} schemas.ResizeResponse
// @Failure 400 {object} error
// @Router /api/v1/resize [get]
func GetResize(ctx *gin.Context) {
	var query schemas.GetResizeRequestQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bad query",
		})
		return
	}

	if query.Url == "" {
		log.Println("url is empty")
		ctx.JSON(400, gin.H{
			"message": "url is empty",
		})
		return
	}

	encodeFormat := "webp"
	url := query.Url
	urlSplint := strings.Split(url, "/")
	filenameSprit := strings.Split(urlSplint[len(urlSplint)-1], ".")
	extension := filenameSprit[len(filenameSprit)-1]

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	value := bytes.NewBuffer([]byte{})
	io.Copy(value, response.Body)

	// log.Println(value)

	log.Println(response)
	resizedImg, err := services.Resize(ctx, value)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to resize",
		})
		return
	}

	formattedFile, err := services.Format(ctx, resizedImg, extension, encodeFormat)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to format",
		})
		return
	}
	ctx.Data(http.StatusOK, "image/"+encodeFormat, formattedFile.Bytes())
}

func GetResizeStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println(id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
