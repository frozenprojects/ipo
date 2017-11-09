package inputs

import (
	"bytes"
	"image"
	"io/ioutil"

	"github.com/aerogo/ipo"
)

// Force interface implementations
var (
	_ Image     = (*FileSystemImage)(nil)
	_ ipo.Input = (*FileSystemImage)(nil)
)

// FileSystemImage fetches an image from a URL.
type FileSystemImage struct {
	URL    string
	img    image.Image
	data   []byte
	format string
}

// Read ...
func (fsImage *FileSystemImage) Read() (obj interface{}, err error) {
	data, err := ioutil.ReadFile(fsImage.URL)

	if err != nil {
		return nil, err
	}

	fsImage.data = data
	fsImage.img, fsImage.format, err = image.Decode(bytes.NewReader(fsImage.data))

	return fsImage, err
}

// Image ...
func (fsImage *FileSystemImage) Image() image.Image {
	return fsImage.img
}

// Data ...
func (fsImage *FileSystemImage) Data() []byte {
	return fsImage.data
}

// Format ...
func (fsImage *FileSystemImage) Format() string {
	return fsImage.format
}

// Extension ...
func (fsImage *FileSystemImage) Extension() string {
	switch fsImage.format {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	case "gif":
		return ".gif"
	case "webp":
		return ".webp"
	default:
		return ""
	}
}
