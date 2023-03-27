package imagex

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"strings"
)

const (
	end         = "\x1b[39;49m\n"
	pixelFormat = "\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm\u2580"
)

func ConvertImage(img image.Image, cols int) (output string) {
	img = resizeImage(img, cols)
	bounds := img.Bounds()
	lines := make([]string, bounds.Max.Y)

	for row := 0; row < bounds.Max.Y; row += 2 {
		line := make([]string, bounds.Max.X+1)

		for col := 0; col <= bounds.Max.X; col++ {
			pixelUp := img.At(col, row)
			pixelDown := img.At(col, row+1)

			ru, gu, bu, _ := pixelUp.RGBA()
			rd, gd, bd, _ := pixelDown.RGBA()

			line[col] = fmt.Sprintf(pixelFormat,
				ru/257, gu/257, bu/257,
				rd/257, gd/257, bd/257)
		}
		line[bounds.Max.X] = end
		lines[row] = strings.Join(line, "")
	}
	return strings.Join(lines, "")
}

func resizeImage(img image.Image, cols int) image.Image {
	width := uint(cols)
	height := uint(float64(cols) / float64(img.Bounds().Max.X) * float64(img.Bounds().Max.Y) * 0.8)

	return resize.Resize(width, height, img, resize.Lanczos3)
}
