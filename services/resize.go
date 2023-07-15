package services

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"

	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"golang.org/x/image/draw"
)

func Resize(ctx *gin.Context, file io.Reader) (*image.RGBA, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return nil, ctx.Error(err)
	}

	bou := img.Bounds()
	size := bou.Size()
	log.Println(size)

	resizedImg := image.NewRGBA(image.Rect(0, 0, int(size.X/2), int(size.Y/2)))
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	return resizedImg, nil
}

func Format(ctx *gin.Context, img *image.RGBA, extension string, encodeFormat string) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer([]byte{})

	switch encodeFormat {
	case "webp":
		options := webp.Options{Lossless: false, Quality: 100}
		if err := webp.Encode(buf, img, &options); err != nil {
			log.Println(err)
			return nil, err
		}
		return buf, nil
	case "png":
		if err := png.Encode(buf, img); err != nil {
			log.Println(err)
			return nil, err
		}
		return buf, nil
	case "jpeg":
		if err := jpeg.Encode(buf, img, nil); err != nil {
			log.Println(err)
			return nil, err
		}
		return buf, nil
	default:
		log.Println("invalid format")
		return nil, errors.New("invalid format")
	}
}
