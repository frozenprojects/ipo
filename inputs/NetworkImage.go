package inputs

import (
	"bytes"
	"errors"
	"image"
	"net/http"
	"strconv"

	"github.com/aerogo/ipo"

	"github.com/aerogo/http/client"
)

// Force interface implementations
var (
	_ Image     = (*NetworkImage)(nil)
	_ ipo.Input = (*NetworkImage)(nil)
)

// NetworkImage fetches an image from a URL.
type NetworkImage struct {
	URL    string
	img    image.Image
	data   []byte
	format string
}

// Read ...
func (networkImage *NetworkImage) Read() (obj interface{}, err error) {
	response, err := client.Get(networkImage.URL).End()

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, errors.New("Unexpected status code: " + strconv.Itoa(response.StatusCode()))
	}

	networkImage.data = response.Bytes()
	networkImage.img, networkImage.format, err = image.Decode(bytes.NewReader(networkImage.data))

	return networkImage, err
}

// Image ...
func (networkImage *NetworkImage) Image() image.Image {
	return networkImage.img
}

// Data ...
func (networkImage *NetworkImage) Data() []byte {
	return networkImage.data
}

// Format ...
func (networkImage *NetworkImage) Format() string {
	return networkImage.format
}

// Extension ...
func (networkImage *NetworkImage) Extension() string {
	switch networkImage.format {
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
