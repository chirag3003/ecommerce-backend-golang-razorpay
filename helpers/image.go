package helpers

import (
	"bytes"
	"github.com/h2non/bimg"
	"io"
	"mime/multipart"
)

func OptimizeImage(file *multipart.File) (*bytes.Reader, error) {
	buffer, err := io.ReadAll(*file)
	if err != nil {
		panic(err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	return bytes.NewReader(newImage), err
}
