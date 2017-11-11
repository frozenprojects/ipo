package outputs

import (
	"errors"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"path"

	"github.com/aerogo/ipo/inputs"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"

	"github.com/aerogo/ipo"
)

// Force interface implementation
var _ ipo.Output = (*ImageFile)(nil)

// The threshold where we start seeing 2 aspect ratios as different
const aspectRatioThreshold = 0.03

// ImageFile writes an image to the filesystem.
type ImageFile struct {
	Directory string
	BaseName  string
	Format    string
	Width     int
	Height    int
	Quality   int
}

// Write ...
func (file *ImageFile) Write(obj interface{}) error {
	input := obj.(inputs.Image)
	extension := input.Extension()

	if extension == "" {
		return errors.New("Unknown format: " + input.Format())
	}

	img := input.Image()
	width := img.Bounds().Dx()
	height := img.Bounds().Dx()
	resizeXRequired := file.Width != 0 && file.Width < width
	resizeYRequired := file.Height != 0 && file.Height < height
	resizeRequired := resizeXRequired || resizeYRequired

	// fmt.Println(file.BaseName+extension, "|", width, "x", height, "|", len(input.Data())/1024, "KB")

	// Original file output
	if file.Format == "" && !resizeRequired {
		fullPath := path.Join(file.Directory, file.BaseName+extension)
		return ioutil.WriteFile(fullPath, input.Data(), 0644)
	}

	// Resize if needed
	if resizeRequired {
		newWidth := uint(0)
		newHeight := uint(0)

		if file.Width != 0 && file.Height != 0 {
			// Take aspect ratio of original image
			aspectRatio := float64(width) / float64(height)

			// Take required aspect ratio
			requiredAspectRatio := float64(file.Width) / float64(file.Height)

			// Decide what to do depending on the aspect ratio differences
			if math.Abs(aspectRatio-requiredAspectRatio) <= aspectRatioThreshold {
				newWidth = uint(file.Width)
				newHeight = uint(file.Height)
			} else if aspectRatio > requiredAspectRatio {
				newHeight = uint(file.Height)
			} else {
				newWidth = uint(file.Width)
			}
		} else {
			newWidth = uint(file.Width)
			newHeight = uint(file.Height)
		}

		img = resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	}

	// Set format automatically if needed
	if file.Format == "" {
		file.Format = input.Format()
	}

	// Write data to file
	fullPath := path.Join(file.Directory, file.BaseName+file.Extension())
	stream, err := os.Create(fullPath)

	if err != nil {
		return err
	}

	defer stream.Close()

	switch file.Format {
	case "jpg", "jpeg":
		err = jpeg.Encode(stream, img, &jpeg.Options{
			Quality: file.Quality,
		})
	case "png":
		err = png.Encode(stream, img)
	case "gif":
		err = gif.Encode(stream, img, nil)
	case "webp":
		err = webp.Encode(stream, img, &webp.Options{
			Quality: float32(file.Quality),
		})
	}

	return err
}

// Extension ...
func (file *ImageFile) Extension() string {
	switch file.Format {
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
