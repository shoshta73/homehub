package user

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/shoshta73/homehub/constants"
	"github.com/shoshta73/homehub/log"
)

const defaultSize = 1024

func GenerateIdenticon(u User) error {
	log.Info("Generating identicon for", "user", u)
	const gridSize = 128
	cellSize := defaultSize / gridSize

	img := image.NewRGBA(image.Rect(0, 0, defaultSize, defaultSize))
	backgroundColor := color.White

	for x := 0; x < defaultSize; x++ {
		for y := 0; y < defaultSize; y++ {
			img.Set(x, y, backgroundColor)
		}
	}

	fillColor := parseColorFromHash(u.Avatar[:], 0)

	for x := 0; x < gridSize; x++ {
		for y := 0; y <= gridSize; y++ {
			// for y := 0; y <= gridSize/2; y++ {
			idx := (x*gridSize + y) % len(u.Avatar)
			if idx >= len(u.Avatar) {
				idx = idx % len(u.Avatar)
			}
			if u.Avatar[idx]%2 == 0 {
				fillRectangle(img, x*cellSize, y*cellSize, cellSize, cellSize, fillColor)
			} else {
				fillRectangle(img, x*cellSize, y*cellSize, cellSize, cellSize, parseColorFromHash(u.Avatar[:], idx))
			}
		}
	}

	_, err := os.Stat(filepath.Join(constants.DATA_DIR, "identicons"))
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(constants.DATA_DIR, "identicons"), 0755)
		} else {
			return err
		}
	}

	outFile, err := os.Create(filepath.Join(constants.DATA_DIR, "identicons", u.Username+".png"))
	if err != nil {
		return err
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func parseColorFromHash(hashStr string, ofset int) color.Color {
	col1 := uint8(hashStr[(ofset+0)%len(hashStr)])
	col2 := uint8(hashStr[(ofset+1)%len(hashStr)])
	col3 := uint8(hashStr[(ofset+2)%len(hashStr)])
	col4 := uint8(hashStr[(ofset+3)%len(hashStr)])

	log.Info("Taking red channel from", "idx", col1)
	log.Info("Taking green channel from", "idx", col2)
	log.Info("Taking blue channel from", "idx", col3)
	log.Info("Taking alpha channel from", "idx", col4)

	r := uint8(hashStr[col1])
	g := uint8(hashStr[col2])
	b := uint8(hashStr[col3])
	a := uint8(hashStr[col4])

	log.Info("Color", "r", r, "g", g, "b", b, "a", a)
	return color.RGBA{r, g, b, a}
}

func fillRectangle(img *image.RGBA, x, y, width, height int, col color.Color) {
	for i := x; i < x+width; i++ {
		for j := y; j < y+height; j++ {
			img.Set(i, j, col)
		}
	}
}
