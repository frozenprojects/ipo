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

// Force interface implementation
var _ ipo.Input = (*NetworkImage)(nil)

// NetworkImage fetches an image from a URL.
type NetworkImage struct {
	URL string
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

	img, _, err := image.Decode(bytes.NewReader(response.Bytes()))
	return img, err
}
